package recipe

import "github.com/google/uuid"

func (s *Service) GetRecipeOwner(userId uuid.UUID) (uuid.UUID, error) {
	return s.repo.GetRecipeOwner(userId)
}
