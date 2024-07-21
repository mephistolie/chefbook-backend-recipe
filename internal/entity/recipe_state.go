package entity

import (
	"github.com/google/uuid"
)

type RecipeState struct {
	Id      uuid.UUID
	Version int32

	OwnerId uuid.UUID

	Translations map[string][]uuid.UUID

	Rating float32
	Score  *int32
	Votes  int32

	Tags        []string
	Collections []uuid.UUID
	IsFavourite bool
}

type RecipeBook struct {
	Recipes           []RecipeState
	Collections       []Collection
	Tags              map[string]Tag
	TagGroups         map[string]string
	HasEncryptedVault bool
	ProfilesInfo      map[string]ProfileInfo
}
