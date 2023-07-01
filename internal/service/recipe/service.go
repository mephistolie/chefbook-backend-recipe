package recipe

import (
	amqp "github.com/mephistolie/chefbook-backend-common/mq/publisher"
	"github.com/mephistolie/chefbook-backend-recipe/internal/repository/grpc"
	"github.com/mephistolie/chefbook-backend-recipe/internal/service/dependencies/repository"
)

type Service struct {
	repo        repository.Recipe
	grpc        *grpc.Repository
	mqPublisher *amqp.Publisher
}

func NewService(
	repo repository.Recipe,
	grpc *grpc.Repository,
	mqPublisher *amqp.Publisher,
) *Service {
	return &Service{
		repo:        repo,
		grpc:        grpc,
		mqPublisher: mqPublisher,
	}
}
