package postgres

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"github.com/mephistolie/chefbook-backend-recipe/internal/repository/postgres/dto"
)

func (r *Repository) SaveToRecipeBook(recipeId, userId uuid.UUID) error {
	query := fmt.Sprintf(`
			INSERT INTO %s (recipe_id, user_id)
			VALUES ($1, $2)
		`, usersTable)

	if _, err := r.db.Exec(query, recipeId, userId); err != nil {
		log.Errorf("unable to add recipe %s to user %s recipe book: %s", recipeId, userId, err)
		return fail.GrpcNotFound
	}

	return nil
}

func (r *Repository) RemoveFromRecipeBook(recipeId, userId uuid.UUID) error {
	query := fmt.Sprintf(`
			DELETE FROM %s
			WHERE recipe_id=$1 AND user_id=$2
		`, usersTable)

	if _, err := r.db.Exec(query, recipeId, userId); err != nil {
		log.Errorf("unable to delete recipe %s from user %s recipe book: %s", recipeId, userId, err)
		return fail.GrpcNotFound
	}

	return nil
}

func (r *Repository) SetRecipeFavouriteStatus(recipeId, userId uuid.UUID, isFavourite bool) error {
	query := fmt.Sprintf(`
			UPDATE %s
			SET favourite=$1
			WHERE recipe_id=$2 AND user_id=$3
		`, usersTable)

	if _, err := r.db.Exec(query, isFavourite, recipeId, userId); err != nil {
		log.Errorf("unable to change recipe %s favourite status for user %s recipe book: %s", recipeId, userId, err)
		return fail.GrpcNotFound
	}

	return nil
}

func (r *Repository) SetRecipeCategories(recipeId, userId uuid.UUID, categories []uuid.UUID) error {
	query := fmt.Sprintf(`
			UPDATE %s
			SET categories=$1
			WHERE recipe_id=$2 AND user_id=$3
		`, usersTable)

	if _, err := r.db.Exec(query, dto.NewCategories(categories), recipeId, userId); err != nil {
		log.Errorf("unable to set recipe %s categories for user %s recipe book: %s", recipeId, userId, err)
		return fail.GrpcNotFound
	}

	return nil
}
