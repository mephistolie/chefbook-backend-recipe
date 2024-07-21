package grpc

import (
	"context"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-recipe/internal/transport/grpc/dto"
)

func (s *RecipeServer) TranslateRecipe(_ context.Context, req *api.TranslateRecipeRequest) (*api.TranslateRecipeResponse, error) {
	recipeId, err := uuid.Parse(req.RecipeId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	translation, err := dto.NewRecipeTranslation(req)
	if err != nil {
		return nil, err
	}

	if err = s.recipeService.TranslateRecipe(recipeId, translation); err != nil {
		return nil, err
	}
	return &api.TranslateRecipeResponse{Message: "recipe translation saved"}, nil
}

func (s *RecipeServer) DeleteRecipeTranslation(_ context.Context, req *api.DeleteRecipeTranslationRequest) (*api.DeleteRecipeTranslationResponse, error) {
	requesterId, err := uuid.Parse(req.RequesterId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	recipeId, err := uuid.Parse(req.RecipeId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	if len(req.Language) != 2 {
		return nil, fail.GrpcInvalidBody
	}

	if err = s.recipeService.DeleteRecipeTranslation(recipeId, requesterId, req.Language); err != nil {
		return nil, err
	}
	return &api.DeleteRecipeTranslationResponse{Message: "recipe translation deleted"}, nil
}
