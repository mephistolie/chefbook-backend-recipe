package postgres

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"github.com/mephistolie/chefbook-backend-recipe/api/model"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	"github.com/mephistolie/chefbook-backend-recipe/internal/repository/postgres/dto"
)

func (r *Repository) GetRecipeBook(userId uuid.UUID) ([]entity.BaseRecipeState, error) {
	var recipes []entity.BaseRecipeState

	query := fmt.Sprintf(`
		SELECT
			%[1]v.recipe_id,
			%[1]v.owner_id,
			%[1]v.rating, %[1]v.votes, coalesce(%[3]v.score, 0),
			%[1]v.tags, coalesce(%[2]v.categories, '[]'::jsonb), coalesce(%[2]v.favourite, false),
			%[1]v.version
		FROM
			%[1]v
		LEFT JOIN
			%[2]v ON %[2]v.recipe_id=%[1]v.recipe_id
		LEFT JOIN
			%[3]v ON %[3]v.recipe_id=%[1]v.recipe_id
		WHERE
			%[2]v.user_id=$1 AND (%[1]v.owner_id=$1 OR %[1]v.visibility<>'%[4]v')
	`, recipesTable, usersTable, scoresTable, model.VisibilityPrivate)

	rows, err := r.db.Query(query, userId)
	if err != nil {
		log.Errorf("unable to get recipes: %s", err)
		return []entity.BaseRecipeState{}, fail.GrpcUnknown
	}

	for rows.Next() {
		recipe := dto.RecipeState{}
		if err = rows.Scan(
			&recipe.Id,
			&recipe.OwnerId,
			&recipe.Rating, &recipe.Votes, &recipe.Score,
			&recipe.Tags, &recipe.Categories, &recipe.IsFavourite,
			&recipe.Version,
		); err != nil {
			log.Warnf("unable to parse recipe info: ", err)
			continue
		}
		recipes = append(recipes, recipe.Entity())
	}

	return recipes, nil
}

func (r *Repository) SaveToRecipeBook(recipeId, userId uuid.UUID) error {
	query := fmt.Sprintf(`
			INSERT INTO %s (recipe_id, user_id)
			VALUES ($1, $2)
		`, usersTable)

	if _, err := r.db.Exec(query, recipeId, userId); err != nil {
		if isUniqueViolationError(err) {
			return nil
		}
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
