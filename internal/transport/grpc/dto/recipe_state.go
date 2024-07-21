package dto

import (
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
)

func NewGetRecipeBookResponse(data entity.RecipeBook) *api.GetRecipeBookResponse {
	return &api.GetRecipeBookResponse{
		Recipes:           newRecipeStates(data.Recipes),
		Collections:       newCollections(data.Collections),
		Tags:              newTags(data.Tags),
		TagGroups:         data.TagGroups,
		HasEncryptedVault: data.HasEncryptedVault,
		ProfilesInfo:      newProfilesInfo(data.ProfilesInfo),
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
	var collections []string
	for _, collection := range recipe.Collections {
		collections = append(collections, collection.String())
	}

	return &api.RecipeState{
		RecipeId:     recipe.Id.String(),
		Version:      recipe.Version,
		Translations: NewRecipeTranslations(recipe.Translations),
		Rating:       recipe.Rating,
		Votes:        recipe.Votes,
		Score:        recipe.Score,
		Tags:         recipe.Tags,
		Collections:  collections,
		IsFavourite:  recipe.IsFavourite,
	}
}
