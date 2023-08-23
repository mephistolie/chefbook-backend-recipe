package dto

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	recipeFail "github.com/mephistolie/chefbook-backend-recipe/internal/entity/fail"
	"k8s.io/utils/strings/slices"
)

const (
	maxCookingStepsCount     = 40
	maxCookingItemTextLength = 2500
	maxCookingStepTime       = 86400 // 1 day
	encryptedCookingSize     = 1
)

func newCooking(cooking []*api.CookingItem, isEncrypted bool) ([]entity.CookingItem, error) {
	if isEncrypted && len(cooking) != encryptedCookingSize {
		return nil, fail.GrpcInvalidBody
	}

	cookingItemIds := make(map[uuid.UUID]bool)
	hasSteps := false

	stepsCount := len(cooking)
	if stepsCount > maxCookingStepsCount {
		stepsCount = maxCookingStepsCount
	}

	response := make([]entity.CookingItem, stepsCount)
	for i, rawItem := range cooking {
		if i >= len(response) {
			break
		}

		item, err := newCookingItem(rawItem, isEncrypted)
		if err != nil {
			return nil, err
		}

		if exists, ok := cookingItemIds[item.Id]; ok && exists {
			return nil, recipeFail.GrpcCookingMatchingIds
		}
		cookingItemIds[item.Id] = true

		response[i] = item

		if item.Type == entity.TypeStep {
			hasSteps = true
		}
	}

	if !hasSteps {
		return nil, recipeFail.GrpcEmptyCooking
	}

	return response, nil
}

func newCookingItem(item *api.CookingItem, isEncrypted bool) (entity.CookingItem, error) {
	if item.Text == nil {
		return entity.CookingItem{}, recipeFail.GrpcEmptyCookingItemText
	}
	if !slices.Contains(entity.AvailableCookingTypes, item.Type) {
		return entity.CookingItem{}, recipeFail.GrpcInvalidCookingItemType
	}
	if isEncrypted && item.Type != entity.TypeCookingEncryptedData ||
		!isEncrypted && item.Type == entity.TypeCookingEncryptedData {
		return entity.CookingItem{}, recipeFail.GrpcInvalidEncryptedFormat
	}

	itemId, err := uuid.Parse(item.Id)
	if err != nil {
		return entity.CookingItem{}, recipeFail.GrpcInvalidCookingItemId
	}
	if item.Text != nil && len(*item.Text) > maxCookingItemTextLength && !isEncrypted {
		text := (*item.Text)[0:maxCookingItemTextLength]
		item.Text = &text
	}
	var recipeId *uuid.UUID
	if item.RecipeId != nil {
		if id, err := uuid.Parse(*item.RecipeId); err == nil {
			recipeId = &id
		}
	}
	if item.Time != nil {
		if *item.Time <= 0 {
			item.Time = nil
		} else if *item.Time > maxCookingStepTime {
			*item.Time = maxCookingStepTime
		}
	}

	if item.Type != entity.TypeStep {
		item.Time = nil
		recipeId = nil
	}

	return entity.CookingItem{
		Id:       itemId,
		Text:     item.Text,
		Type:     item.Type,
		Time:     item.Time,
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
	var recipeId *string
	if cookingItem.RecipeId != nil {
		id := cookingItem.RecipeId.String()
		recipeId = &id
	}

	return &api.CookingItem{
		Id:       cookingItem.Id.String(),
		Text:     cookingItem.Text,
		Type:     cookingItem.Type,
		Time:     cookingItem.Time,
		RecipeId: recipeId,
	}
}
