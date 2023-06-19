package dto

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	"k8s.io/utils/strings/slices"
	"regexp"
)

const languageRegex = "^[a-z]+$"

func NewRecipeCreationInput(req *api.CreateRecipeRequest) (entity.RecipeInput, error) {
	return newRecipeInput(
		req.UserId,
		req.RecipeId,
		req.Name,
		req.Visibility,
		req.IsEncrypted,
		req.Language,
		req.Description,
		req.Tags,
		req.Servings,
		req.Time,
		req.Calories,
		req.Macronutrients,
		req.Ingredients,
		req.Cooking,
		nil,
	)
}

func NewRecipeUpdateInput(req *api.UpdateRecipeRequest) (entity.RecipeInput, error) {
	var version *int32
	if req.Version > 0 {
		version = &req.Version
	}

	return newRecipeInput(
		req.UserId,
		req.RecipeId,
		req.Name,
		req.Visibility,
		req.IsEncrypted,
		req.Language,
		req.Description,
		req.Tags,
		req.Servings,
		req.Time,
		req.Calories,
		req.Macronutrients,
		req.Ingredients,
		req.Cooking,
		version,
	)
}

func newRecipeInput(
	rawUserId string,
	rawRecipeId string,
	rawName string,
	rawVisibility string,
	isEncrypted bool,
	rawLanguage string,
	rawDescription string,
	tags []string,
	rawServings int32,
	rawTime int32,
	rawCalories int32,
	rawMacronutrients *api.Macronutrients,
	rawIngredients []*api.IngredientItem,
	rawCooking []*api.CookingItem,
	version *int32,
) (entity.RecipeInput, error) {
	if len(rawName) == 0 {
		return entity.RecipeInput{}, fail.GrpcInvalidBody
	}
	userId, err := uuid.Parse(rawUserId)
	if err != nil {
		return entity.RecipeInput{}, fail.GrpcInvalidBody
	}
	var recipeIdPtr *uuid.UUID
	if recipeId, err := uuid.Parse(rawRecipeId); err == nil {
		recipeIdPtr = &recipeId
	}
	name := rawName
	if len(name) > 150 {
		name = name[0:150]
	}
	visibility := rawVisibility
	if !slices.Contains(entity.AvailableVisibilities, visibility) {
		visibility = entity.VisibilityPrivate
	}
	language := rawLanguage
	matched, err := regexp.MatchString(languageRegex, language)
	if len(language) != 2 || !matched || err != nil {
		language = "en"
	}
	var descriptionPtr *string
	if len(rawDescription) > 0 {
		description := rawDescription
		if len(description) > 1500 {
			description = description[0:1500]
		}
		descriptionPtr = &description
	}
	var servings *int32
	if rawServings > 0 {
		servings = &rawServings
	}
	var time *int32
	if rawTime > 0 {
		time = &rawTime
	}
	var calories *int32
	if rawCalories > 0 {
		calories = &rawCalories
	}
	var macronutrientsPtr *entity.Macronutrients
	if rawMacronutrients != nil && (rawMacronutrients.Protein > 0 || rawMacronutrients.Fats > 0 || rawMacronutrients.Carbohydrates > 0) {
		macronutrients := entity.Macronutrients{}
		if rawMacronutrients.Protein > 0 {
			macronutrients.Protein = &rawMacronutrients.Protein
		}
		if rawMacronutrients.Fats > 0 {
			macronutrients.Fats = &rawMacronutrients.Fats
		}
		if rawMacronutrients.Carbohydrates > 0 {
			macronutrients.Carbohydrates = &rawMacronutrients.Carbohydrates
		}
		macronutrientsPtr = &macronutrients
	}

	ingredients, err := newIngredients(rawIngredients, isEncrypted)
	if err != nil {
		return entity.RecipeInput{}, err
	}
	cooking, err := newCooking(rawCooking, isEncrypted)
	if err != nil {
		return entity.RecipeInput{}, err
	}

	return entity.RecipeInput{
		Id:             recipeIdPtr,
		Name:           name,
		UserId:         userId,
		Visibility:     visibility,
		IsEncrypted:    isEncrypted,
		Language:       language,
		Description:    descriptionPtr,
		Tags:           tags,
		Servings:       servings,
		Time:           time,
		Calories:       calories,
		Macronutrients: macronutrientsPtr,
		Ingredients:    ingredients,
		Cooking:        cooking,
		Version:        version,
	}, nil
}
