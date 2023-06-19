package dto

import (
	"github.com/google/uuid"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/utils/strings/slices"
	"time"
)

func NewRecipesQuery(req *api.GetRecipesRequest) entity.RecipesQuery {
	var pageSizePtr *int
	if req.PageSize > 0 {
		pageSize := int(req.PageSize)
		if pageSize > 50 {
			pageSize = 50
		}
		pageSizePtr = &pageSize
	}
	var authorIdPtr *uuid.UUID
	if authorId, err := uuid.Parse(req.AuthorId); err == nil {
		authorIdPtr = &authorId
	}
	var search *string
	if len(req.Search) > 0 {
		search = &req.Search
	}
	sorting := req.Sorting
	if !slices.Contains(entity.AvailableSortings, sorting) {
		sorting = entity.SortingCreationTimestamp
	}
	var lastRecipeIdPtr *uuid.UUID
	if lastRecipeId, err := uuid.Parse(req.LastRecipeId); err == nil {
		lastRecipeIdPtr = &lastRecipeId
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
	var lastRating *float32
	if req.LastRating > 0 {
		lastRating = &req.LastRating
	}
	var lastVotes *int32
	if req.LastVotes > 0 {
		lastVotes = &req.LastVotes
	}
	var lastTime *int32
	if req.LastTime > 0 {
		lastTime = &req.LastTime
	}
	var lastCalories *int32
	if req.LastCalories > 0 {
		lastCalories = &req.LastCalories
	}
	var minRating *int32
	if req.MinRating > 0 {
		minRating = &req.MinRating
	}
	var maxRating *int32
	if req.MaxRating > 0 {
		maxRating = &req.MaxRating
	}
	var minTime *int32
	if req.MinTime > 0 {
		minTime = &req.MinTime
	}
	var maxTime *int32
	if req.MaxTime > 0 {
		maxTime = &req.MaxTime
	}
	var minServings *int32
	if req.MinServings > 0 {
		minServings = &req.MinServings
	}
	var maxServings *int32
	if req.MaxServings > 0 {
		maxServings = &req.MaxServings
	}
	var minCalories *int32
	if req.MinCalories > 0 {
		minCalories = &req.MinCalories
	}
	var maxCalories *int32
	if req.MaxCalories > 0 {
		maxCalories = &req.MaxCalories
	}
	var languages *[]string
	if len(req.RecipeLanguages) > 0 {
		languages = &req.RecipeLanguages
	}

	return entity.RecipesQuery{
		PageSize:              pageSizePtr,
		AuthorId:              authorIdPtr,
		Owned:                 req.Owned,
		Saved:                 req.Saved,
		Search:                search,
		Sorting:               sorting,
		LastRecipeId:          lastRecipeIdPtr,
		LastCreationTimestamp: lastCreationTimestamp,
		LastUpdateTimestamp:   lastUpdateTimestamp,
		LastRating:            lastRating,
		LastVotes:             lastVotes,
		LastTime:              lastTime,
		LastCalories:          lastCalories,
		MinRating:             minRating,
		MaxRating:             maxRating,
		MinTime:               minTime,
		MaxTime:               maxTime,
		MinServings:           minServings,
		MaxServings:           maxServings,
		MinCalories:           minCalories,
		MaxCalories:           maxCalories,
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
	ownerName := ""
	if recipe.OwnerName != nil {
		ownerName = *recipe.OwnerName
	}
	ownerAvatar := ""
	if recipe.OwnerAvatar != nil {
		ownerName = *recipe.OwnerAvatar
	}
	preview := ""
	if recipe.Preview != nil {
		preview = *recipe.Preview
	}
	var score int32 = 0
	if recipe.Score != nil {
		score = int32(*recipe.Score)
	}
	var categories []string
	for _, category := range recipe.Categories {
		categories = append(categories, category.String())
	}
	var servings int32 = 0
	if recipe.Servings != nil {
		servings = int32(*recipe.Servings)
	}
	var timePtr int32 = 0
	if recipe.Time != nil {
		timePtr = int32(*recipe.Time)
	}
	var calories int32 = 0
	if recipe.Calories != nil {
		calories = int32(*recipe.Calories)
	}

	return &api.RecipeInfo{
		RecipeId: recipe.Id.String(),
		Name:     recipe.Name,

		OwnerId:     recipe.OwnerId.String(),
		OwnerName:   ownerName,
		OwnerAvatar: ownerAvatar,

		IsOwned:     recipe.IsOwned,
		IsSaved:     recipe.IsSaved,
		Visibility:  recipe.Visibility,
		IsEncrypted: recipe.IsEncrypted,

		Language: recipe.Language,
		Preview:  preview,

		CreationTimestamp: timestamppb.New(recipe.CreationTimestamp),
		UpdateTimestamp:   timestamppb.New(recipe.UpdateTimestamp),
		Version:           recipe.Version,

		Rating: recipe.Rating,
		Votes:  recipe.Votes,
		Score:  score,

		Tags:        recipe.Tags,
		Categories:  categories,
		IsFavourite: recipe.IsFavourite,

		Servings: servings,
		Time:     timePtr,

		Calories: calories,
	}
}
