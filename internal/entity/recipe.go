package entity

import (
	"github.com/google/uuid"
	"time"
)

type BaseRecipe struct {
	Id   uuid.UUID
	Name string

	OwnerId uuid.UUID

	IsOwned     bool
	IsSaved     bool
	Visibility  string
	IsEncrypted bool

	Language    string
	Description *string
	Preview     *string

	CreationTimestamp time.Time
	UpdateTimestamp   time.Time
	Version           int32

	Rating float32
	Score  *int32
	Votes  int32

	Tags        []string
	Categories  []uuid.UUID
	IsFavourite bool

	Servings *int32
	Time     *int32

	Calories       *int32
	Macronutrients *Macronutrients

	Ingredients []IngredientItem
	Cooking     []CookingItem
}

type Recipe struct {
	BaseRecipe

	OwnerName   *string
	OwnerAvatar *string
}

type DetailedRecipe struct {
	Recipe     Recipe
	Tags       map[string]Tag
	Categories map[string]Category
}
