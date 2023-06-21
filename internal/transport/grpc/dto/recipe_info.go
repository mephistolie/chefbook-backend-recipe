package dto

import (
	"github.com/google/uuid"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/utils/strings/slices"
	"time"
)

const maxPageSize = 50

func NewRecipesQuery(req *api.GetRecipesRequest) entity.RecipesQuery {
	if req.PageSize != nil && *req.PageSize > 0 {
		*req.PageSize = maxPageSize
	}
	var authorIdPtr *uuid.UUID
	if req.AuthorId != nil {
		if authorId, err := uuid.Parse(*req.AuthorId); err == nil {
			authorIdPtr = &authorId
		}
	}
	sorting := entity.SortingCreationTimestamp
	if req.Sorting != nil && slices.Contains(entity.AvailableSortings, *req.Sorting) {
		sorting = *req.Sorting
	}
	var lastRecipeIdPtr *uuid.UUID
	if req.LastRecipeId != nil {
		if lastRecipeId, err := uuid.Parse(*req.LastRecipeId); err == nil {
			lastRecipeIdPtr = &lastRecipeId
		}
	}
	var lastCreationTimestamp *time.Time
	if req.LastCreationTimestamp != nil {
		timestamp := req.LastCreationTimestamp.AsTime()
		lastCreationTimestamp = &timestamp
	}
	var lastUpdateTimestamp *time.Time
	if req.LastUpdateTimestamp != nil {
		timestamp := req.LastUpdateTimestamp.AsTime()
		lastUpdateTimestamp = &timestamp
	}
	var languages *[]string
	if len(req.RecipeLanguages) > 0 {
		languages = &req.RecipeLanguages
	}

	return entity.RecipesQuery{
		PageSize:              req.PageSize,
		AuthorId:              authorIdPtr,
		Owned:                 req.Owned,
		Saved:                 req.Saved,
		Search:                req.Search,
		Sorting:               sorting,
		LastRecipeId:          lastRecipeIdPtr,
		LastCreationTimestamp: lastCreationTimestamp,
		LastUpdateTimestamp:   lastUpdateTimestamp,
		LastRating:            req.LastRating,
		LastVotes:             req.LastVotes,
		LastTime:              req.LastTime,
		LastCalories:          req.LastCalories,
		MinRating:             req.MinRating,
		MaxRating:             req.MaxRating,
		MinTime:               req.MinTime,
		MaxTime:               req.MaxTime,
		MinServings:           req.MinServings,
		MaxServings:           req.MaxServings,
		MinCalories:           req.MinCalories,
		MaxCalories:           req.MaxCalories,
		Languages:             languages,
	}
}

func NewGetRecipesResponse(data entity.DetailedRecipesInfo) *api.GetRecipesResponse {
	return &api.GetRecipesResponse{
		Recipes:    newRecipeInfos(data.Recipes),
		Tags:       newTags(data.Tags),
		Categories: newCategoriesMap(data.Categories),
	}
}

func newRecipeInfos(recipes []entity.RecipeInfo) []*api.RecipeInfo {
	response := make([]*api.RecipeInfo, len(recipes))
	for i, recipe := range recipes {
		response[i] = newRecipeInfo(recipe)
	}
	return response
}

func newRecipeInfo(recipe entity.RecipeInfo) *api.RecipeInfo {
	var categories []string
	for _, category := range recipe.Categories {
		categories = append(categories, category.String())
	}

	return &api.RecipeInfo{
		RecipeId: recipe.Id.String(),
		Name:     recipe.Name,

		OwnerId:     recipe.OwnerId.String(),
		OwnerName:   recipe.OwnerName,
		OwnerAvatar: recipe.OwnerAvatar,

		IsOwned:     recipe.IsOwned,
		IsSaved:     recipe.IsSaved,
		Visibility:  recipe.Visibility,
		IsEncrypted: recipe.IsEncrypted,

		Language: recipe.Language,
		Preview:  recipe.Preview,

		CreationTimestamp: timestamppb.New(recipe.CreationTimestamp),
		UpdateTimestamp:   timestamppb.New(recipe.UpdateTimestamp),
		Version:           recipe.Version,

		Rating: recipe.Rating,
		Votes:  recipe.Votes,
		Score:  recipe.Score,

		Tags:        recipe.Tags,
		Categories:  categories,
		IsFavourite: recipe.IsFavourite,

		Servings: recipe.Servings,
		Time:     recipe.Time,

		Calories: recipe.Calories,
	}
}
