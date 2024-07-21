package entity

import (
	"github.com/google/uuid"
	"time"
)

type Recipe struct {
	Id   uuid.UUID
	Name string

	OwnerId uuid.UUID

	IsOwned     bool
	IsSaved     bool
	Visibility  string
	IsEncrypted bool

	Language     string
	Translations map[string][]uuid.UUID
	Description  *string

	CreationTimestamp time.Time
	UpdateTimestamp   time.Time
	Version           int32

	Rating float32
	Score  *int32
	Votes  int32

	Tags        []string
	Collections []uuid.UUID
	IsFavourite bool

	Servings *int32
	Time     *int32

	Calories       *int32
	Macronutrients *Macronutrients

	Ingredients []IngredientItem
	Cooking     []CookingItem
	Pictures    RecipePictures
}

type DetailedRecipe struct {
	Recipe       Recipe
	Tags         map[string]Tag
	TagGroups    map[string]string
	Collections  map[uuid.UUID]CollectionInfo
	ProfilesInfo map[string]ProfileInfo
}
