package recipe

import (
	amqp "github.com/mephistolie/chefbook-backend-common/mq/publisher"
	"github.com/mephistolie/chefbook-backend-recipe/internal/helpers"
	"github.com/mephistolie/chefbook-backend-recipe/internal/repository/grpc"
	"github.com/mephistolie/chefbook-backend-recipe/internal/repository/s3"
	"github.com/mephistolie/chefbook-backend-recipe/internal/service/dependencies/repository"
)

type Service struct {
	repo                repository.Recipe
	grpc                *grpc.Repository
	s3                  *s3.Repository
	mqPublisher         *amqp.Publisher
	subscriptionLimiter helpers.SubscriptionLimiter
}

func NewService(
	repo repository.Recipe,
	grpc *grpc.Repository,
	s3 *s3.Repository,
	mqPublisher *amqp.Publisher,
	subscriptionLimiter helpers.SubscriptionLimiter,
) *Service {
	return &Service{
		repo:                repo,
		grpc:                grpc,
		s3:                  s3,
		mqPublisher:         mqPublisher,
		subscriptionLimiter: subscriptionLimiter,
	}
}
