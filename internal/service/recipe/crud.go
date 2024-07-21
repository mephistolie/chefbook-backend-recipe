package recipe

import (
	"context"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"github.com/mephistolie/chefbook-backend-recipe/api/model"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	recipeFail "github.com/mephistolie/chefbook-backend-recipe/internal/entity/fail"
	api "github.com/mephistolie/chefbook-backend-tag/api/proto/implementation/v1"
	"sync"
)

func (s *Service) CreateRecipe(input entity.RecipeInput) (uuid.UUID, int32, error) {
	id, version, err := s.recipeRepo.CreateRecipe(input)
	if err == nil {
		go s.validateTags(input)
	}
	return id, version, err
}

func (s *Service) GetRecipe(recipeId, userId uuid.UUID, language string, translatorId *uuid.UUID) (entity.DetailedRecipe, error) {
	baseRecipe, err := s.recipeRepo.GetRecipe(recipeId, userId)
	if err != nil {
		return entity.DetailedRecipe{}, err
	}
	if baseRecipe.OwnerId != userId && baseRecipe.Visibility == model.VisibilityPrivate {
		return entity.DetailedRecipe{}, fail.GrpcAccessDenied
	}
	return s.fillBaseRecipe(baseRecipe, language, translatorId), nil
}

func (s *Service) GetRandomRecipe(userId uuid.UUID, recipeLanguages *[]string, userLanguage string) (entity.DetailedRecipe, error) {
	baseRecipe, err := s.recipeRepo.GetRandomRecipe(userId, recipeLanguages)
	if err != nil {
		return entity.DetailedRecipe{}, err
	}
	return s.fillBaseRecipe(baseRecipe, userLanguage, nil), nil
}

func (s *Service) fillBaseRecipe(recipe entity.Recipe, language string, translatorId *uuid.UUID) entity.DetailedRecipe {
	wg := sync.WaitGroup{}
	wg.Add(3)

	var tags map[string]entity.Tag
	var tagGroups map[string]string
	go func() {
		tags, tagGroups = s.getTags(recipe.Tags, language)
		wg.Done()
	}()

	var collections map[uuid.UUID]entity.CollectionInfo
	go func() {
		collections = s.getCollectionsMap(recipe.Collections)
		wg.Done()
	}()

	var profilesInfo map[string]entity.ProfileInfo
	go func() {
		profilesInfo = s.getRecipeProfilesInfo(recipe)
		wg.Done()
	}()

	s.fillRecipeTranslation(&recipe, language, translatorId)

	wg.Wait()

	return entity.DetailedRecipe{
		Recipe:       recipe,
		Tags:         tags,
		TagGroups:    tagGroups,
		Collections:  collections,
		ProfilesInfo: profilesInfo,
	}
}

func (s *Service) fillRecipeTranslation(recipe *entity.Recipe, language string, translatorId *uuid.UUID) {
	if recipe.Language == language {
		return
	}
	if _, ok := recipe.Translations[language]; !ok {
		return
	}

	translation := s.recipeRepo.GetRecipeTranslation(recipe.Id, language, translatorId)
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

func (s *Service) UpdateRecipe(input entity.RecipeInput) (int32, error) {
	policy, err := s.recipeRepo.GetRecipePolicy(*input.RecipeId)
	if err != nil {
		return 0, err
	}
	if policy.OwnerId != input.UserId {
		return 0, fail.GrpcAccessDenied
	}
	if policy.IsEncrypted != input.IsEncrypted {
		return 0, recipeFail.GrpcChangedEncryptionStatus
	}

	version, err := s.recipeRepo.UpdateRecipe(input)
	if err == nil {
		go s.validateTags(input)
	}
	return version, err
}

func (s *Service) DeleteRecipe(recipeId, userId uuid.UUID) error {
	policy, err := s.recipeRepo.GetRecipePolicy(recipeId)
	if err != nil {
		return err
	}
	if policy.OwnerId != userId {
		return fail.GrpcAccessDenied
	}
	msg, err := s.recipeRepo.DeleteRecipe(recipeId)

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
			_ = s.recipeRepo.SetRecipeTags(*input.RecipeId, usedTags)
		}
	}
}
