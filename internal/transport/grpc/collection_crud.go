package grpc

import (
	"context"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-recipe/internal/transport/grpc/dto"
)

func (s *RecipeServer) GetCollections(_ context.Context, req *api.GetCollectionsRequest) (*api.GetCollectionsResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	requesterId, err := uuid.Parse(req.RequesterId)
	if err != nil {

		return nil, fail.GrpcInvalidBody
	}

	collections := s.collectionService.GetCollections(userId, requesterId)
	if err != nil {
		return nil, err
	}

	return dto.NewGetCollectionsResponse(collections), nil
}

func (s *RecipeServer) CreateCollection(_ context.Context, req *api.CreateCollectionRequest) (*api.CreateCollectionResponse, error) {
	input, err := dto.NewCreateCollectionInput(req)
	if err != nil {
		return nil, err
	}

	id, err := s.collectionService.CreateCollection(input)
	if err != nil {
		return nil, err
	}

	return &api.CreateCollectionResponse{CollectionId: id.String()}, nil
}

func (s *RecipeServer) GetCollection(_ context.Context, req *api.GetCollectionRequest) (*api.GetCollectionResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	collectionId, err := uuid.Parse(req.CollectionId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	collection, err := s.collectionService.GetCollection(collectionId, userId)
	if err != nil {
		return nil, err
	}

	return dto.NewGetCollectionResponse(collection), nil
}

func (s *RecipeServer) UpdateCollection(_ context.Context, req *api.UpdateCollectionRequest) (*api.UpdateCollectionResponse, error) {
	input, err := dto.NewUpdateCollectionInput(req)
	if err != nil {
		return nil, err
	}

	if err = s.collectionService.UpdateCollection(input); err != nil {
		return nil, err
	}

	return &api.UpdateCollectionResponse{Message: "collection updated"}, nil
}

func (s *RecipeServer) DeleteCollection(_ context.Context, req *api.DeleteCollectionRequest) (*api.DeleteCollectionResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	collectionId, err := uuid.Parse(req.CollectionId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	if err = s.collectionService.DeleteCollection(collectionId, userId); err != nil {
		return nil, err
	}

	return &api.DeleteCollectionResponse{Message: "collection deleted"}, nil
}
