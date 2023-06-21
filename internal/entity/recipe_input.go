package entity

import (
	"github.com/google/uuid"
	"time"
)

type RecipeInput struct {
	RecipeId *uuid.UUID
	Name     string

	UserId uuid.UUID

	Visibility  string
	IsEncrypted bool

	Language    string
	Description *string

	Version *int32

	Tags []string

	Servings *int32
	Time     *int32

	Calories       *int32
	Macronutrients *Macronutrients

	Ingredients []IngredientItem
	Cooking     []CookingItem

	CreationTimestamp *time.Time
}
