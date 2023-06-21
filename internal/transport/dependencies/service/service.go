package service

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/firebase"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-recipe/internal/config"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	"github.com/mephistolie/chefbook-backend-recipe/internal/repository/grpc"
	"github.com/mephistolie/chefbook-backend-recipe/internal/service/dependencies/repository"
	"github.com/mephistolie/chefbook-backend-recipe/internal/service/mq"
	"github.com/mephistolie/chefbook-backend-recipe/internal/service/recipe"
)

type Service struct {
	Recipe
	MQ
}

type Recipe interface {
	GetRecipes(params entity.RecipesQuery, userId uuid.UUID, language string) entity.DetailedRecipesInfo
	GetRandomRecipe(userId uuid.UUID, recipeLanguages *[]string, userLanguage string) (entity.DetailedRecipe, error)
	GetRecipesBook(userId uuid.UUID, language string) (entity.DetailedRecipesState, error)
	GetRecipeNames(recipeIds []uuid.UUID, userId uuid.UUID) (map[uuid.UUID]string, error)
	CreateRecipe(input entity.RecipeInput) (uuid.UUID, int32, error)
	GetRecipe(recipeId, userId uuid.UUID, language string) (entity.DetailedRecipe, error)
	UpdateRecipe(input entity.RecipeInput) (int32, error)
	DeleteRecipe(recipeId, userId uuid.UUID) error

	RateRecipe(recipeId, userId uuid.UUID, score int) error
	SaveToRecipeBook(recipeId, userId uuid.UUID) error
	RemoveFromRecipeBook(recipeId, userId uuid.UUID) error
	SetRecipeFavouriteStatus(recipeId, userId uuid.UUID, favourite bool) error
	SetRecipeCategories(recipeId, userId uuid.UUID, categories []uuid.UUID) error

	GetRecipeOwner(userId uuid.UUID) (uuid.UUID, error)
}

type MQ interface {
	ImportFirebaseRecipes(userId uuid.UUID, firebaseId string, messageId uuid.UUID) error
	DeleteUserRecipes(userId uuid.UUID, deleteSharedData bool, messageId uuid.UUID) error
}

func New(
	cfg *config.Config,
	repo repository.Recipe,
	grpc *grpc.Repository,
) (*Service, error) {
	var err error = nil
	var client *firebase.Client = nil
	if len(*cfg.Firebase.Credentials) > 0 {
		credentials := []byte(*cfg.Firebase.Credentials)
		client, err = firebase.NewClient(credentials, "")
		if err != nil {
			return nil, err
		}
		log.Info("Firebase client initialized")
	}

	return &Service{
		Recipe: recipe.NewService(repo, grpc),
		MQ:     mq.NewService(repo, grpc, client),
	}, nil
}
