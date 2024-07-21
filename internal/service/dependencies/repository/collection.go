package repository

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	"time"
)

type Collection interface {
	GetCollections(userId, requesterId uuid.UUID) []entity.Collection
	GetCollectionsMap(collectionIds []uuid.UUID) map[uuid.UUID]entity.CollectionInfo

	CreateCollection(input entity.CollectionInput) (uuid.UUID, error)
	GetCollection(collectionId, userId uuid.UUID) (entity.Collection, error)
	UpdateCollection(input entity.CollectionInput) error
	DeleteCollection(collectionId, userId uuid.UUID) error

	GetCollectionKey(collectionId uuid.UUID) (uuid.UUID, time.Time, error)
	IsCollectionKeyValid(collectionId, key uuid.UUID) (bool, error)
	AddCollectionContributor(collectionId, contributorId uuid.UUID, role string) error
	RemoveCollectionContributors(collectionId uuid.UUID, contributorIds []uuid.UUID) error

	SaveCollectionToRecipeBook(collectionId, userId uuid.UUID) error
	RemoveCollectionFromRecipeBook(collectionId, userId uuid.UUID) error
}
