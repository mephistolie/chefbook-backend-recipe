package dto

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
)

type Collection struct {
	Id   uuid.UUID
	Name string

	Visibility   string
	Contributors Contributors

	RecipesCount int32
}

func (c *Collection) Entity() entity.Collection {
	return entity.Collection{
		Id:   c.Id,
		Name: c.Name,

		Contributors: c.Contributors.Entity(),
		Visibility:   c.Visibility,

		RecipesCount: c.RecipesCount,
	}
}

type CollectionIds []CollectionId

func (c CollectionIds) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *CollectionIds) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	if err := json.Unmarshal(b, &c); err != nil {
		return errors.New("unable to unmarshal collection IDs")
	}

	return nil
}

type CollectionId struct {
	uuid.UUID
}

func (c CollectionId) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *CollectionId) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	if err := json.Unmarshal(b, &c); err != nil {
		return errors.New("unable to unmarshal collection ID")
	}

	return nil
}
