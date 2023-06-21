package entity

import (
	"github.com/google/uuid"
)

type RecipePolicy struct {
	Id          uuid.UUID
	OwnerId     uuid.UUID
	Visibility  string
	IsEncrypted bool
}
