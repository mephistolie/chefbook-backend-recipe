package grpc

import (
	"context"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	"github.com/mephistolie/chefbook-backend-recipe/internal/transport/grpc/dto"
)

func (s *RecipeServer) GetRecipes(_ context.Context, req *api.GetRecipesRequest) (*api.GetRecipesResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	recipes := s.service.GetRecipes(dto.NewRecipesQuery(req), userId, entity.ValidatedLanguage(req.UserLanguage))
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

	recipe, err := s.service.GetRandomRecipe(userId, languages, entity.ValidatedLanguage(req.UserLanguage))
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

	recipeBook, err := s.service.GetRecipesBook(userId, entity.ValidatedLanguage(req.UserLanguage))
	if err != nil {
		return nil, err
	}

	return dto.NewGetRecipeBookResponse(recipeBook), nil
}

func (s *RecipeServer) GetRecipeNames(_ context.Context, req *api.GetRecipeNamesRequest) (*api.GetRecipeNamesResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	var recipeIds []uuid.UUID
	for _, rawId := range req.RecipeIds {
		if id, err := uuid.Parse(rawId); err == nil {
			recipeIds = append(recipeIds, id)
		}
	}

	recipeNames, err := s.service.GetRecipeNames(recipeIds, userId)
	if err != nil {
		return nil, err
	}

	response := make(map[string]string)
	for id, name := range recipeNames {
		response[id.String()] = name
	}

	return &api.GetRecipeNamesResponse{RecipeNames: response}, nil
}
