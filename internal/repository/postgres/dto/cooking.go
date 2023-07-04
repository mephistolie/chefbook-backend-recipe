package dto

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
)

type Cooking []CookingItem

func (c *Cooking) Entity(pictures RecipePictures) []entity.CookingItem {
	picturesMap := make(map[uuid.UUID][]uuid.UUID)
	if pictures.Cooking != nil {
		picturesMap = *pictures.Cooking
	}

	var cooking []entity.CookingItem
	for _, dto := range *c {
		var stepPictures []uuid.UUID
		if dto.Type == entity.TypeStep {
			stepPictures, _ = picturesMap[dto.Id]
		}

		step := entity.CookingItem{
			Id:         dto.Id,
			Text:       dto.Text,
			Type:       dto.Type,
			Time:       dto.Time,
			PictureIds: stepPictures,
			RecipeId:   dto.RecipeId,
		}
		cooking = append(cooking, step)
	}
	return cooking
}

func NewCooking(cooking []entity.CookingItem) Cooking {
	dtos := make(Cooking, len(cooking))
	for i, ingredient := range cooking {
		dto := CookingItem{
			Id:       ingredient.Id,
			Text:     ingredient.Text,
			Type:     ingredient.Type,
			Time:     ingredient.Time,
			RecipeId: ingredient.RecipeId,
		}
		dtos[i] = dto
	}
	return dtos
}

func (c Cooking) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *Cooking) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	if err := json.Unmarshal(b, &c); err != nil {
		return errors.New("unable to unmarshal cooking")
	}

	return nil
}

type CookingItem struct {
	Id       uuid.UUID  `json:"id"`
	Text     *string    `json:"text,omitempty"`
	Type     string     `json:"type"`
	Time     *int32     `json:"time,omitempty"`
	RecipeId *uuid.UUID `json:"recipeId,omitempty"`
}

func (i CookingItem) Value() (driver.Value, error) {
	return json.Marshal(i)
}

func (i *CookingItem) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	if err := json.Unmarshal(b, &i); err != nil {
		return errors.New("unable to unmarshal recipe cooking item")
	}

	return nil
}
