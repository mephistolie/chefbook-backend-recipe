package entity

import (
	"github.com/google/uuid"
	"time"
)

type RecipesQuery struct {
	PageSize int32

	RecipeIds []uuid.UUID
	AuthorId  *uuid.UUID

	Owned bool
	Saved bool

	Search *string

	Sorting               string
	LastRecipeId          *uuid.UUID
	LastCreationTimestamp *time.Time
	LastUpdateTimestamp   *time.Time
	LastRating            *float32
	LastVotes             *int32
	LastTime              *int32
	LastCalories          *int32

	MinRating *int32
	MaxRating *int32

	MinTime     *int32
	MaxTime     *int32
	MinServings *int32
	MaxServings *int32
	MinCalories *int32
	MaxCalories *int32

	Languages []string
}
