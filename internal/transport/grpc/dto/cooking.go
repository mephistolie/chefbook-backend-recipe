package dto

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
)

func newCooking(cooking []*api.CookingItem, isEncrypted bool) ([]entity.CookingItem, error) {
	if isEncrypted && len(cooking) != 1 {
		return nil, fail.GrpcInvalidBody
	}

	ingredientsCount := len(cooking)
	if ingredientsCount > 30 {
		ingredientsCount = 30
	}

	response := make([]entity.CookingItem, ingredientsCount)
	for i, rawItem := range cooking {
		if i >= len(response) {
			break
		}

		item, err := newCookingItem(rawItem, isEncrypted)
		if err != nil {
			return nil, err
		}
		if item == nil {
			continue
		}
		response[i] = *item
	}
	return response, nil
}

func newCookingItem(item *api.CookingItem, isEncrypted bool) (*entity.CookingItem, error) {
	if isEncrypted && item.Type != entity.TypeCookingEncryptedData {
		return nil, fail.GrpcInvalidBody
	}
	itemId, err := uuid.Parse(item.Id)
	if err != nil {
		return nil, nil
	}

	var textPtr *string
	if len(item.Text) > 0 {
		text := item.Text
		if len(text) > 2500 {
			text = text[0:2500]
		}
		textPtr = &text
	}
	var time *int32
	if item.Time > 0 {
		time = &item.Time
	}
	var recipeId *uuid.UUID
	if id, err := uuid.Parse(item.RecipeId); err == nil {
		recipeId = &id
	}

	if textPtr == nil {
		return nil, nil
	}

	return &entity.CookingItem{
		Id:       itemId,
		Text:     textPtr,
		Type:     item.Type,
		Time:     time,
		RecipeId: recipeId,
	}, nil
}

func newCookingResponse(cooking []entity.CookingItem) []*api.CookingItem {
	response := make([]*api.CookingItem, len(cooking))
	for i, cookingItem := range cooking {
		response[i] = newCookingItemResponse(cookingItem)
	}
	return response
}

func newCookingItemResponse(cookingItem entity.CookingItem) *api.CookingItem {
	text := ""
	if cookingItem.Text != nil {
		text = *cookingItem.Text
	}
	var time int32 = 0
	if cookingItem.Time != nil {
		time = *cookingItem.Time
	}
	recipeId := ""
	if cookingItem.RecipeId != nil {
		recipeId = cookingItem.RecipeId.String()
	}

	return &api.CookingItem{
		Id:       cookingItem.Id.String(),
		Text:     text,
		Type:     cookingItem.Type,
		Time:     time,
		RecipeId: recipeId,
		Pictures: cookingItem.Pictures,
	}
}
