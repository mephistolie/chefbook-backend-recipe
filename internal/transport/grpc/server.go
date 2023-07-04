package grpc

import (
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-recipe/internal/helpers"
	"github.com/mephistolie/chefbook-backend-recipe/internal/transport/dependencies/service"
)

type RecipeServer struct {
	api.UnsafeRecipeServiceServer
	service             service.Recipe
	subscriptionLimiter helpers.SubscriptionLimiter
}

func NewServer(service service.Recipe, subscriptionLimiter helpers.SubscriptionLimiter) *RecipeServer {
	return &RecipeServer{
		service:             service,
		subscriptionLimiter: subscriptionLimiter,
	}
}
