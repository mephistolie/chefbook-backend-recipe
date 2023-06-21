package dto

import (
	"github.com/google/uuid"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	recipeFail "github.com/mephistolie/chefbook-backend-recipe/internal/entity/fail"
	"k8s.io/utils/strings/slices"
)

const (
	maxIngredientsCount      = 50
	maxIngredientTextLength  = 150
	maxIngredientAmount      = 10000
	maxIngredientUnitLength  = 10
	encryptedIngredientsSize = 1
)

func newIngredients(ingredients []*api.IngredientItem, isEncrypted bool) ([]entity.IngredientItem, error) {
	if len(ingredients) == 0 {
		return nil, recipeFail.GrpcEmptyIngredients
	}
	if isEncrypted && len(ingredients) != encryptedIngredientsSize {
		return nil, recipeFail.GrpcInvalidEncryptedFormat
	}

	ingredientsCount := len(ingredients)
	if ingredientsCount > maxIngredientsCount {
		ingredientsCount = maxIngredientsCount
	}

	response := make([]entity.IngredientItem, ingredientsCount)
	for i, rawIngredient := range ingredients {
		if i >= len(response) {
			break
		}

		ingredient, err := newIngredient(rawIngredient, isEncrypted)
		if err != nil {
			return nil, err
		}
		response[i] = ingredient
	}
	return response, nil
}

func newIngredient(ingredient *api.IngredientItem, isEncrypted bool) (entity.IngredientItem, error) {
	if ingredient.Text == nil {
		return entity.IngredientItem{}, recipeFail.GrpcEmptyIngredientText
	}
	if !slices.Contains(entity.AvailableIngredientTypes, ingredient.Type) {
		return entity.IngredientItem{}, recipeFail.GrpcInvalidIngredientType
	}
	if isEncrypted && ingredient.Type != entity.TypeCookingEncryptedData ||
		!isEncrypted && ingredient.Type == entity.TypeCookingEncryptedData {
		return entity.IngredientItem{}, recipeFail.GrpcInvalidEncryptedFormat
	}

	ingredientId, err := uuid.Parse(ingredient.Id)
	if err != nil {
		return entity.IngredientItem{}, recipeFail.GrpcInvalidIngredientId
	}

	if ingredient.Text != nil && len(*ingredient.Text) > maxIngredientTextLength {
		text := (*ingredient.Text)[0:maxIngredientTextLength]
		ingredient.Text = &text
	}
	if ingredient.Amount != nil && *ingredient.Amount > maxIngredientAmount {
		*ingredient.Amount = maxIngredientAmount
	}
	if ingredient.Unit != nil && len(*ingredient.Unit) > maxIngredientUnitLength {
		unit := (*ingredient.Unit)[0:maxIngredientUnitLength]
		ingredient.Unit = &unit
	}
	var recipeId *uuid.UUID
	if ingredient.RecipeId != nil {
		if id, err := uuid.Parse(*ingredient.RecipeId); err == nil {
			recipeId = &id
		}
	}

	return entity.IngredientItem{
		Id:       ingredientId,
		Text:     ingredient.Text,
		Type:     ingredient.Type,
		Amount:   ingredient.Amount,
		Unit:     ingredient.Unit,
		RecipeId: recipeId,
	}, nil
}

func newIngredientsResponse(ingredients []entity.IngredientItem) []*api.IngredientItem {
	response := make([]*api.IngredientItem, len(ingredients))
	for i, ingredient := range ingredients {
		response[i] = newIngredientResponse(ingredient)
	}
	return response
}

func newIngredientResponse(ingredient entity.IngredientItem) *api.IngredientItem {
	var recipeIdPtr *string
	if ingredient.RecipeId != nil {
		recipeId := ingredient.RecipeId.String()
		recipeIdPtr = &recipeId
	}

	return &api.IngredientItem{
		Id:       ingredient.Id.String(),
		Text:     ingredient.Text,
		Type:     ingredient.Type,
		Amount:   ingredient.Amount,
		Unit:     ingredient.Unit,
		RecipeId: recipeIdPtr,
	}
}