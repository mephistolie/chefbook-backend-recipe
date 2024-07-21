package recipe

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"github.com/mephistolie/chefbook-backend-recipe/api/model"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	recipeFail "github.com/mephistolie/chefbook-backend-recipe/internal/entity/fail"
)

func (s *Service) TranslateRecipe(recipeId uuid.UUID, translation entity.RecipeTranslation) error {
	if err := s.validateTranslation(recipeId, translation); err != nil {
		return err
	}

	return s.recipeRepo.TranslateRecipe(recipeId, translation)
}

func (s *Service) validateTranslation(recipeId uuid.UUID, translation entity.RecipeTranslation) error {
	recipe, err := s.recipeRepo.GetRecipe(recipeId, translation.AuthorId)
	if err != nil {
		return err
	}
	if recipe.Visibility != model.VisibilityPublic || recipe.Language == translation.Language {
		return fail.GrpcAccessDenied
	}

	if recipe.Description == nil && translation.Description != nil {
		translation.Description = nil
	}
	if len(recipe.Ingredients) != len(translation.Ingredients) || len(recipe.Cooking) != len(translation.Cooking) {
		return recipeFail.GrpcTranslationMismatch
	}

	approvedIngredientsCount := 0
	for _, ingredient := range recipe.Ingredients {
		if _, ok := translation.Ingredients[ingredient.Id]; ok {
			approvedIngredientsCount += 1
		}
	}
	if approvedIngredientsCount < len(recipe.Ingredients) {
		return recipeFail.GrpcTranslationMismatch
	}
	approvedCookingCount := 0
	for _, cooking := range recipe.Cooking {
		if _, ok := translation.Cooking[cooking.Id]; ok {
			approvedCookingCount += 1
		}
	}
	if approvedCookingCount < len(recipe.Cooking) {
		return recipeFail.GrpcTranslationMismatch
	}

	return nil
}

func (s *Service) DeleteRecipeTranslation(recipeId, userId uuid.UUID, language string) error {
	return s.recipeRepo.DeleteRecipeTranslation(recipeId, userId, language)
}
