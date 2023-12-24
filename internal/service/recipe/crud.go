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

func (s *Service) GetRecipe(recipeId, userId uuid.UUID, language string, translatorId *uuid.UUID) (entity.DetailedRecipe, error) {
	baseRecipe, err := s.repo.GetRecipe(recipeId, userId)
	if err != nil {
		return entity.DetailedRecipe{}, err
	}
	if baseRecipe.OwnerId != userId && baseRecipe.Visibility == model.VisibilityPrivate {
		return entity.DetailedRecipe{}, fail.GrpcAccessDenied
	}
	return s.fillBaseRecipe(baseRecipe, userId, language, translatorId), nil
}

func (s *Service) GetRandomRecipe(userId uuid.UUID, recipeLanguages *[]string, userLanguage string) (entity.DetailedRecipe, error) {
	baseRecipe, err := s.repo.GetRandomRecipe(userId, recipeLanguages)
	if err != nil {
		return entity.DetailedRecipe{}, err
	}
	return s.fillBaseRecipe(baseRecipe, userId, userLanguage, nil), nil
}

func (s *Service) fillBaseRecipe(baseRecipe entity.BaseRecipe, userId uuid.UUID, language string, translatorId *uuid.UUID) entity.DetailedRecipe {
	recipe := entity.Recipe{BaseRecipe: baseRecipe}

	tags := make(map[string]entity.Tag)
	categories := make(map[string]entity.Category)
	wg := s.getCategoriesAndTagsAsync(baseRecipe.Tags, baseRecipe.Categories, userId, language, &tags, &categories)

	s.fillRecipeTranslation(&recipe, language, translatorId)
	s.fillProfilesData(&recipe)

	wg.Wait()

	return entity.DetailedRecipe{
		Recipe:     recipe,
		Tags:       tags,
		Categories: categories,
	}
}

func (s *Service) fillRecipeTranslation(recipe *entity.Recipe, language string, translatorId *uuid.UUID) {
	if recipe.Language == language {
		return
	}
	if _, ok := recipe.Translations[language]; !ok {
		return
	}

	translation := s.repo.GetRecipeTranslation(recipe.Id, language, translatorId)
	if translation == nil {
		return
	}

	recipe.Name = translation.Name
	if translation.Description != nil {
		recipe.Description = translation.Description
	}
	for i, ingredient := range recipe.Ingredients {
		if ingredientTranslation, ok := translation.Ingredients[ingredient.Id]; ok {
			recipe.Ingredients[i].Text = &ingredientTranslation.Text
			if ingredientTranslation.Unit != nil {
				recipe.Ingredients[i].Unit = ingredientTranslation.Unit
			}
		}
	}
	for i, step := range recipe.Cooking {
		if stepTranslation, ok := translation.Cooking[step.Id]; ok {
			recipe.Cooking[i].Text = &stepTranslation
		}
	}
}

func (s *Service) fillProfilesData(recipe *entity.Recipe) {
	profiles := []string{recipe.OwnerId.String()}
	for i, _ := range recipe.Translations {
		for j, _ := range recipe.Translations[i] {
			profiles = append(profiles, recipe.Translations[i][j].AuthorId.String())
		}
	}

	profilesMap := s.getProfilesInfo(profiles)

	if info, ok := profilesMap[recipe.OwnerId.String()]; ok && info != nil {
		recipe.OwnerName = info.VisibleName
		recipe.OwnerAvatar = info.Avatar
	}

	for i, _ := range recipe.Translations {
		for j, _ := range recipe.Translations[i] {
			if info, ok := profilesMap[recipe.Translations[i][j].AuthorId.String()]; ok && info != nil {
				recipe.Translations[i][j].AuthorName = info.VisibleName
				recipe.Translations[i][j].AuthorAvatar = info.Avatar
			}

		}
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
