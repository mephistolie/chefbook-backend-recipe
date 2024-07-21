package app

import (
	"context"
	"fmt"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/shutdown"
	recipepb "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-recipe/internal/config"
	"github.com/mephistolie/chefbook-backend-recipe/internal/helpers"
	grpcRepo "github.com/mephistolie/chefbook-backend-recipe/internal/repository/grpc"
	"github.com/mephistolie/chefbook-backend-recipe/internal/repository/postgres"
	"github.com/mephistolie/chefbook-backend-recipe/internal/repository/s3"
	"github.com/mephistolie/chefbook-backend-recipe/internal/transport/dependencies/service"
	recipe "github.com/mephistolie/chefbook-backend-recipe/internal/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"time"
)

func Run(cfg *config.Config) {
	log.Init(*cfg.LogsPath, *cfg.Environment == config.EnvDev)
	cfg.Print()

	subscriptionLimiter := helpers.NewSubscriptionLimiter(cfg.Subscription)

	db, err := postgres.Connect(cfg.Database)
	if err != nil {
		log.Fatal(err)
		return
	}

	repository := postgres.NewRepository(db, cfg.Recipes)

	grpcRepository, err := grpcRepo.NewRepository(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}

	s3Repository, err := s3.NewRepository(cfg, subscriptionLimiter)
	if err != nil {
		log.Fatal(err)
		return
	}

	mqPublisher, err := NewMqPublisher(cfg.Amqp, repository)
	if err != nil {
		log.Fatal(err)
		return
	}

	recipeService, err := service.New(cfg, repository, repository, repository, grpcRepository, s3Repository, mqPublisher, subscriptionLimiter)
	if err != nil {
		log.Fatal(err)
		return
	}

	mqSubscriber, err := NewMqSubscriber(cfg.Amqp, recipeService.MQ)
	if err != nil {
		log.Fatal(err)
		return
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *cfg.Port))
	if err != nil {
		log.Fatal(err)
		return
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			log.UnaryServerInterceptor(),
		),
	)

	healthServer := health.NewServer()
	recipeServer := recipe.NewServer(recipeService.Recipe, recipeService.Collection, subscriptionLimiter)

	go monitorHealthChecking(db, healthServer)

	recipepb.RegisterRecipeServiceServer(grpcServer, recipeServer)
	healthpb.RegisterHealthServer(grpcServer, healthServer)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Errorf("error occurred while running http server: %s\n", err.Error())
		} else {
			log.Info("gRPC server started")
		}
	}()

	wait := shutdown.Graceful(context.Background(), 5*time.Second, map[string]shutdown.Operation{
		"grpc-server": func(ctx context.Context) error {
			grpcServer.GracefulStop()
			return nil
		},
		"database": func(ctx context.Context) error {
			return db.Close()
		},
		"services": func(ctx context.Context) error {
			return grpcRepository.Stop()
		},
		"mq": func(ctx context.Context) error {
			_ = mqPublisher.Stop()
			_ = mqSubscriber.Stop()
			return nil
		},
	})
	<-wait
}
