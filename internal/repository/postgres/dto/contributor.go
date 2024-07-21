package dto

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
)

type Contributors []Contributor

func (c *Contributors) Entity() []entity.Contributor {
	var contributors []entity.Contributor
	for _, dto := range *c {
		contributor := entity.Contributor{
			Id:   dto.Id,
			Role: dto.Role,
		}
		contributors = append(contributors, contributor)
	}
	return contributors
}

func (c Contributors) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *Contributors) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	if err := json.Unmarshal(b, &c); err != nil {
		return errors.New("unable to unmarshal contributors")
	}

	return nil
}

type Contributor struct {
	Id   uuid.UUID `json:"contributor_id"`
	Role string    `json:"role"`
}

func (c Contributor) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *Contributor) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	if err := json.Unmarshal(b, &c); err != nil {
		return errors.New("unable to unmarshal contributor")
	}

	return nil
}
