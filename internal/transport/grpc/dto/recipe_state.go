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
	var categories []string
	for _, category := range recipe.Categories {
		categories = append(categories, category.String())
	}

	return &api.RecipeState{
		RecipeId:    recipe.Id.String(),
		OwnerName:   recipe.OwnerName,
		OwnerAvatar: recipe.OwnerAvatar,
		Version:     recipe.Version,
		Rating:      recipe.Rating,
		Votes:       recipe.Votes,
		Score:       recipe.Score,
		Categories:  categories,
		IsFavourite: recipe.IsFavourite,
	}
}
