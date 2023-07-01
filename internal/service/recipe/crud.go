package recipe

import (
	"context"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"github.com/mephistolie/chefbook-backend-recipe/api/model"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	recipeFail "github.com/mephistolie/chefbook-backend-recipe/internal/entity/fail"
	api "github.com/mephistolie/chefbook-backend-tag/api/proto/implementation/v1"
)

func (s *Service) CreateRecipe(input entity.RecipeInput) (uuid.UUID, int32, error) {
	id, version, err := s.repo.CreateRecipe(input)
	if err == nil {
		go s.validateTags(input)
	}
	return id, version, err
}

func (s *Service) GetRecipe(recipeId, userId uuid.UUID, language string) (entity.DetailedRecipe, error) {
	baseRecipe, err := s.repo.GetRecipe(recipeId, userId)
	if err != nil {
		return entity.DetailedRecipe{}, err
	}
	if baseRecipe.OwnerId != userId && baseRecipe.Visibility == model.VisibilityPrivate {
		return entity.DetailedRecipe{}, fail.GrpcAccessDenied
	}
	return s.fillBaseRecipe(baseRecipe, userId, language), nil
}

func (s *Service) GetRandomRecipe(userId uuid.UUID, recipeLanguages *[]string, userLanguage string) (entity.DetailedRecipe, error) {
	baseRecipe, err := s.repo.GetRandomRecipe(userId, recipeLanguages)
	if err != nil {
		return entity.DetailedRecipe{}, err
	}
	return s.fillBaseRecipe(baseRecipe, userId, userLanguage), nil
}

func (s *Service) fillBaseRecipe(baseRecipe entity.BaseRecipe, userId uuid.UUID, language string) entity.DetailedRecipe {
	recipe := entity.Recipe{BaseRecipe: baseRecipe}

	tags := make(map[string]entity.Tag)
	categories := make(map[string]entity.Category)
	wg := s.getCategoriesAndTagsAsync(baseRecipe.Tags, baseRecipe.Categories, userId, language, &tags, &categories)

	ownerId := baseRecipe.OwnerId.String()
	authorsMap := s.getRecipeAuthorsInfo([]string{ownerId})

	if info, ok := authorsMap[ownerId]; ok && info != nil {
		recipe.OwnerName = info.VisibleName
		recipe.OwnerAvatar = info.Avatar
	}

	wg.Wait()

	return entity.DetailedRecipe{
		Recipe:     recipe,
		Tags:       tags,
		Categories: categories,
	}
}

func (s *Service) UpdateRecipe(input entity.RecipeInput) (int32, error) {
	policy, err := s.repo.GetRecipePolicy(*input.RecipeId)
	if err != nil {
		return 0, err
	}
	if policy.OwnerId != input.UserId {
		return 0, fail.GrpcAccessDenied
	}
	if policy.IsEncrypted != input.IsEncrypted {
		return 0, recipeFail.GrpcChangedEncryptionStatus
	}

	version, err := s.repo.UpdateRecipe(input)
	if err == nil {
		go s.validateTags(input)
	}
	return version, err
}

func (s *Service) DeleteRecipe(recipeId, userId uuid.UUID) error {
	policy, err := s.repo.GetRecipePolicy(recipeId)
	if err != nil {
		return err
	}
	if policy.OwnerId != userId {
		return fail.GrpcAccessDenied
	}
	msg, err := s.repo.DeleteRecipe(recipeId)

	if err == nil {
		go s.mqPublisher.PublishMessage(msg)
	}

	return err
}

func (s *Service) validateTags(input entity.RecipeInput) {
	if len(input.Tags) == 0 {
		return
	}

	if res, err := s.grpc.Tag.GetTags(context.Background(), &api.GetTagsRequest{}); err == nil {
		var existingTags map[string]bool
		for _, tag := range res.Tags {
			existingTags[tag.TagId] = true
		}

		var usedTags []string
		for _, tag := range input.Tags {
			if exists, ok := existingTags[tag]; ok && exists {
				usedTags = append(usedTags, tag)
			}
		}
		if len(usedTags) < len(input.Tags) {
			_ = s.repo.SetRecipeTags(*input.RecipeId, usedTags)
		}
	}
}
