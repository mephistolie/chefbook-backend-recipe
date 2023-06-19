package dto

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
)

type RecipePictures struct {
	Preview *string                 `json:"preview,omitempty"`
	Cooking *map[uuid.UUID][]string `json:"cooking,omitempty"`
}

func (p RecipePictures) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *RecipePictures) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return errors.New("unable to unmarshal recipe pictures")
	}

	return nil
}
