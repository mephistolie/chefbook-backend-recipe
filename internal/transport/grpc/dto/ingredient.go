package dto

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
)

func newIngredients(ingredients []*api.IngredientItem, isEncrypted bool) ([]entity.IngredientItem, error) {
	if isEncrypted && len(ingredients) != 1 {
		return nil, fail.GrpcInvalidBody
	}

	ingredientsCount := len(ingredients)
	if ingredientsCount > 50 {
		ingredientsCount = 50
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
		if ingredient == nil {
			continue
		}
		response[i] = *ingredient
	}
	return response, nil
}

func newIngredient(ingredient *api.IngredientItem, isEncrypted bool) (*entity.IngredientItem, error) {
	if isEncrypted && ingredient.Type != entity.TypeIngredientEncryptedData {
		return nil, fail.GrpcInvalidBody
	}
	ingredientId, err := uuid.Parse(ingredient.Id)
	if err != nil {
		return nil, nil
	}

	var textPtr *string
	if len(ingredient.Text) > 0 {
		text := ingredient.Text
		if len(text) > 100 {
			text = text[0:100]
		}
		textPtr = &text
	}
	var amount *int32
	if ingredient.Amount > 0 {
		amount = &ingredient.Amount
	}
	var unitPtr *string
	if len(ingredient.Unit) > 0 {
		unit := ingredient.Unit
		if len(unit) > 10 {
			unit = unit[0:10]
		}
		unitPtr = &unit
	}
	var recipeId *uuid.UUID
	if id, err := uuid.Parse(ingredient.RecipeId); err == nil {
		recipeId = &id
	}

	if textPtr == nil {
		return nil, nil
	}

	return &entity.IngredientItem{
		Id:       ingredientId,
		Text:     textPtr,
		Type:     ingredient.Type,
		Amount:   amount,
		Unit:     unitPtr,
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
	text := ""
	if ingredient.Text != nil {
		text = *ingredient.Text
	}
	var amount int32 = 0
	if ingredient.Amount != nil {
		amount = *ingredient.Amount
	}
	unit := ""
	if ingredient.Unit != nil {
		unit = *ingredient.Unit
	}
	recipeId := ""
	if ingredient.RecipeId != nil {
		recipeId = ingredient.RecipeId.String()
	}

	return &api.IngredientItem{
		Id:       ingredient.Id.String(),
		Text:     text,
		Type:     ingredient.Type,
		Amount:   amount,
		Unit:     unit,
		RecipeId: recipeId,
	}
}
