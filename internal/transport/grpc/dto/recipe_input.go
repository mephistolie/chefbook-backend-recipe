package dto

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	sliceUtils "github.com/mephistolie/chefbook-backend-common/utils/slices"
	"github.com/mephistolie/chefbook-backend-recipe/api/model"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	recipeFail "github.com/mephistolie/chefbook-backend-recipe/internal/entity/fail"
	"k8s.io/utils/strings/slices"
)

const (
	maxNameLength        = 75
	maxDescriptionLength = 1000

	maxRecipeTagsCount = 10

	maxServings    = 1000
	maxCookingTime = 10080 // 1 week
	maxCalories    = 10000
)

func NewRecipeInput(
	req *api.RecipeInput,
	isUpdateInput bool,
	isEncryptedRecipeAllowed bool,
) (entity.RecipeInput, error) {
	if len(req.Name) == 0 {
		return entity.RecipeInput{}, fail.GrpcInvalidBody
	}
	if req.IsEncrypted && req.Visibility == model.VisibilityPublic {
		return entity.RecipeInput{}, recipeFail.GrpcEncryptedPublicRecipe
	}
	if req.IsEncrypted && !isEncryptedRecipeAllowed {
		return entity.RecipeInput{}, fail.GrpcPremiumRequired
	}

	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return entity.RecipeInput{}, fail.GrpcInvalidBody
	}
	var recipeIdPtr *uuid.UUID
	if req.RecipeId != nil {
		if recipeId, err := uuid.Parse(*req.RecipeId); err == nil {
			recipeIdPtr = &recipeId
		}
	}
	if recipeIdPtr == nil && isUpdateInput {
		return entity.RecipeInput{}, fail.GrpcInvalidBody
	}
	if len([]rune(req.Name)) > maxNameLength {
		req.Name = string([]rune(req.Name)[0:maxNameLength])
	}
	if !slices.Contains(model.AvailableVisibilities, req.Visibility) {
		req.Visibility = model.VisibilityPrivate
	}
	if req.Description != nil && len([]rune(*req.Description)) > maxDescriptionLength {
		description := string([]rune(*req.Description)[0:maxDescriptionLength])
		req.Description = &description
	}
	req.Tags = sliceUtils.RemoveDuplicates(req.Tags)
	if len(req.Tags) > maxRecipeTagsCount {
		req.Tags = req.Tags[0:maxRecipeTagsCount]
	}
	if req.Servings != nil {
		if *req.Servings <= 0 {
			req.Servings = nil
		} else if *req.Servings > maxServings {
			*req.Servings = maxServings
		}
	}
	if req.Time != nil {
		if *req.Time <= 0 {
			req.Time = nil
		} else if *req.Time > maxCookingTime {
			*req.Time = maxCookingTime
		}
	}
	if req.Calories != nil {
		if *req.Calories <= 0 {
			req.Calories = nil
		} else if *req.Calories > maxCalories {
			*req.Calories = maxCalories
		}
	}

	ingredients, err := newIngredients(req.Ingredients, req.IsEncrypted)
	if err != nil {
		return entity.RecipeInput{}, err
	}
	cooking, err := newCooking(req.Cooking, req.IsEncrypted)
	if err != nil {
		return entity.RecipeInput{}, err
	}

	var version *int32
	if isUpdateInput {
		version = req.Version
	}

	return entity.RecipeInput{
		RecipeId:       recipeIdPtr,
		Name:           req.Name,
		UserId:         userId,
		Visibility:     req.Visibility,
		IsEncrypted:    req.IsEncrypted,
		Language:       entity.ValidatedLanguage(req.Language),
		Description:    req.Description,
		Tags:           req.Tags,
		Servings:       req.Servings,
		Time:           req.Time,
		Calories:       req.Calories,
		Macronutrients: newMacronutrients(req.Macronutrients),
		Ingredients:    ingredients,
		Cooking:        cooking,
		Version:        version,
	}, nil
}
