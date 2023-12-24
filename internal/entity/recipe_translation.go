package entity

import "github.com/google/uuid"

type RecipeTranslationInfo struct {
	AuthorId     uuid.UUID
	AuthorName   *string
	AuthorAvatar *string
}

type RecipeTranslation struct {
	AuthorId    uuid.UUID
	Language    string
	Name        string
	Description *string
	Ingredients map[uuid.UUID]IngredientTranslation
	Cooking     map[uuid.UUID]string
}

type IngredientTranslation struct {
	Text string
	Unit *string
}
