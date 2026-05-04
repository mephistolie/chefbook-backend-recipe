package recipe

import (
	"context"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
)

func (s *Service) GetRecipePolicy(ctx context.Context, recipeId uuid.UUID) (entity.RecipePolicy, error) {
	return s.recipeRepo.GetRecipePolicy(ctx, recipeId)
}
