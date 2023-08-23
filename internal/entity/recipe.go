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
	PreviewId   *uuid.UUID

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
	PictureIds  RecipePictureIds
}

type Recipe struct {
	BaseRecipe

	OwnerName   *string
	OwnerAvatar *string

	Pictures RecipePictures
}

type DetailedRecipe struct {
	Recipe     Recipe
	Tags       map[string]Tag
	Categories map[string]Category
}
