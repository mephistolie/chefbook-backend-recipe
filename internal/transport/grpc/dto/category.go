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

func newCategoriesMap(categories map[string]entity.Category) map[string]*api.RecipeCategoryInfo {
	response := make(map[string]*api.RecipeCategoryInfo)
	for id, category := range categories {
		response[id] = newCategoryInfo(category)
	}
	return response
}

func newCategoryInfo(category entity.Category) *api.RecipeCategoryInfo {
	return &api.RecipeCategoryInfo{
		Name:  category.Name,
		Emoji: category.Emoji,
	}
}

func newCategory(category entity.Category) *api.RecipeCategory {
	return &api.RecipeCategory{
		Id:    category.Id,
		Name:  category.Name,
		Emoji: category.Emoji,
	}
}
