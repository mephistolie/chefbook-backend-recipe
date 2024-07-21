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
	Saved        bool

	RecipesCount int
}

func (c *Collection) Entity() entity.Collection {
	return entity.Collection{
		Id:   c.Id,
		Name: c.Name,

		Contributors: c.Contributors.Entity(),
		Visibility:   c.Visibility,
		Saved:        c.Saved,

		RecipesCount: c.RecipesCount,
	}
}

type Collections []uuid.UUID

func (c Collections) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *Collections) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	if err := json.Unmarshal(b, &c); err != nil {
		return errors.New("unable to unmarshal category IDs")
	}

	return nil
}
