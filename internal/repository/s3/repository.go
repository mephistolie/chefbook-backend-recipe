package s3

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"github.com/mephistolie/chefbook-backend-recipe/internal/config"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	"github.com/mephistolie/chefbook-backend-recipe/internal/helpers"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"time"
)

const (
	recipesDir  = "recipes"
	picturesDir = "pictures"

	idLength = 36
)

type Repository struct {
	client              *minio.Client
	bucket              string
	subscriptionLimiter helpers.SubscriptionLimiter
}

func NewRepository(cfg *config.Config, subscriptionLimiter helpers.SubscriptionLimiter) (*Repository, error) {
	client, err := minio.New(*cfg.S3.Host, &minio.Options{
		Creds:  credentials.NewStaticV4(*cfg.S3.AccessKeyId, *cfg.S3.SecretAccessKey, ""),
		Secure: true,
		Region: *cfg.S3.Region,
	})
	if err != nil {
		return nil, err
	}

	return &Repository{
		client: client,
		bucket: *cfg.S3.Bucket,

		subscriptionLimiter: subscriptionLimiter,
	}, nil
}

func (r *Repository) GetRecipePictureLink(recipeId, pictureId uuid.UUID) string {
	objectPath := r.getRecipePicturePath(recipeId, pictureId)
	return fmt.Sprintf("https://%s/%s", r.bucket, objectPath)
}

func (r *Repository) GenerateRecipePictureUploadLink(recipeId, pictureId uuid.UUID, subscriptionPlan string, isEncrypted bool) (entity.PictureUpload, error) {
	maxPictureSize := r.subscriptionLimiter.GetPictureMaxSize(subscriptionPlan)
	return r.generatePictureUploadLink(pictureId, r.getRecipePicturePath(recipeId, pictureId), maxPictureSize, isEncrypted)
}

func (r *Repository) CheckRecipePicturesExist(recipeId uuid.UUID, pictures []uuid.UUID) bool {
	picturesPath := fmt.Sprintf("%s/%s/%s", recipesDir, recipeId, picturesDir)
	existingPictures := make(map[uuid.UUID]bool)

	for object := range r.client.ListObjects(context.Background(), r.bucket, minio.ListObjectsOptions{
		Prefix:    picturesPath,
		Recursive: true,
	}) {
		keyLength := len(object.Key)
		rawPicureId := object.Key[keyLength-idLength : keyLength]
		pictureId, err := uuid.Parse(rawPicureId)
		if err != nil {
			log.Debugf("unable to parse picture id by key %s: %s", object.Key, err)
			continue
		}
		existingPictures[pictureId] = true
	}

	for _, picture := range pictures {
		if exists, ok := existingPictures[picture]; !ok || !exists {
			return false
		}
	}

	return true
}

func (r *Repository) DeleteUnusedRecipePictures(recipeId uuid.UUID, usedPictures []uuid.UUID) {
	picturesPath := fmt.Sprintf("%s/%s/%s", recipesDir, recipeId, picturesDir)
	opts := minio.RemoveObjectOptions{ForceDelete: true}

	usedPicturesMap := make(map[uuid.UUID]bool)
	for _, usedPicture := range usedPictures {
		usedPicturesMap[usedPicture] = true
	}

	for object := range r.client.ListObjects(context.Background(), r.bucket, minio.ListObjectsOptions{
		Prefix:    picturesPath,
		Recursive: true,
	}) {
		keyLength := len(object.Key)
		rawPictureId := object.Key[keyLength-idLength : keyLength]
		pictureId, err := uuid.Parse(rawPictureId)
		if err != nil {
			log.Debugf("unable to parse picture id by key %s: %s", object.Key, err)
			continue
		}

		if exists, ok := usedPicturesMap[pictureId]; !ok || !exists {
			if err = r.client.RemoveObject(context.Background(), r.bucket, object.Key, opts); err != nil {
				log.Warn("unable to delete picture %s: %s", object.Key, err)
			}
		}
	}
}

func (r *Repository) getRecipePicturePath(recipeId, pictureId uuid.UUID) string {
	return fmt.Sprintf("%s/%s/%s/%s", recipesDir, recipeId, picturesDir, pictureId)
}

func (r *Repository) generatePictureUploadLink(pictureId uuid.UUID, objectName string, maxSize int64, isEncrypted bool) (entity.PictureUpload, error) {
	policy := minio.NewPostPolicy()

	if err := policy.SetBucket(r.bucket); err != nil {
		log.Error("unable to set bucket in post policy: ", err)
		return entity.PictureUpload{}, fail.GrpcUnknown
	}
	if err := policy.SetKey(objectName); err != nil {
		log.Errorf("unable to set object %s in post policy: %s", objectName, err)
		return entity.PictureUpload{}, fail.GrpcUnknown
	}
	if !isEncrypted {
		if err := policy.SetContentTypeStartsWith("image"); err != nil {
			log.Errorf("unable to set content type in post policy: %s", objectName, err)
			return entity.PictureUpload{}, fail.GrpcUnknown
		}
	}
	if err := policy.SetContentLengthRange(0, maxSize); err != nil {
		log.Errorf("unable to set content length in post policy: %s", objectName, err)
		return entity.PictureUpload{}, fail.GrpcUnknown
	}
	if err := policy.SetExpires(time.Now().Add(1 * time.Hour)); err != nil {
		log.Errorf("unable to set expiration in post policy: %s", objectName, err)
		return entity.PictureUpload{}, fail.GrpcUnknown
	}

	originalUrl, formData, err := r.client.PresignedPostPolicy(context.Background(), policy)
	if err != nil {
		log.Errorf("unable to generate presigned link for uploading object %s: %s", objectName, err)
		return entity.PictureUpload{}, fail.GrpcUnknown
	}
	_ = fmt.Sprintf("https://%s", r.bucket)

	return entity.PictureUpload{
		PictureId: pictureId,
		URL:       originalUrl.String(),
		FormData:  formData,
		MaxSize:   maxSize,
	}, nil
}
