package postgres

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"github.com/mephistolie/chefbook-backend-recipe/api/model"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	"github.com/mephistolie/chefbook-backend-recipe/internal/repository/postgres/dto"
)

func (r *Repository) GetRecipeBook(userId uuid.UUID) ([]entity.RecipeState, error) {
	var recipes []entity.RecipeState

	query := fmt.Sprintf(`
		SELECT
			%[1]v.recipe_id,
			%[1]v.owner_id,
			(
				SELECT COALESCE(
					jsonb_agg(json_build_object('language', %[5]v.language, 'author_id', %[5]v.author_id))
					FILTER (WHERE %[5]v.author_id IS NOT NULL),
					'[]'
				)
				FROM %[5]v
				WHERE %[5]v.recipe_id=%[1]v.recipe_id
			) AS translations,
			%[1]v.rating, %[1]v.votes, COALESCE(%[4]v.score, 0),
			%[1]v.tags,
			(%[6]v) as collections,
			(
				SELECT EXISTS
				(
					SELECT 1
					FROM %[3]v
					WHERE %[3]v.recipe_id=%[3]v.recipe_id AND user_id=$1
				)
			) AS favoutie,
			%[1]v.version
		FROM
			%[2]v
		LEFT JOIN
			%[1]v ON %[1]v.recipe_id=%[2]v.recipe_id AND user_id=$1
		LEFT JOIN
			%[4]v ON %[4]v.recipe_id=%[2]v.recipe_id AND user_id=$1
		WHERE
			%[2]v.user_id=$1 AND (%[1]v.owner_id=$1 OR %[1]v.visibility<>'%[7]v')
	`, recipesTable, recipeBookTable, favouritesTable, scoresTable, translationsTable,
		getRecipeCollectionIdsSubquery, model.VisibilityPrivate)

	rows, err := r.db.Query(query, userId)
	if err != nil {
		log.Errorf("unable to get recipes: %s", query)
		return []entity.RecipeState{}, fail.GrpcUnknown
	}

	for rows.Next() {
		recipe := dto.RecipeState{}
		m := pgtype.NewMap()
		if err = rows.Scan(
			&recipe.Id,
			&recipe.OwnerId,
			m.SQLScanner(&recipe.Translations),
			&recipe.Rating, &recipe.Votes, &recipe.Score,
			m.SQLScanner(&recipe.Tags), m.SQLScanner(&recipe.Collections), &recipe.IsFavourite,
			&recipe.Version,
		); err != nil {
			log.Warnf("unable to parse recipe info: ", err)
			continue
		}
		recipes = append(recipes, recipe.Entity())
	}

	return recipes, nil
}

func (r *Repository) SaveRecipeToRecipeBook(recipeId, userId uuid.UUID) error {
	query := fmt.Sprintf(`
			INSERT INTO %s (recipe_id, user_id)
			VALUES ($1, $2)
			ON CONFLICT (recipe_id, user_id) DO NOTHING
		`, recipeBookTable)

	if _, err := r.db.Exec(query, recipeId, userId); err != nil {
		if isUniqueViolationError(err) {
			return nil
		}
		log.Errorf("unable to add recipe %s to user %s recipe book: %s", recipeId, userId, err)
		return fail.GrpcNotFound
	}

	return nil
}

func (r *Repository) RemoveRecipeFromRecipeBook(recipeId, userId uuid.UUID) error {
	query := fmt.Sprintf(`
			DELETE FROM %s
			WHERE recipe_id=$1 AND user_id=$2
		`, recipeBookTable)

	if _, err := r.db.Exec(query, recipeId, userId); err != nil {
		log.Errorf("unable to delete recipe %s from user %s recipe book: %s", recipeId, userId, err)
		return fail.GrpcNotFound
	}

	return nil
}

func (r *Repository) SaveRecipeToFavourites(recipeId, userId uuid.UUID) error {
	query := fmt.Sprintf(`
			INSERT INTO %s (recipe_id, user_id)
			VALUES ($1, $2)
			ON CONFLICT (recipe_id, user_id) DO NOTHING
		`, favouritesTable)

	if _, err := r.db.Exec(query, recipeId, userId); err != nil {
		if isUniqueViolationError(err) {
			return nil
		}
		log.Errorf("unable to add recipe %s to user %s favourites: %s", recipeId, userId, err)
		return fail.GrpcNotFound
	}

	return nil
}

func (r *Repository) RemoveRecipeFromFavourites(recipeId, userId uuid.UUID) error {
	query := fmt.Sprintf(`
			DELETE FROM %s
			WHERE recipe_id=$1 AND user_id=$2
		`, favouritesTable)

	if _, err := r.db.Exec(query, recipeId, userId); err != nil {
		log.Errorf("unable to delete recipe %s from user %s favourites: %s", recipeId, userId, err)
		return fail.GrpcNotFound
	}

	return nil
}
