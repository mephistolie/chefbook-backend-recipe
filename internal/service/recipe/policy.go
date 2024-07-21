package recipe

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
)

func (s *Service) GetRecipePolicy(recipeId uuid.UUID) (entity.RecipePolicy, error) {
	return s.recipeRepo.GetRecipePolicy(recipeId)
}
