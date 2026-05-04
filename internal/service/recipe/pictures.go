package recipe

import (
	"context"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	recipeFail "github.com/mephistolie/chefbook-backend-recipe/internal/entity/fail"
)

func (s *Service) GenerateRecipePicturesUploadLinks(ctx context.Context, recipeId, userId uuid.UUID, picturesCount int, subscriptionPlan string) ([]entity.PictureUpload, error) {
	policy, err := s.recipeRepo.GetRecipePolicy(ctx, recipeId)
	if err != nil || policy.OwnerId != userId || policy.IsEncrypted && !s.subscriptionLimiter.IsEncryptionAllowed(subscriptionPlan) {
		return nil, fail.GrpcAccessDenied
	}

	pictureIds, err := s.recipeRepo.GetRecipePictureIdsToUpload(ctx, recipeId, picturesCount)
	if err != nil {
		return nil, err
	}

	var uploads []entity.PictureUpload
	for _, pictureId := range pictureIds {
		if upload, err := s.s3.GenerateRecipePictureUploadLink(ctx, recipeId, pictureId, subscriptionPlan, policy.IsEncrypted); err == nil {
			uploads = append(uploads, upload)
		}
	}

	return uploads, nil
}

func (s *Service) SetRecipePictures(
	ctx context.Context,
	recipeId,
	userId uuid.UUID,
	pictures entity.RecipePictures,
	version *int32,
	subscriptionPlan string,
) (int32, entity.RecipePictures, error) {
	if policy, err := s.recipeRepo.GetRecipePolicy(ctx, recipeId); err != nil || policy.OwnerId != userId {
		return 0, entity.RecipePictures{}, fail.GrpcAccessDenied
	}

	validatedPictures, pictureIds := s.validatePictureLinks(recipeId, pictures)

	maxPicturesCount := s.subscriptionLimiter.GetMaxPicturesCount(subscriptionPlan)
	if len(pictureIds) > maxPicturesCount {
		return 0, entity.RecipePictures{}, recipeFail.GrpcRecipePicturesCountLimit
	}

	if !s.s3.CheckRecipePicturesExist(ctx, recipeId, pictureIds) {
		return 0, entity.RecipePictures{}, recipeFail.GrpcRecipePictureNotFound
	}

	newVersion, err := s.recipeRepo.SetRecipePictures(ctx, recipeId, validatedPictures, pictureIds, version)
	if err != nil {
		return 0, entity.RecipePictures{}, err
	}

	go func() {
		s.s3.DeleteUnusedRecipePictures(context.WithoutCancel(ctx), recipeId, pictureIds)
	}()

	return newVersion, validatedPictures, nil
}

func (s *Service) validatePictureLinks(recipeId uuid.UUID, pictures entity.RecipePictures) (entity.RecipePictures, []uuid.UUID) {
	validatedPictures := pictures
	var pictureIds []uuid.UUID
	if validatedPictures.Preview != nil {
		if pictureId := s.s3.GetRecipePictureIdByLink(recipeId, *validatedPictures.Preview); pictureId != nil {
			pictureIds = append(pictureIds, *pictureId)
		} else {
			validatedPictures.Preview = nil
		}
	}
	for stepId, stepPictures := range validatedPictures.Cooking {
		validatedStepPictures := stepPictures
		for _, stepPicture := range stepPictures {
			if pictureId := s.s3.GetRecipePictureIdByLink(recipeId, stepPicture); pictureId != nil {
				pictureIds = append(pictureIds, *pictureId)
			} else {
				var filteredPictures []string
				for _, filteredPicture := range validatedStepPictures {
					if filteredPicture != stepPicture {
						filteredPictures = append(filteredPictures, filteredPicture)
					}
				}
				validatedStepPictures = filteredPictures
			}
		}
		if len(validatedStepPictures) > 0 {
			validatedPictures.Cooking[stepId] = validatedStepPictures
		} else {
			validatedPictures.Cooking[stepId] = nil
		}
	}
	return validatedPictures, pictureIds
}
