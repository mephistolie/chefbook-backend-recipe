package entity

import (
	"github.com/google/uuid"
)

const (
	TypeIngredient              = "ingredient"
	TypeIngredientsSection      = "section"
	TypeIngredientEncryptedData = "encrypted_data"
)

var AvailableIngredientTypes = []string{
	TypeIngredient,
	TypeIngredientsSection,
	TypeIngredientEncryptedData,
}

type IngredientItem struct {
	Id       uuid.UUID
	Text     *string
	Type     string
	Amount   *float32
	Unit     *string
	RecipeId *uuid.UUID
}
