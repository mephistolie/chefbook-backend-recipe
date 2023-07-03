package mq

import "github.com/google/uuid"

func (s *Service) DeleteUserEncryptedRecipes(userId uuid.UUID, messageId uuid.UUID) error {
	return s.repo.DeleteUserEncryptedRecipes(userId, messageId)
}

func (s *Service) DeleteUserRecipes(userId uuid.UUID, deleteSharedData bool, messageId uuid.UUID) error {
	return s.repo.DeleteUserRecipes(userId, deleteSharedData, messageId)
}
