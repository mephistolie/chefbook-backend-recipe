package recipe

import (
	"context"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	api "github.com/mephistolie/chefbook-backend-tag/api/proto/implementation/v1"
)

func (s *Service) CreateRecipe(input entity.RecipeInput) (uuid.UUID, int32, error) {
	id, version, err := s.repo.CreateRecipe(input)
	if err == nil {
		go s.ValidateTags(input)
	}
	return id, version, err
}

func (s *Service) GetRecipe(recipeId, userId uuid.UUID, language string) (entity.DetailedRecipe, error) {
	baseRecipe, err := s.repo.GetRecipe(recipeId, userId)
	if err != nil {
		return entity.DetailedRecipe{}, err
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
		if len(info.VisibleName) > 0 {
			recipe.OwnerName = &info.VisibleName
		}
		if len(info.Avatar) > 0 {
			recipe.OwnerAvatar = &info.Avatar
		}
	}

	wg.Wait()

	return entity.DetailedRecipe{
		Recipe:     recipe,
		Tags:       tags,
		Categories: categories,
	}
}

func (s *Service) UpdateRecipe(input entity.RecipeInput) (int32, error) {
	version, err := s.repo.UpdateRecipe(input)
	if err == nil {
		go s.ValidateTags(input)
	}
	return version, err
}

func (s *Service) DeleteRecipe(recipeId, userId uuid.UUID) error {
	return s.repo.DeleteRecipe(recipeId, userId)
}

func (s *Service) ValidateTags(input entity.RecipeInput) {
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
			input.Tags = usedTags
			_, _ = s.repo.UpdateRecipe(input)
		}
	}
}
