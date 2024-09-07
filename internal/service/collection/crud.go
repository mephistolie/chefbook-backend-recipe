package collection

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
)

func (s *Service) GetCollections(userId uuid.UUID, requesterId uuid.UUID) entity.DetailedCollections {
	collections := s.repo.GetCollections(userId, requesterId)

	var profileIds []string
	for _, collection := range collections {
		for _, contributor := range collection.Contributors {
			profileIds = append(profileIds, contributor.Id.String())
		}
	}
	profilesInfo := s.getProfilesInfo(profileIds)

	return entity.DetailedCollections{
		Collections:  collections,
		ProfilesInfo: profilesInfo,
	}
}

func (s *Service) CreateCollection(collection entity.CollectionInput) (uuid.UUID, error) {
	return s.repo.CreateCollection(collection)
}

func (s *Service) GetCollection(collectionId uuid.UUID, userId uuid.UUID) (entity.DetailedCollection, error) {
	collection, err := s.repo.GetCollection(collectionId, userId)
	if err != nil {
		return entity.DetailedCollection{}, err
	}

	var profileIds []string
	for _, contributor := range collection.Contributors {
		profileIds = append(profileIds, contributor.Id.String())
	}
	profilesInfo := s.getProfilesInfo(profileIds)

	return entity.DetailedCollection{
		Collection:   collection,
		ProfilesInfo: profilesInfo,
	}, nil
}

func (s *Service) UpdateCollection(collection entity.CollectionInput) error {
	return s.repo.UpdateCollection(collection)
}

func (s *Service) DeleteCollection(collectionId, userId uuid.UUID) error {
	return s.repo.DeleteCollection(collectionId, userId)
}
