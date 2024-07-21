package entity

import "github.com/google/uuid"

type Collection struct {
	Id   uuid.UUID
	Name string

	Visibility   string
	Contributors []Contributor
	Saved        bool

	RecipesCount int
}

type DetailedCollection struct {
	Collection   Collection
	ProfilesInfo map[string]ProfileInfo
}

type DetailedCollections struct {
	Collections  []Collection
	ProfilesInfo map[string]ProfileInfo
}

type CollectionInfo struct {
	Name string
}

type CollectionInput struct {
	Id         uuid.UUID
	UserId     uuid.UUID
	Name       string
	Visibility string
}
