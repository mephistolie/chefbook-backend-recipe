package dto

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
)

type Categories []uuid.UUID

func NewCategories(categories []uuid.UUID) Categories {
	dtos := make([]uuid.UUID, len(categories))
	for i, category := range categories {
		dtos[i] = category
	}
	return dtos
}

func (c Categories) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *Categories) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	if err := json.Unmarshal(b, &c); err != nil {
		return errors.New("unable to unmarshal category IDs")
	}

	return nil
}
