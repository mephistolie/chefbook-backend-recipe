package grpc

import (
	"context"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
)

func (s *RecipeServer) GetRecipePolicy(_ context.Context, req *api.GetRecipePolicyRequest) (*api.GetRecipePolicyResponse, error) {
	recipeId, err := uuid.Parse(req.RecipeId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	policy, err := s.recipeService.GetRecipePolicy(recipeId)
	if err != nil {
		return nil, err
	}

	return &api.GetRecipePolicyResponse{
		OwnerId:     policy.OwnerId.String(),
		Visibility:  policy.Visibility,
		IsEncrypted: policy.IsEncrypted,
	}, nil
}
