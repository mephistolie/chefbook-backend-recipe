package dto

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
)

type Ingredients []IngredientItem

func (i *Ingredients) Entity() []entity.IngredientItem {
	var ingredients []entity.IngredientItem
	for _, dto := range *i {
		ingredient := entity.IngredientItem{
			Id:       dto.Id,
			Text:     dto.Text,
			Type:     dto.Type,
			Amount:   dto.Amount,
			Unit:     dto.Unit,
			RecipeId: dto.RecipeId,
		}
		ingredients = append(ingredients, ingredient)
	}
	return ingredients
}

func NewIngredients(ingredients []entity.IngredientItem) Ingredients {
	dtos := make(Ingredients, len(ingredients))
	for i, ingredient := range ingredients {
		dto := IngredientItem{
			Id:       ingredient.Id,
			Text:     ingredient.Text,
			Type:     ingredient.Type,
			Amount:   ingredient.Amount,
			Unit:     ingredient.Unit,
			RecipeId: ingredient.RecipeId,
		}
		dtos[i] = dto
	}
	return dtos
}

func (i Ingredients) Value() (driver.Value, error) {
	return json.Marshal(i)
}

func (i *Ingredients) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	if err := json.Unmarshal(b, &i); err != nil {
		return errors.New("unable to unmarshal ingredients")
	}

	return nil
}

type IngredientItem struct {
	Id       uuid.UUID  `json:"id"`
	Text     *string    `json:"text,omitempty"`
	Type     string     `json:"type"`
	Amount   *int32     `json:"amount,omitempty"`
	Unit     *string    `json:"unit,omitempty"`
	RecipeId *uuid.UUID `json:"recipeId,omitempty"`
}

func (i IngredientItem) Value() (driver.Value, error) {
	return json.Marshal(i)
}

func (i *IngredientItem) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	if err := json.Unmarshal(b, &i); err != nil {
		return errors.New("unable to unmarshal recipe ingredient item")
	}

	return nil
}
