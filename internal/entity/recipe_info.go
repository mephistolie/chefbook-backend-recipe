package entity

import (
	"github.com/google/uuid"
	"time"
)

type BaseRecipeInfo struct {
	Id   uuid.UUID
	Name string

	OwnerId uuid.UUID

	IsOwned     bool
	IsSaved     bool
	Visibility  string
	IsEncrypted bool

	Language  string
	PreviewId *uuid.UUID

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

	Calories *int32
}

type RecipeInfo struct {
	BaseRecipeInfo

	OwnerName   *string
	OwnerAvatar *string

	Preview *string
}

type DetailedRecipesInfo struct {
	Recipes    []RecipeInfo
	Tags       map[string]Tag
	Categories map[string]Category
}
