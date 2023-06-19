package postgres

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
)

func (r *Repository) GetRecipeOwner(recipeId uuid.UUID) (uuid.UUID, error) {
	var id uuid.UUID

	query := fmt.Sprintf(`
		SELECT owner_id
		FROM %s
		WHERE recipe_id=$1
	`, recipesTable)

	if err := r.db.Get(&id, query, recipeId); err != nil {
		log.Warnf("unable to get recipe %s owner", recipeId, err)
		return uuid.UUID{}, fail.GrpcNotFound
	}

	return id, nil
}
