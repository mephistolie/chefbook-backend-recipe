package grpc

import (
	"context"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
)

func (s *RecipeServer) SaveCollectionToRecipeBook(_ context.Context, req *api.SaveCollectionToRecipeBookRequest) (*api.SaveCollectionToRecipeBookResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	collectionId, err := uuid.Parse(req.CollectionId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	if err = s.collectionService.SaveCollectionToRecipeBook(collectionId, userId); err != nil {
		return nil, err
	}

	return &api.SaveCollectionToRecipeBookResponse{Message: "collection added to recipe book"}, nil
}

func (s *RecipeServer) RemoveCollectionFromRecipeBook(_ context.Context, req *api.RemoveCollectionFromRecipeBookRequest) (*api.RemoveCollectionFromRecipeBookResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	collectionId, err := uuid.Parse(req.CollectionId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	if err = s.collectionService.RemoveCollectionFromRecipeBook(collectionId, userId); err != nil {
		return nil, err
	}

	return &api.RemoveCollectionFromRecipeBookResponse{Message: "collection removed from recipe book"}, nil
}
