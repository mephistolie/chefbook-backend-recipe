package repository

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/mq/model"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
)

type Recipe interface {
	GetRecipes(params entity.RecipesQuery, userId uuid.UUID) []entity.BaseRecipeInfo
	GetRandomRecipe(userId uuid.UUID, languages *[]string) (entity.BaseRecipe, error)
	GetRecipeBook(userId uuid.UUID) ([]entity.BaseRecipeState, error)
	GetRecipeNames(recipeIds []uuid.UUID, userId uuid.UUID) (map[uuid.UUID]string, error)

	CreateRecipe(input entity.RecipeInput) (uuid.UUID, int32, error)
	GetRecipe(recipeId, userId uuid.UUID) (entity.BaseRecipe, error)
	UpdateRecipe(input entity.RecipeInput) (int32, error)
	SetRecipeTags(recipeId uuid.UUID, tags []string) error
	DeleteRecipe(recipeId uuid.UUID) (*model.MessageData, error)

	GetRecipePictureIdsToUpload(recipeId uuid.UUID, picturesCount int) ([]uuid.UUID, error)
	SetRecipePictures(recipeId uuid.UUID, pictures entity.RecipePictures, pictureIds []uuid.UUID, version *int32) (int32, error)

	GetRecipeRatingAndVotes(recipeId uuid.UUID) (float32, int, error)
	RateRecipe(recipeId, userId uuid.UUID, score int) (*model.MessageData, error)

	SaveToRecipeBook(recipeId, userId uuid.UUID) error
	RemoveFromRecipeBook(recipeId, userId uuid.UUID) error
	SetRecipeFavouriteStatus(recipeId, userId uuid.UUID, isFavourite bool) error
	SetRecipeCategories(recipeId, userId uuid.UUID, categories []uuid.UUID) error

	GetRecipeTranslation(recipeId uuid.UUID, language string, authorId *uuid.UUID) *entity.RecipeTranslation
	TranslateRecipe(recipeId uuid.UUID, translation entity.RecipeTranslation) error
	DeleteRecipeTranslation(recipeId uuid.UUID, userId uuid.UUID, language string) error

	GetRecipePolicy(recipeId uuid.UUID) (entity.RecipePolicy, error)

	ConfirmFirebaseDataLoad(messageId uuid.UUID) error
	DeleteUserEncryptedRecipes(userId uuid.UUID, messageId uuid.UUID) error
	DeleteUserRecipes(userId uuid.UUID, deleteSharedData bool, messageId uuid.UUID) error
}
