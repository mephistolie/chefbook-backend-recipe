package grpc

import (
	"context"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	"github.com/mephistolie/chefbook-backend-recipe/internal/transport/grpc/dto"
)

func (s *RecipeServer) CreateRecipe(_ context.Context, req *api.RecipeInput) (*api.CreateRecipeResponse, error) {
	isEncryptedRecipeAllowed := s.subscriptionLimiter.IsEncryptionAllowed(req.UserSubscription)
	input, err := dto.NewRecipeInput(req, false, isEncryptedRecipeAllowed)
	if err != nil {
		return nil, err
	}

	id, version, err := s.recipeService.CreateRecipe(input)
	if err != nil {
		return nil, err
	}

	return &api.CreateRecipeResponse{
		RecipeId: id.String(),
		Version:  version,
	}, nil
}

func (s *RecipeServer) GetRecipe(_ context.Context, req *api.GetRecipeRequest) (*api.GetRecipeResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	recipeId, err := uuid.Parse(req.RecipeId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	var translatorIdPtr *uuid.UUID
	if req.TranslatorId != nil {
		if id, err := uuid.Parse(*req.TranslatorId); err == nil {
			translatorIdPtr = &id
		}
	}

	recipe, err := s.recipeService.GetRecipe(recipeId, userId, entity.ValidatedLanguage(req.Language), translatorIdPtr)
	if err != nil {
		return nil, err
	}

	return dto.NewGetRecipeResponse(recipe), nil
}

func (s *RecipeServer) UpdateRecipe(_ context.Context, req *api.RecipeInput) (*api.UpdateRecipeResponse, error) {
	isEncryptedRecipeAllowed := s.subscriptionLimiter.IsEncryptionAllowed(req.UserSubscription)
	input, err := dto.NewRecipeInput(req, true, isEncryptedRecipeAllowed)
	if err != nil {
		return nil, err
	}

	version, err := s.recipeService.UpdateRecipe(input)
	if err != nil {
		return nil, err
	}

	return &api.UpdateRecipeResponse{Version: version}, nil
}

func (s *RecipeServer) DeleteRecipe(_ context.Context, req *api.DeleteRecipeRequest) (*api.DeleteRecipeResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	recipeId, err := uuid.Parse(req.RecipeId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	if err = s.recipeService.DeleteRecipe(recipeId, userId); err != nil {
		return nil, err
	}

	return &api.DeleteRecipeResponse{Message: "recipe deleted"}, nil
}
