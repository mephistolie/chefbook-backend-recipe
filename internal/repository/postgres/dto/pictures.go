package dto

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
)

type RecipePictures struct {
	Preview *uuid.UUID                 `json:"preview,omitempty"`
	Cooking *map[uuid.UUID][]uuid.UUID `json:"cooking,omitempty"`
}

func NewRecipePicturesDto(entity entity.RecipePictures) RecipePictures {
	var cooking *map[uuid.UUID][]uuid.UUID
	if len(entity.Cooking) > 0 {
		cooking = &entity.Cooking
	}

	return RecipePictures{
		Preview: entity.Preview,
		Cooking: cooking,
	}
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

type RecipePicturesUpload []uuid.UUID

func (p RecipePicturesUpload) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *RecipePicturesUpload) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return errors.New("unable to unmarshal recipe pictures")
	}

	return nil
}
