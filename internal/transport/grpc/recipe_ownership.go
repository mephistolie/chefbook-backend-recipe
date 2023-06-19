package grpc

import (
	"context"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
)

func (s *RecipeServer) GetRecipeOwner(_ context.Context, req *api.GetRecipeOwnerRequest) (*api.GetRecipeOwnerResponse, error) {
	recipeId, err := uuid.Parse(req.RecipeId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	ownerId, err := s.service.GetRecipeOwner(recipeId)
	if err != nil {
		return nil, err
	}

	return &api.GetRecipeOwnerResponse{OwnerId: ownerId.String()}, nil
}
