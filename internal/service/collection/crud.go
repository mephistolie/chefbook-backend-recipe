package collection

import (
	"context"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
)

func (s *Service) GetCollections(ctx context.Context, userId uuid.UUID, requesterId uuid.UUID) entity.DetailedCollections {
	collections := s.repo.GetCollections(ctx, userId, requesterId)

	var profileIds []string
	for _, collection := range collections {
		for _, contributor := range collection.Contributors {
			profileIds = append(profileIds, contributor.Id.String())
		}
	}
	profilesInfo := s.getProfilesInfo(ctx, profileIds)

	return entity.DetailedCollections{
		Collections:  collections,
		ProfilesInfo: profilesInfo,
	}
}

func (s *Service) CreateCollection(ctx context.Context, collection entity.CollectionInput) (uuid.UUID, error) {
	return s.repo.CreateCollection(ctx, collection)
}

func (s *Service) GetCollection(ctx context.Context, collectionId uuid.UUID, userId uuid.UUID) (entity.DetailedCollection, error) {
	collection, err := s.repo.GetCollection(ctx, collectionId, userId)
	if err != nil {
		return entity.DetailedCollection{}, err
	}

	var profileIds []string
	for _, contributor := range collection.Contributors {
		profileIds = append(profileIds, contributor.Id.String())
	}
	profilesInfo := s.getProfilesInfo(ctx, profileIds)

	return entity.DetailedCollection{
		Collection:   collection,
		ProfilesInfo: profilesInfo,
	}, nil
}

func (s *Service) UpdateCollection(ctx context.Context, collection entity.CollectionInput) error {
	return s.repo.UpdateCollection(ctx, collection)
}

func (s *Service) DeleteCollection(ctx context.Context, collectionId, userId uuid.UUID) error {
	return s.repo.DeleteCollection(ctx, collectionId, userId)
}
