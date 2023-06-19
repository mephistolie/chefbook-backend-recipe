package dto

import (
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewGetRecipeResponse(data entity.DetailedRecipe) *api.GetRecipeResponse {
	return &api.GetRecipeResponse{
		Recipe:     newRecipe(data.Recipe),
		Tags:       newTags(data.Tags),
		Categories: newCategoriesMap(data.Categories),
	}
}

func newRecipe(recipe entity.Recipe) *api.Recipe {
	ownerName := ""
	if recipe.OwnerName != nil {
		ownerName = *recipe.OwnerName
	}
	ownerAvatar := ""
	if recipe.OwnerAvatar != nil {
		ownerName = *recipe.OwnerAvatar
	}
	description := ""
	if recipe.Description != nil {
		description = *recipe.Description
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
	var time int32 = 0
	if recipe.Time != nil {
		time = int32(*recipe.Time)
	}
	var calories int32 = 0
	if recipe.Calories != nil {
		calories = int32(*recipe.Calories)
	}
	var macronutrientsPtr *api.Macronutrients
	if recipe.Macronutrients != nil {
		macronutrients := api.Macronutrients{}
		if recipe.Macronutrients.Protein != nil {
			macronutrients.Protein = int32(*recipe.Macronutrients.Protein)
		}
		if recipe.Macronutrients.Fats != nil {
			macronutrients.Fats = int32(*recipe.Macronutrients.Fats)
		}
		if recipe.Macronutrients.Carbohydrates != nil {
			macronutrients.Carbohydrates = int32(*recipe.Macronutrients.Carbohydrates)
		}
	}

	return &api.Recipe{
		RecipeId: recipe.Id.String(),
		Name:     recipe.Name,

		OwnerId:     recipe.OwnerId.String(),
		OwnerName:   ownerName,
		OwnerAvatar: ownerAvatar,

		IsOwned:     recipe.IsOwned,
		IsSaved:     recipe.IsSaved,
		Visibility:  recipe.Visibility,
		IsEncrypted: recipe.IsEncrypted,

		Language:    recipe.Language,
		Description: description,
		Preview:     preview,

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
		Time:     time,

		Calories:       calories,
		Macronutrients: macronutrientsPtr,

		Ingredients: newIngredientsResponse(recipe.Ingredients),
		Cooking:     newCookingResponse(recipe.Cooking),
	}
}
