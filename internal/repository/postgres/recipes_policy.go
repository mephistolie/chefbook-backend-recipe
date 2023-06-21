package postgres

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
)

func (r *Repository) GetRecipePolicy(recipeId uuid.UUID) (entity.RecipePolicy, error) {
	var policy entity.RecipePolicy

	query := fmt.Sprintf(`
		SELECT recipe_id, owner_id, visibility, encrypted
		FROM %s
		WHERE recipe_id=$1
	`, recipesTable)

	row := r.db.QueryRow(query, recipeId)
	if err := row.Scan(&policy.Id, &policy.OwnerId, &policy.Visibility, &policy.IsEncrypted); err != nil {
		log.Warnf("unable to get recipe %s policy: %s", recipeId, err)
		return entity.RecipePolicy{}, fail.GrpcNotFound
	}

	return entity.RecipePolicy{}, nil
}
