package mq

import "github.com/google/uuid"

func (s *Service) DeleteUserRecipes(userId uuid.UUID, deleteSharedData bool, messageId uuid.UUID) error {
	return s.repo.DeleteUserRecipes(userId, deleteSharedData, messageId)
}
