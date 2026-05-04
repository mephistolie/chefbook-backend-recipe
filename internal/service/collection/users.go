package collection

import (
	"context"
	"github.com/google/uuid"
)

func (s *Service) SaveCollectionToRecipeBook(ctx context.Context, collectionId, userId uuid.UUID) error {
	if _, err := s.repo.GetCollection(ctx, collectionId, userId); err != nil {
		return err
	}
	return s.repo.SaveCollectionToRecipeBook(ctx, collectionId, userId)
}

func (s *Service) RemoveCollectionFromRecipeBook(ctx context.Context, collectionId, userId uuid.UUID) error {
	return s.repo.RemoveCollectionFromRecipeBook(ctx, collectionId, userId)
}
