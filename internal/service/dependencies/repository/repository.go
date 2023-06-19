package repository

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
)

type Recipe interface {
	GetRecipes(params entity.RecipesQuery, userId uuid.UUID) []entity.BaseRecipeInfo
	GetRandomRecipe(userId uuid.UUID, languages *[]string) (entity.BaseRecipe, error)
	GetRecipeBook(userId uuid.UUID) ([]entity.BaseRecipeState, error)

	CreateRecipe(input entity.RecipeInput) (uuid.UUID, int32, error)
	GetRecipe(recipeId, userId uuid.UUID) (entity.BaseRecipe, error)
	UpdateRecipe(input entity.RecipeInput) (int32, error)
	DeleteRecipe(recipeId, userId uuid.UUID) error

	GetRecipeRatingAndVotes(recipeId uuid.UUID) (float32, int, error)
	GetUserRecipeScore(recipeId, userId uuid.UUID) (int, error)
	RateRecipe(recipeId, userId uuid.UUID, score int) error

	SaveToRecipeBook(recipeId, userId uuid.UUID) error
	RemoveFromRecipeBook(recipeId, userId uuid.UUID) error
	SetRecipeFavouriteStatus(recipeId, userId uuid.UUID, isFavourite bool) error
	SetRecipeCategories(recipeId, userId uuid.UUID, categories []uuid.UUID) error

	GetRecipeOwner(recipeId uuid.UUID) (uuid.UUID, error)

	ConfirmFirebaseDataLoad(messageId uuid.UUID) error
	DeleteUserRecipes(userId uuid.UUID, deleteSharedData bool, messageId uuid.UUID) error
}
