package grpc

import (
	"context"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-recipe/internal/transport/grpc/dto"
)

func (s *RecipeServer) GetRecipes(_ context.Context, req *api.GetRecipesRequest) (*api.GetRecipesResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	recipes := s.service.GetRecipes(dto.NewRecipesQuery(req), userId, req.UserLanguage)
	return dto.NewGetRecipesResponse(recipes), nil
}

func (s *RecipeServer) GetRandomRecipe(_ context.Context, req *api.GetRandomRecipeRequest) (*api.GetRecipeResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	var languages *[]string
	if len(req.RecipeLanguages) > 0 {
		languages = &req.RecipeLanguages
	}

	recipe, err := s.service.GetRandomRecipe(userId, languages, req.UserLanguage)
	if err != nil {
		return nil, err
	}

	return dto.NewGetRecipeResponse(recipe), nil
}

func (s *RecipeServer) GetRecipeBook(_ context.Context, req *api.GetRecipeBookRequest) (*api.GetRecipeBookResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	recipe, err := s.service.GetRecipesBook(userId, req.UserLanguage)
	if err != nil {
		return nil, err
	}

	return dto.NewGetRecipeBookResponse(recipe), nil
}
