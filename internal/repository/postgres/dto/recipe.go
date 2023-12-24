package dto

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	"time"
)

type Recipe struct {
	Id   uuid.UUID `db:"recipe_id"`
	Name string    `db:"name"`

	OwnerId uuid.UUID `db:"owner_id"`

	Visibility  string `db:"visibility"`
	IsEncrypted bool   `db:"encrypted"`

	Language     string                  `db:"language"`
	Translations []RecipeTranslationInfo `db:"translations"`
	Description  *string                 `db:"description"`

	Rating float32 `db:"rating"`
	Score  int32   `db:"score"`
	Votes  int32   `db:"votes"`

	Tags []string `db:"tags"`

	IsSaved     bool       `db:"isSaved"`
	IsFavourite bool       `db:"isFavourite"`
	Categories  Categories `db:"categories"`

	Ingredients Ingredients    `db:"ingredients"`
	Cooking     Cooking        `db:"cooking"`
	Pictures    RecipePictures `db:"pictures"`

	Servings *int32 `db:"servings"`
	Time     *int32 `db:"cooking_time"`

	Calories      *int32 `db:"calories"`
	Protein       *int32 `db:"protein"`
	Fats          *int32 `db:"fats"`
	Carbohydrates *int32 `db:"carbohydrates"`

	CreationTimestamp time.Time `db:"creation_timestamp"`
	UpdateTimestamp   time.Time `db:"update_timestamp"`
	Version           int32     `db:"version"`
}

func (r *Recipe) Entity(userId uuid.UUID) entity.BaseRecipe {
	var macronutrients *entity.Macronutrients
	if r.Protein != nil || r.Fats != nil || r.Carbohydrates != nil {
		macronutrients = &entity.Macronutrients{
			Protein:       r.Protein,
			Fats:          r.Fats,
			Carbohydrates: r.Carbohydrates,
		}
	}
	var score *int32
	if r.Score > 0 {
		score = &r.Score
	}

	return entity.BaseRecipe{
		Id:   r.Id,
		Name: r.Name,

		OwnerId: r.OwnerId,

		IsOwned:     userId == r.OwnerId,
		IsSaved:     r.IsSaved,
		Visibility:  r.Visibility,
		IsEncrypted: r.IsEncrypted,

		Language:     r.Language,
		Translations: r.TranslationsEntity(),
		Description:  r.Description,

		CreationTimestamp: r.CreationTimestamp,
		UpdateTimestamp:   r.UpdateTimestamp,
		Version:           r.Version,

		Rating: r.Rating,
		Score:  score,
		Votes:  r.Votes,

		Tags:        r.Tags,
		Categories:  r.Categories,
		IsFavourite: r.IsFavourite,

		Servings: r.Servings,
		Time:     r.Time,

		Calories:       r.Calories,
		Macronutrients: macronutrients,

		Ingredients: r.Ingredients.Entity(),
		Cooking:     r.Cooking.Entity(),
		Pictures:    r.Pictures.Entity(),
	}
}

func (r *Recipe) TranslationsEntity() map[string][]entity.RecipeTranslationInfo {
	translations := map[string][]entity.RecipeTranslationInfo{}
	for _, translation := range r.Translations {
		languageTranslations, _ := translations[translation.Language]
		languageTranslations = append(languageTranslations, entity.RecipeTranslationInfo{AuthorId: translation.AuthorId})
		translations[translation.Language] = languageTranslations
	}
	return translations
}
