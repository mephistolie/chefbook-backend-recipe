package grpc

import (
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-recipe/internal/transport/dependencies/service"
)

type RecipeServer struct {
	api.UnsafeRecipeServiceServer
	service service.Service
}

func NewServer(service service.Service) *RecipeServer {
	return &RecipeServer{
		service: service,
	}
}
