package dto

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
)

type RecipeTranslationInfo struct {
	Language string    `db:"language" json:"language"`
	AuthorId uuid.UUID `db:"author_id" json:"authorId"`
}

type RecipeTranslations []RecipeTranslationInfo

func (c RecipeTranslations) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *RecipeTranslations) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	if err := json.Unmarshal(b, &c); err != nil {
		return errors.New("unable to unmarshal translation IDs")
	}

	return nil
}

func TranslationsEntity(dto RecipeTranslations) map[string][]uuid.UUID {
	translations := map[string][]uuid.UUID{}
	for _, translation := range dto {
		languageTranslations, _ := translations[translation.Language]
		languageTranslations = append(languageTranslations, translation.AuthorId)
		translations[translation.Language] = languageTranslations
	}
	return translations
}

type IngredientTranslation struct {
	Text string  `json:"text"`
	Unit *string `json:"unit"`
}

type IngredientsTranslation map[uuid.UUID]IngredientTranslation

func (i *IngredientsTranslation) Entity() map[uuid.UUID]entity.IngredientTranslation {
	ingredients := map[uuid.UUID]entity.IngredientTranslation{}
	for id, dto := range *i {
		ingredient := entity.IngredientTranslation{
			Text: dto.Text,
			Unit: dto.Unit,
		}
		ingredients[id] = ingredient
	}
	return ingredients
}

func NewIngredientsTranslation(translation map[uuid.UUID]entity.IngredientTranslation) IngredientsTranslation {
	dtos := IngredientsTranslation{}
	for id, ingredient := range translation {
		dto := IngredientTranslation{
			Text: ingredient.Text,
			Unit: ingredient.Unit,
		}
		dtos[id] = dto
	}
	return dtos
}

func (i IngredientsTranslation) Value() (driver.Value, error) {
	return json.Marshal(i)
}

func (i *IngredientsTranslation) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	if err := json.Unmarshal(b, &i); err != nil {
		return errors.New("unable to unmarshal recipe ingredients translation")
	}

	return nil
}

type CookingTranslation map[uuid.UUID]string

func (c CookingTranslation) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *CookingTranslation) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	if err := json.Unmarshal(b, &c); err != nil {
		return errors.New("unable to unmarshal recipe cooking translation")
	}

	return nil
}

type RecipeTranslation struct {
	AuthorId    uuid.UUID              `db:"author_id"`
	Language    string                 `db:"language"`
	Name        string                 `db:"name"`
	Description *string                `db:"description"`
	Ingredients IngredientsTranslation `db:"ingredients"`
	Cooking     CookingTranslation     `db:"cooking"`
}

func (t *RecipeTranslation) Entity() entity.RecipeTranslation {
	return entity.RecipeTranslation{
		AuthorId:    t.AuthorId,
		Language:    t.Language,
		Name:        t.Name,
		Description: t.Description,
		Ingredients: t.Ingredients.Entity(),
		Cooking:     t.Cooking,
	}
}

func NewRecipeTranslation(translation entity.RecipeTranslation) RecipeTranslation {
	return RecipeTranslation{
		AuthorId:    translation.AuthorId,
		Language:    translation.Language,
		Name:        translation.Name,
		Description: translation.Description,
		Ingredients: NewIngredientsTranslation(translation.Ingredients),
		Cooking:     translation.Cooking,
	}
}
