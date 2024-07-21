package dto

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"github.com/mephistolie/chefbook-backend-recipe/api/model"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	"k8s.io/utils/strings/slices"
)

const (
	maxCollectionNameLength = 64
)

func NewGetCollectionsResponse(response entity.DetailedCollections) *api.GetCollectionsResponse {
	collections := make([]*api.Collection, len(response.Collections))
	for i, collection := range response.Collections {
		collections[i] = newCollection(collection)
	}
	return &api.GetCollectionsResponse{
		Collections:  collections,
		ProfilesInfo: newProfilesInfo(response.ProfilesInfo),
	}
}

func newCollections(collections []entity.Collection) []*api.Collection {
	dtos := make([]*api.Collection, len(collections))
	for i, collection := range collections {
		dtos[i] = newCollection(collection)
	}
	return dtos
}

func newCollectionsMap(categories map[uuid.UUID]entity.CollectionInfo) map[string]*api.CollectionInfo {
	response := make(map[string]*api.CollectionInfo)
	for id, category := range categories {
		response[id.String()] = newCollectionInfo(category)
	}
	return response
}

func NewGetCollectionResponse(response entity.DetailedCollection) *api.GetCollectionResponse {
	return &api.GetCollectionResponse{
		Collection:   newCollection(response.Collection),
		ProfilesInfo: newProfilesInfo(response.ProfilesInfo),
	}
}

func newCollection(collection entity.Collection) *api.Collection {
	contributors := make([]*api.CollectionContributor, len(collection.Contributors))
	for i, contributor := range collection.Contributors {
		contributors[i] = &api.CollectionContributor{
			ContributorId: contributor.Id.String(),
			Role:          contributor.Role,
		}
	}

	return &api.Collection{
		CollectionId: collection.Id.String(),
		Name:         collection.Name,
		Visibility:   collection.Visibility,
		Contributors: contributors,
	}
}

func newCollectionInfo(collection entity.CollectionInfo) *api.CollectionInfo {
	return &api.CollectionInfo{
		Name: collection.Name,
	}
}

func NewCreateCollectionInput(req *api.CreateCollectionRequest) (entity.CollectionInput, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return entity.CollectionInput{}, fail.GrpcInvalidBody
	}

	if len(req.Name) == 0 {
		return entity.CollectionInput{}, fail.GrpcInvalidBody
	}

	var collectionId *uuid.UUID
	if req.CollectionId != nil {
		if parsedId, err := uuid.Parse(*req.CollectionId); err == nil {
			collectionId = &parsedId
		}
	}
	if collectionId == nil {
		id := uuid.New()
		collectionId = &id
	}

	if len([]rune(req.Name)) > maxCollectionNameLength {
		req.Name = string([]rune(req.Name)[0:maxCollectionNameLength])
	}

	if !slices.Contains(model.AvailableVisibilities, req.Visibility) {
		req.Visibility = model.VisibilityPrivate
	}

	return entity.CollectionInput{
		Id:         *collectionId,
		UserId:     userId,
		Name:       req.Name,
		Visibility: req.Visibility,
	}, nil
}

func NewUpdateCollectionInput(req *api.UpdateCollectionRequest) (entity.CollectionInput, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return entity.CollectionInput{}, fail.GrpcInvalidBody
	}
	collectionId, err := uuid.Parse(req.CollectionId)
	if err != nil {
		return entity.CollectionInput{}, fail.GrpcInvalidBody
	}

	if len(req.Name) == 0 {
		return entity.CollectionInput{}, fail.GrpcInvalidBody
	}

	if len([]rune(req.Name)) > maxCollectionNameLength {
		req.Name = string([]rune(req.Name)[0:maxCollectionNameLength])
	}

	if !slices.Contains(model.AvailableVisibilities, req.Visibility) {
		req.Visibility = model.VisibilityPrivate
	}

	return entity.CollectionInput{
		Id:         collectionId,
		UserId:     userId,
		Name:       req.Name,
		Visibility: req.Visibility,
	}, nil
}
