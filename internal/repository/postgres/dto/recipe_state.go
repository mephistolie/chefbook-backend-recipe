package dto

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
)

type RecipeState struct {
	Id      uuid.UUID `db:"recipe_id"`
	Version int32     `db:"version"`

	OwnerId uuid.UUID `db:"owner_id"`

	Translations RecipeTranslations `db:"translations"`

	Rating float32 `db:"rating"`
	Score  int32   `db:"score"`
	Votes  int32   `db:"votes"`

	Tags        []string      `db:"tags"`
	Collections CollectionIds `db:"collections"`
	IsFavourite bool          `db:"favourite"`
}

func (r *RecipeState) Entity() entity.RecipeState {
	var collections []uuid.UUID
	for _, collection := range r.Collections {
		collections = append(collections, collection.UUID)
	}

	var score *int32
	if r.Score > 0 {
		score = &r.Score
	}

	return entity.RecipeState{
		Id:      r.Id,
		Version: r.Version,

		OwnerId: r.OwnerId,

		Translations: TranslationsEntity(r.Translations),

		Rating: r.Rating,
		Score:  score,
		Votes:  r.Votes,

		Tags:        r.Tags,
		Collections: collections,
		IsFavourite: r.IsFavourite,
	}
}
