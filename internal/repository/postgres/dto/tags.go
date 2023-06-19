package dto

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Tags []string

func NewTags(tags []string) Tags {
	dtos := make([]string, len(tags))
	for i, tag := range tags {
		dtos[i] = tag
	}
	return dtos
}

func (t Tags) Value() (driver.Value, error) {
	return json.Marshal(t)
}

func (t *Tags) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	if err := json.Unmarshal(b, &t); err != nil {
		return errors.New("unable to unmarshal tag IDs")
	}

	return nil
}
