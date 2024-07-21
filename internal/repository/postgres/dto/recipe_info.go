package dto

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	"time"
)

type RecipeInfo struct {
	Id   uuid.UUID `db:"recipe_id"`
	Name string    `db:"name"`

	OwnerId uuid.UUID `db:"owner_id"`

	IsSaved     bool   `db:"saved"`
	Visibility  string `db:"visibility"`
	IsEncrypted bool   `db:"encrypted"`

	Language     string   `db:"language"`
	Translations []string `db:"translations"`

	Rating float32 `db:"rating"`
	Score  int32   `db:"score"`
	Votes  int32   `db:"votes"`

	Tags []string `db:"tags"`

	IsFavourite bool        `db:"favourite"`
	Collections Collections `db:"collections"`

	Pictures RecipePictures `db:"pictures"`

	Servings *int32 `db:"servings"`
	Time     *int32 `db:"cooking_time"`

	Calories *int32 `db:"calories"`

	CreationTimestamp time.Time `db:"creation_timestamp"`
	UpdateTimestamp   time.Time `db:"update_timestamp"`
	Version           int32     `db:"version"`
}

func (r *RecipeInfo) Entity(userId uuid.UUID) entity.RecipeInfo {
	var score *int32
	if r.Score > 0 {
		score = &r.Score
	}

	return entity.RecipeInfo{
		Id:   r.Id,
		Name: r.Name,

		OwnerId: r.OwnerId,

		IsOwned:     userId == r.OwnerId,
		IsSaved:     r.IsSaved,
		Visibility:  r.Visibility,
		IsEncrypted: r.IsEncrypted,

		Language:     r.Language,
		Translations: r.Translations,
		Preview:      r.Pictures.Preview,

		CreationTimestamp: r.CreationTimestamp,
		UpdateTimestamp:   r.UpdateTimestamp,
		Version:           r.Version,

		Rating: r.Rating,
		Score:  score,
		Votes:  r.Votes,

		Tags:        r.Tags,
		Collections: r.Collections,
		IsFavourite: r.IsFavourite,

		Servings: r.Servings,
		Time:     r.Time,

		Calories: r.Calories,
	}
}
