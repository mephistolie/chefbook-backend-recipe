package dto

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
)

type RecipeState struct {
	Id      uuid.UUID `db:"recipe_id"`
	Version int32     `db:"version"`

	OwnerId uuid.UUID `db:"owner_id"`

	Rating float32 `db:"rating"`
	Score  int32   `db:"score"`
	Votes  int32   `db:"votes"`

	Tags        Tags       `db:"tags"`
	Categories  Categories `db:"categories"`
	IsFavourite bool       `db:"favourite"`
}

func (r *RecipeState) Entity() entity.BaseRecipeState {
	var score *int32
	if r.Score > 0 {
		score = &r.Score
	}

	return entity.BaseRecipeState{
		Id:      r.Id,
		Version: r.Version,

		OwnerId: r.OwnerId,

		Rating: r.Rating,
		Score:  score,
		Votes:  r.Votes,

		Tags:        r.Tags,
		Categories:  r.Categories,
		IsFavourite: r.IsFavourite,
	}
}
