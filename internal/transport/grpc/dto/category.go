package dto

import (
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
)

func newCategories(categories []entity.Category) []*api.RecipeCategory {
	response := make([]*api.RecipeCategory, len(categories))
	for i, category := range categories {
		response[i] = newCategory(category)
	}
	return response
}

func newCategoriesMap(categories map[string]entity.Category) map[string]*api.RecipeCategory {
	response := make(map[string]*api.RecipeCategory)
	for id, category := range categories {
		response[id] = newCategory(category)
	}
	return response
}

func newCategory(category entity.Category) *api.RecipeCategory {
	emoji := ""
	if category.Emoji != nil {
		emoji = *category.Emoji
	}

	return &api.RecipeCategory{
		Name:  category.Name,
		Emoji: emoji,
	}
}
