package grpc

import (
	"context"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	sliceUtils "github.com/mephistolie/chefbook-backend-common/utils/slices"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
)

const (
	maxRecipeCategoriesCount = 20
)

func (s *RecipeServer) RateRecipe(_ context.Context, req *api.RateRecipeRequest) (*api.RateRecipeResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	recipeId, err := uuid.Parse(req.RecipeId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	if err = s.service.RateRecipe(recipeId, userId, int(req.Score)); err != nil {
		return nil, err
	}

	return &api.RateRecipeResponse{Message: "score set"}, nil
}

func (s *RecipeServer) SaveToRecipeBook(_ context.Context, req *api.SaveToRecipeBookRequest) (*api.SaveToRecipeBookResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	recipeId, err := uuid.Parse(req.RecipeId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	if err = s.service.SaveToRecipeBook(recipeId, userId); err != nil {
		return nil, err
	}

	return &api.SaveToRecipeBookResponse{Message: "recipe saved to recipe book"}, nil
}

func (s *RecipeServer) RemoveFromRecipeBook(_ context.Context, req *api.RemoveFromRecipeBookRequest) (*api.RemoveFromRecipeBookResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	recipeId, err := uuid.Parse(req.RecipeId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	if err = s.service.RemoveFromRecipeBook(recipeId, userId); err != nil {
		return nil, err
	}

	return &api.RemoveFromRecipeBookResponse{Message: "recipe removed from recipe book"}, nil
}

func (s *RecipeServer) SetRecipeFavouriteStatus(_ context.Context, req *api.SetRecipeFavouriteStatusRequest) (*api.SetRecipeFavouriteStatusResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	recipeId, err := uuid.Parse(req.RecipeId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	if err = s.service.SetRecipeFavouriteStatus(recipeId, userId, req.Favourite); err != nil {
		return nil, err
	}

	return &api.SetRecipeFavouriteStatusResponse{Message: "recipe favourite status set"}, nil
}

func (s *RecipeServer) SetRecipeCategories(_ context.Context, req *api.SetRecipeCategoriesRequest) (*api.SetRecipeCategoriesResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	recipeId, err := uuid.Parse(req.RecipeId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	req.CategoryIds = sliceUtils.RemoveDuplicates(req.CategoryIds)
	var categoryIds []uuid.UUID
	for i, rawId := range req.CategoryIds {
		if i > maxRecipeCategoriesCount {
			break
		}
		if id, err := uuid.Parse(rawId); err == nil {
			categoryIds = append(categoryIds, id)
		}
	}

	if err = s.service.SetRecipeCategories(recipeId, userId, categoryIds); err != nil {
		return nil, err
	}

	return &api.SetRecipeCategoriesResponse{Message: "recipe categories set"}, nil
}
