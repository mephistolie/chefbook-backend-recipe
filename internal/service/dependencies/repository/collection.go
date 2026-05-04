package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	"time"
)

type Collection interface {
	GetCollections(ctx context.Context, userId, requesterId uuid.UUID) []entity.Collection
	GetCollectionsMap(ctx context.Context, collectionIds []uuid.UUID) map[uuid.UUID]entity.CollectionInfo

	CreateCollection(ctx context.Context, input entity.CollectionInput) (uuid.UUID, error)
	GetCollection(ctx context.Context, collectionId, userId uuid.UUID) (entity.Collection, error)
	UpdateCollection(ctx context.Context, input entity.CollectionInput) error
	DeleteCollection(ctx context.Context, collectionId, userId uuid.UUID) error

	GetCollectionKey(ctx context.Context, collectionId uuid.UUID) (uuid.UUID, time.Time, error)
	IsCollectionKeyValid(ctx context.Context, collectionId, key uuid.UUID) (bool, error)
	AddCollectionContributor(ctx context.Context, collectionId, contributorId uuid.UUID, role string) error
	RemoveCollectionContributors(ctx context.Context, collectionId uuid.UUID, contributorIds []uuid.UUID) error

	SaveCollectionToRecipeBook(ctx context.Context, collectionId, userId uuid.UUID) error
	RemoveCollectionFromRecipeBook(ctx context.Context, collectionId, userId uuid.UUID) error
}
