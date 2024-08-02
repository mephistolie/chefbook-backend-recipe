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
	if req.Score < 0 || req.Score > 5 {
		return nil, fail.GrpcInvalidBody
	}

	if err = s.recipeService.RateRecipe(recipeId, userId, int(req.Score)); err != nil {
		return nil, err
	}

	return &api.RateRecipeResponse{Message: "score set"}, nil
}

func (s *RecipeServer) SaveRecipeToRecipeBook(_ context.Context, req *api.SaveRecipeToRecipeBookRequest) (*api.SaveRecipeToRecipeBookResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	recipeId, err := uuid.Parse(req.RecipeId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	if err = s.recipeService.SaveRecipeToRecipeBook(recipeId, userId); err != nil {
		return nil, err
	}

	return &api.SaveRecipeToRecipeBookResponse{Message: "recipe saved to recipe book"}, nil
}

func (s *RecipeServer) RemoveRecipeFromRecipeBook(_ context.Context, req *api.RemoveRecipeFromRecipeBookRequest) (*api.RemoveRecipeFromRecipeBookResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	recipeId, err := uuid.Parse(req.RecipeId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	if err = s.recipeService.RemoveRecipeFromRecipeBook(recipeId, userId); err != nil {
		return nil, err
	}

	return &api.RemoveRecipeFromRecipeBookResponse{Message: "recipe removed from recipe book"}, nil
}

func (s *RecipeServer) SaveRecipeToFavourites(_ context.Context, req *api.SaveRecipeToFavouritesRequest) (*api.SaveRecipeToFavouritesResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	recipeId, err := uuid.Parse(req.RecipeId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	if err = s.recipeService.SaveRecipeToFavourites(recipeId, userId); err != nil {
		return nil, err
	}

	return &api.SaveRecipeToFavouritesResponse{Message: "recipe saved to favourites"}, nil
}

func (s *RecipeServer) RemoveRecipeFromFavourites(_ context.Context, req *api.RemoveRecipeFromFavouritesRequest) (*api.RemoveRecipeFromFavouritesResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	recipeId, err := uuid.Parse(req.RecipeId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	if err = s.recipeService.RemoveRecipeFromFavourites(recipeId, userId); err != nil {
		return nil, err
	}

	return &api.RemoveRecipeFromFavouritesResponse{Message: "recipe removed from favourites"}, nil
}

func (s *RecipeServer) AddRecipeToCollection(_ context.Context, req *api.AddRecipeToCollectionRequest) (*api.AddRecipeToCollectionResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	recipeId, err := uuid.Parse(req.RecipeId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	collectionId, err := uuid.Parse(req.CollectionId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	if err = s.recipeService.AddRecipeToCollection(recipeId, collectionId, userId); err != nil {
		return nil, err
	}

	return &api.AddRecipeToCollectionResponse{Message: "recipe saved to collection"}, nil
}

func (s *RecipeServer) RemoveRecipeFromCollection(_ context.Context, req *api.RemoveRecipeFromCollectionRequest) (*api.RemoveRecipeFromCollectionResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	recipeId, err := uuid.Parse(req.RecipeId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	collectionId, err := uuid.Parse(req.CollectionId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	if err = s.recipeService.RemoveRecipeFromCollection(recipeId, collectionId, userId); err != nil {
		return nil, err
	}

	return &api.RemoveRecipeFromCollectionResponse{Message: "recipe removed from collection"}, nil
}

func (s *RecipeServer) SetRecipeCollections(_ context.Context, req *api.SetRecipeCollectionsRequest) (*api.SetRecipeCollectionsResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	recipeId, err := uuid.Parse(req.RecipeId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	req.CollectionIds = sliceUtils.RemoveDuplicates(req.CollectionIds)
	var collectionIds []uuid.UUID
	for i, rawId := range req.CollectionIds {
		if i > maxRecipeCategoriesCount {
			break
		}
		if id, err := uuid.Parse(rawId); err == nil {
			collectionIds = append(collectionIds, id)
		}
	}

	if err = s.recipeService.SetRecipeCollections(recipeId, userId, collectionIds); err != nil {
		return nil, err
	}

	return &api.SetRecipeCollectionsResponse{Message: "recipe collections set"}, nil
}
