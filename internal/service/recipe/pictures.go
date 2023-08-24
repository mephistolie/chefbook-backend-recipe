package recipe

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	recipeFail "github.com/mephistolie/chefbook-backend-recipe/internal/entity/fail"
)

func (s *Service) GenerateRecipePicturesUploadLinks(recipeId, userId uuid.UUID, picturesCount int, subscriptionPlan string) ([]entity.PictureUpload, error) {
	policy, err := s.repo.GetRecipePolicy(recipeId)
	if err != nil || policy.OwnerId != userId || policy.IsEncrypted && s.subscriptionLimiter.IsEncryptionAllowed(subscriptionPlan) {
		return nil, fail.GrpcAccessDenied
	}

	pictureIds, err := s.repo.GetRecipePictureIdsToUpload(recipeId, picturesCount)
	if err != nil {
		return nil, err
	}

	var uploads []entity.PictureUpload
	for _, pictureId := range pictureIds {
		if upload, err := s.s3.GenerateRecipePictureUploadLink(recipeId, pictureId, subscriptionPlan, policy.IsEncrypted); err == nil {
			upload.PictureId = pictureId
			uploads = append(uploads, upload)
		}
	}

	return uploads, nil
}

func (s *Service) SetRecipePictures(
	recipeId,
	userId uuid.UUID,
	pictures entity.RecipePictureIds,
	version *int32,
	subscriptionPlan string,
) (map[uuid.UUID]string, int32, error) {
	if policy, err := s.repo.GetRecipePolicy(recipeId); err != nil || policy.OwnerId != userId {
		return map[uuid.UUID]string{}, 0, fail.GrpcAccessDenied
	}

	pictureIds := pictures.GetIds()
	maxPicturesCount := s.subscriptionLimiter.GetMaxPicturesCount(subscriptionPlan)
	if len(pictureIds) > maxPicturesCount {
		return map[uuid.UUID]string{}, 0, recipeFail.GrpcRecipePicturesCountLimit
	}

	if !s.s3.CheckRecipePicturesExist(recipeId, pictures.GetIds()) {
		return map[uuid.UUID]string{}, 0, recipeFail.GrpcRecipePictureNotFound
	}

	newVersion, err := s.repo.SetRecipePictures(recipeId, pictures, version)
	if err != nil {
		return map[uuid.UUID]string{}, 0, err
	}

	go func() {
		s.s3.DeleteUnusedRecipePictures(recipeId, pictureIds)
	}()

	links := make(map[uuid.UUID]string)
	if pictures.Preview != nil {
		links[*pictures.Preview] = s.s3.GetRecipePictureLink(recipeId, *pictures.Preview)
	}
	for _, stepPictures := range pictures.Cooking {
		for _, pictureId := range stepPictures {
			links[pictureId] = s.s3.GetRecipePictureLink(recipeId, pictureId)
		}
	}

	return links, newVersion, nil
}
