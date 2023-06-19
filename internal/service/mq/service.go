package mq

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/firebase"
	"github.com/mephistolie/chefbook-backend-recipe/internal/repository/grpc"
	"github.com/mephistolie/chefbook-backend-recipe/internal/service/dependencies/repository"
)

type Service struct {
	repo     repository.Recipe
	grpc     *grpc.Repository
	firebase *firebase.Client
}

func NewService(
	repo repository.Recipe,
	grpc *grpc.Repository,
	firebase *firebase.Client,
) *Service {
	return &Service{
		repo:     repo,
		grpc:     grpc,
		firebase: firebase,
	}
}

func (s *Service) DeleteUserRecipes(userId uuid.UUID, deleteSharedData bool, messageId uuid.UUID) error {
	return s.repo.DeleteUserRecipes(userId, deleteSharedData, messageId)
}
