package recipe

import "github.com/google/uuid"

func (s *Service) GetRecipeOwner(recipeId uuid.UUID) (uuid.UUID, error) {
	policy, err := s.repo.GetRecipePolicy(recipeId)
	return policy.OwnerId, err
}
