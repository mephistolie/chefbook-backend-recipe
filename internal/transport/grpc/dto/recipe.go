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
	var categories []string
	for _, category := range recipe.Categories {
		categories = append(categories, category.String())
	}
	var macronutrients *api.Macronutrients
	if recipe.Macronutrients != nil {
		macronutrients = &api.Macronutrients{
			Protein:       recipe.Macronutrients.Protein,
			Fats:          recipe.Macronutrients.Fats,
			Carbohydrates: recipe.Macronutrients.Carbohydrates,
		}
	}

	return &api.Recipe{
		RecipeId: recipe.Id.String(),
		Name:     recipe.Name,

		OwnerId:     recipe.OwnerId.String(),
		OwnerName:   recipe.OwnerName,
		OwnerAvatar: recipe.OwnerAvatar,

		IsOwned:     recipe.IsOwned,
		IsSaved:     recipe.IsSaved,
		Visibility:  recipe.Visibility,
		IsEncrypted: recipe.IsEncrypted,

		Language:     recipe.Language,
		Translations: NewRecipeTranslations(recipe.Translations),
		Description:  recipe.Description,

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

		Calories:       recipe.Calories,
		Macronutrients: macronutrients,

		Ingredients: newIngredientsResponse(recipe.Ingredients),
		Cooking:     newCookingResponse(recipe.Cooking),
		Pictures:    newRecipesResponse(recipe.Pictures),
	}
}
