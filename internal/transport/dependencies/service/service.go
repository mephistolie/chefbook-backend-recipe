package service

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/firebase"
	"github.com/mephistolie/chefbook-backend-common/log"
	mq "github.com/mephistolie/chefbook-backend-common/mq/dependencies"
	amqp "github.com/mephistolie/chefbook-backend-common/mq/publisher"
	"github.com/mephistolie/chefbook-backend-recipe/internal/config"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	"github.com/mephistolie/chefbook-backend-recipe/internal/helpers"
	"github.com/mephistolie/chefbook-backend-recipe/internal/repository/grpc"
	"github.com/mephistolie/chefbook-backend-recipe/internal/repository/s3"
	"github.com/mephistolie/chefbook-backend-recipe/internal/service/collection"
	"github.com/mephistolie/chefbook-backend-recipe/internal/service/dependencies/repository"
	mqInbox "github.com/mephistolie/chefbook-backend-recipe/internal/service/mq"
	"github.com/mephistolie/chefbook-backend-recipe/internal/service/recipe"
)

type Service struct {
	Recipe
	Collection
	MQ mq.Inbox
}

type Recipe interface {
	GetRecipes(params entity.RecipesQuery, userId uuid.UUID, language string) entity.RecipesInfo
	GetRandomRecipe(userId uuid.UUID, recipeLanguages *[]string, userLanguage string) (entity.DetailedRecipe, error)
	GetRecipesBook(userId uuid.UUID, language string) (entity.RecipeBook, error)
	GetRecipeNames(recipeIds []uuid.UUID, userId uuid.UUID) (map[uuid.UUID]string, error)
	CreateRecipe(input entity.RecipeInput) (uuid.UUID, int32, error)
	GetRecipe(recipeId, userId uuid.UUID, language string, translatorId *uuid.UUID) (entity.DetailedRecipe, error)
	UpdateRecipe(input entity.RecipeInput) (int32, error)
	DeleteRecipe(recipeId, userId uuid.UUID) error

	GenerateRecipePicturesUploadLinks(recipeId, userId uuid.UUID, picturesCount int, subscriptionPlan string) ([]entity.PictureUpload, error)
	SetRecipePictures(recipeId, userId uuid.UUID, pictures entity.RecipePictures, version *int32, subscriptionPlan string) (int32, entity.RecipePictures, error)

	RateRecipe(recipeId, userId uuid.UUID, score int) error
	SaveToRecipeBook(recipeId, userId uuid.UUID) error
	RemoveFromRecipeBook(recipeId, userId uuid.UUID) error
	SetRecipeFavouriteStatus(recipeId, userId uuid.UUID, favourite bool) error
	SetRecipeCategories(recipeId, userId uuid.UUID, categories []uuid.UUID) error

	TranslateRecipe(recipeId uuid.UUID, translation entity.RecipeTranslation) error
	DeleteRecipeTranslation(recipeId, userId uuid.UUID, language string) error

	GetRecipePolicy(userId uuid.UUID) (entity.RecipePolicy, error)
}

type Collection interface {
	GetCollections(userId uuid.UUID, requesterId uuid.UUID) entity.DetailedCollections
	CreateCollection(input entity.CollectionInput) (uuid.UUID, error)
	GetCollection(collectionId uuid.UUID, userId uuid.UUID) (entity.DetailedCollection, error)
	UpdateCollection(input entity.CollectionInput) error
	DeleteCollection(collectionId, userId uuid.UUID) error
}

func New(
	cfg *config.Config,
	recipeRepo repository.Recipe,
	collectionRepo repository.Collection,
	mqRepo repository.MQ,
	grpc *grpc.Repository,
	s3 *s3.Repository,
	mqPublisher *amqp.Publisher,
	subscriptionLimiter helpers.SubscriptionLimiter,
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
		Recipe:     recipe.NewService(recipeRepo, collectionRepo, grpc, s3, mqPublisher, subscriptionLimiter),
		Collection: collection.NewService(collectionRepo, grpc),
		MQ:         mqInbox.NewService(mqRepo, recipeRepo, collectionRepo, client),
	}, nil
}
