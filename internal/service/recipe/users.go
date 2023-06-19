package recipe

import (
	"context"
	"github.com/google/uuid"
	api "github.com/mephistolie/chefbook-backend-category/api/proto/implementation/v1"
)

func (s *Service) RateRecipe(recipeId, userId uuid.UUID, score int) error {
	return s.repo.RateRecipe(recipeId, userId, score)
}

func (s *Service) SaveToRecipeBook(recipeId, userId uuid.UUID) error {
	return s.repo.SaveToRecipeBook(recipeId, userId)
}

func (s *Service) RemoveFromRecipeBook(recipeId, userId uuid.UUID) error {
	return s.repo.RemoveFromRecipeBook(recipeId, userId)
}

func (s *Service) SetRecipeFavouriteStatus(recipeId, userId uuid.UUID, favourite bool) error {
	return s.repo.SetRecipeFavouriteStatus(recipeId, userId, favourite)
}

func (s *Service) SetRecipeCategories(recipeId, userId uuid.UUID, categories []uuid.UUID) error {
	err := s.repo.SetRecipeCategories(recipeId, userId, categories)
	if err == nil {
		go s.ValidateCategories(recipeId, userId, categories)
	}
	return err
}

func (s *Service) ValidateCategories(recipeId, userId uuid.UUID, categories []uuid.UUID) {
	if len(categories) == 0 {
		return
	}

	if res, err := s.grpc.Category.GetUserCategories(
		context.Background(),
		&api.GetUserCategoriesRequest{UserId: userId.String()},
	); err == nil {
		var ownedCategoryIds map[string]bool
		for _, category := range res.Categories {
			ownedCategoryIds[category.CategoryId] = true
		}

		var ownedCategories []uuid.UUID
		for _, category := range categories {
			if exists, ok := ownedCategoryIds[category.String()]; ok && exists {
				ownedCategories = append(ownedCategories, category)
			}
		}
		if len(ownedCategories) < len(categories) {
			_ = s.repo.SetRecipeCategories(recipeId, userId, ownedCategories)
		}
	}
}
