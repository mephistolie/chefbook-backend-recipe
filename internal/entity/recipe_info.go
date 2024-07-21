package entity

import (
	"github.com/google/uuid"
	"time"
)

type RecipeInfo struct {
	Id   uuid.UUID
	Name string

	OwnerId uuid.UUID

	IsOwned     bool
	IsSaved     bool
	Visibility  string
	IsEncrypted bool

	Language     string
	Translations []string
	Preview      *string

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

	Calories *int32
}

type RecipesInfo struct {
	Recipes      []RecipeInfo
	Collections  map[uuid.UUID]CollectionInfo
	Tags         map[string]Tag
	TagGroups    map[string]string
	ProfilesInfo map[string]ProfileInfo
}
