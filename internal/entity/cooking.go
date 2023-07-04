package entity

import (
	"github.com/google/uuid"
)

const (
	TypeStep                 = "step"
	TypeCookingSection       = "section"
	TypeCookingEncryptedData = "encrypted_data"
)

var AvailableCookingTypes = []string{
	TypeStep,
	TypeCookingSection,
	TypeCookingEncryptedData,
}

type CookingItem struct {
	Id         uuid.UUID
	Text       *string
	Type       string
	Time       *int32
	PictureIds []uuid.UUID
	Pictures   []string
	RecipeId   *uuid.UUID
}
