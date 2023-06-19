package dto

import (
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
)

func NewGetRecipeBookResponse(data entity.DetailedRecipesState) *api.GetRecipeBookResponse {
	return &api.GetRecipeBookResponse{
		Recipes:    newRecipeStates(data.Recipes),
		Tags:       newTags(data.Tags),
		Categories: newCategories(data.Categories),
	}
}

func newRecipeStates(recipes []entity.RecipeState) []*api.RecipeState {
	response := make([]*api.RecipeState, len(recipes))
	for i, recipe := range recipes {
		response[i] = newRecipeState(recipe)
	}
	return response
}

func newRecipeState(recipe entity.RecipeState) *api.RecipeState {
	ownerName := ""
	if recipe.OwnerName != nil {
		ownerName = *recipe.OwnerName
	}
	ownerAvatar := ""
	if recipe.OwnerAvatar != nil {
		ownerName = *recipe.OwnerAvatar
	}
	var score int32 = 0
	if recipe.Score != nil {
		score = int32(*recipe.Score)
	}
	var categories []string
	for _, category := range recipe.Categories {
		categories = append(categories, category.String())
	}

	return &api.RecipeState{
		RecipeId:    recipe.Id.String(),
		OwnerName:   ownerName,
		OwnerAvatar: ownerAvatar,
		Version:     recipe.Version,
		Rating:      recipe.Rating,
		Votes:       recipe.Votes,
		Score:       score,
		Categories:  categories,
		IsFavourite: recipe.IsFavourite,
	}
}
