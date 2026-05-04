package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/mq/model"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
)

type Recipe interface {
	GetRecipes(ctx context.Context, params entity.RecipesQuery, userId uuid.UUID) []entity.RecipeInfo
	GetRandomRecipe(ctx context.Context, userId uuid.UUID, languages *[]string) (entity.Recipe, error)
	GetRecipeBook(ctx context.Context, userId uuid.UUID) ([]entity.RecipeState, error)
	GetRecipeNames(ctx context.Context, recipeIds []uuid.UUID, userId uuid.UUID) (map[uuid.UUID]string, error)

	CreateRecipe(ctx context.Context, input entity.RecipeInput) (uuid.UUID, int32, error)
	GetRecipe(ctx context.Context, recipeId, userId uuid.UUID) (entity.Recipe, error)
	UpdateRecipe(ctx context.Context, input entity.RecipeInput) (int32, error)
	SetRecipeTags(ctx context.Context, recipeId uuid.UUID, tags []string) error
	DeleteRecipe(ctx context.Context, recipeId uuid.UUID) (*model.MessageData, error)

	GetRecipePictureIdsToUpload(ctx context.Context, recipeId uuid.UUID, picturesCount int) ([]uuid.UUID, error)
	SetRecipePictures(ctx context.Context, recipeId uuid.UUID, pictures entity.RecipePictures, pictureIds []uuid.UUID, version *int32) (int32, error)

	GetRecipeRatingAndVotes(ctx context.Context, recipeId uuid.UUID) (float32, int, error)
	RateRecipe(ctx context.Context, recipeId, userId uuid.UUID, score int) (*model.MessageData, error)

	SaveRecipeToRecipeBook(ctx context.Context, recipeId, userId uuid.UUID) error
	RemoveRecipeFromRecipeBook(ctx context.Context, recipeId, userId uuid.UUID) error
	SaveRecipeToFavourites(ctx context.Context, recipeId, userId uuid.UUID) error
	RemoveRecipeFromFavourites(ctx context.Context, recipeId, userId uuid.UUID) error
	AddRecipeToCollection(ctx context.Context, recipeId, collectionId, userId uuid.UUID) error
	RemoveRecipeFromCollection(ctx context.Context, recipeId, collectionId, userId uuid.UUID) error
	SetRecipeCollections(ctx context.Context, recipeId, userId uuid.UUID, collections []uuid.UUID) error

	GetRecipeTranslation(ctx context.Context, recipeId uuid.UUID, language string, authorId *uuid.UUID) *entity.RecipeTranslation
	TranslateRecipe(ctx context.Context, recipeId uuid.UUID, translation entity.RecipeTranslation) error
	DeleteRecipeTranslation(ctx context.Context, recipeId uuid.UUID, userId uuid.UUID, language string) error

	GetRecipePolicy(ctx context.Context, recipeId uuid.UUID) (entity.RecipePolicy, error)
}
