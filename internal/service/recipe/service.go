package recipe

import (
	"github.com/mephistolie/chefbook-backend-recipe/internal/repository/grpc"
	"github.com/mephistolie/chefbook-backend-recipe/internal/service/dependencies/repository"
)

type Service struct {
	repo repository.Recipe
	grpc *grpc.Repository
}

func NewService(
	repo repository.Recipe,
	grpc *grpc.Repository,
) *Service {
	return &Service{
		repo: repo,
		grpc: grpc,
	}
}
