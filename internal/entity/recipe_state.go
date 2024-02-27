package entity

import (
	"github.com/google/uuid"
)

type BaseRecipeState struct {
	Id      uuid.UUID
	Version int32

	OwnerId uuid.UUID

	Translations []string

	Rating float32
	Score  *int32
	Votes  int32

	Tags        []string
	Categories  []uuid.UUID
	IsFavourite bool
}

type RecipeState struct {
	BaseRecipeState
	OwnerName   *string
	OwnerAvatar *string
}

type DetailedRecipesState struct {
	Recipes           []RecipeState
	Categories        []Category
	Tags              map[string]Tag
	TagGroups         map[string]string
	HasEncryptedVault bool
}
