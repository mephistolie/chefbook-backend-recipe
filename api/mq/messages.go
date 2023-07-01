package mq

import "github.com/google/uuid"

const ExchangeRecipes = "recipes"
const AppId = "recipe-service"

const (
	MsgTypeRecipeRatingChanged = "recipe.rating.changed"
	MsgTypeRecipeDeleted       = "recipe.deleted"
)

type MsgBodyRecipeRatingChanged struct {
	RecipeId  uuid.UUID `json:"recipeId"`
	OwnerId   uuid.UUID `json:"ownerId"`
	ScoreDiff int       `json:"scoreDiff"`
	UserId    uuid.UUID `json:"userId"`
}

type MsgBodyRecipeDeleted struct {
	RecipeId uuid.UUID `json:"recipeId"`
}
