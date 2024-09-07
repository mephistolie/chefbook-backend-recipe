package collection

import (
	"github.com/google/uuid"
)

func (s *Service) SaveCollectionToRecipeBook(collectionId, userId uuid.UUID) error {
	if _, err := s.repo.GetCollection(collectionId, userId); err != nil {
		return err
	}
	return s.repo.SaveCollectionToRecipeBook(collectionId, userId)
}

func (s *Service) RemoveCollectionFromRecipeBook(collectionId, userId uuid.UUID) error {
	return s.repo.RemoveCollectionFromRecipeBook(collectionId, userId)
}
