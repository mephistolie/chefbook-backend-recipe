package postgres

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
	"github.com/mephistolie/chefbook-backend-recipe/internal/repository/postgres/dto"
	"k8s.io/utils/strings/slices"
)

func (r *Repository) GetRecipeTranslations(recipeId uuid.UUID) (map[string][]uuid.UUID, error) {
	query := fmt.Sprintf(`
		SELECT language, author_id
		FROM %s
		WHERE recipe_id=$1 AND hidden=false
	`, translationsTable)

	rows, err := r.db.Query(query, recipeId)
	if err != nil {
		log.Warnf("unable to get recipe %s translations: %s", recipeId, err)
		return nil, fail.GrpcNotFound
	}
	var translations []dto.RecipeTranslationInfo
	for rows.Next() {
		translation := dto.RecipeTranslationInfo{}
		if err = rows.Scan(&translation.Language, &translation.AuthorId); err == nil {
			translations = append(translations, translation)
		}
	}

	return dto.TranslationsEntity(translations), nil
}

func (r *Repository) GetRecipeTranslation(recipeId uuid.UUID, language string, authorId *uuid.UUID) *entity.RecipeTranslation {
	var translations []dto.RecipeTranslation

	query := fmt.Sprintf(`
		SELECT author_id, name, description, ingredients, cooking
		FROM %s
		WHERE recipe_id=$1 AND language=$2 AND hidden=false
	`, translationsTable)

	args := []interface{}{recipeId, language}
	if authorId != nil {
		query += " AND author_id=$3"
		args = append(args, *authorId)
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		log.Warnf("unable to get recipe %s translation to %s: %s", recipeId, language, err)
		return nil
	}
	for rows.Next() {
		translation := dto.RecipeTranslation{Language: language}
		if err = rows.Scan(&translation.AuthorId, &translation.Name, &translation.Description, &translation.Ingredients,
			&translation.Cooking); err != nil {
			log.Errorf("unable to parse recipe %s translation to %s; %s", recipeId, language, err)
			continue
		}
		if authorId != nil && *authorId == translation.AuthorId {
			_ = rows.Close()
			t := translation.Entity()
			return &t
		}

		translations = append(translations, translation)
	}

	if len(translations) == 0 {
		return nil
	}

	t := translations[0].Entity()
	return &t
}

func (r *Repository) TranslateRecipe(recipeId uuid.UUID, translation entity.RecipeTranslation) error {
	t := dto.NewRecipeTranslation(translation)

	tx, err := r.startTransaction()
	if err != nil {
		return err
	}

	addTranslationQuery := fmt.Sprintf(`
		INSERT INTO %s (recipe_id, language, author_id, name, description, ingredients, cooking)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (recipe_id, author_id) DO UPDATE
		SET author_id=$3, name=$4, description=$5, ingredients=$6, cooking=$7;
	`, translationsTable)

	if _, err = tx.Exec(addTranslationQuery, recipeId, t.Language, translation.AuthorId, t.Name, t.Description, t.Ingredients, t.Cooking); err != nil {
		log.Errorf("unable to add recipe %s translation to %s: %s", recipeId, translation.Language, err)
		return errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}

	getCurrentRecipeTranslationsQuery := fmt.Sprintf(`
		SELECT translations
		FROM %s
		WHERE recipe_id=$1
	`, recipesTable)

	var languages []string

	row := tx.QueryRow(getCurrentRecipeTranslationsQuery, recipeId)
	m := pgtype.NewMap()
	if err = row.Scan(m.SQLScanner(&languages)); err != nil {
		log.Warnf("unable to get recipe %s languages: %s", recipeId, err)
		return errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}

	if !slices.Contains(languages, translation.Language) {
		languages = append(languages, translation.Language)
		updateRecipeTranslationsQuery := fmt.Sprintf(`
			UPDATE %s
			SET translations=$2
			WHERE recipe_id=$1
		`, recipesTable)

		if _, err = tx.Exec(updateRecipeTranslationsQuery, recipeId, languages); err != nil {
			log.Errorf("unable to update recipe %s translations: %s", recipeId, err)
			return errorWithTransactionRollback(tx, fail.GrpcUnknown)
		}
	}

	return commitTransaction(tx)
}

func (r *Repository) DeleteRecipeTranslation(recipeId uuid.UUID, userId uuid.UUID, language string) error {
	tx, err := r.startTransaction()
	if err != nil {
		return err
	}

	deleteTranslationQuery := fmt.Sprintf(`
		DELETE FROM %s
		WHERE recipe_id=$1 AND author_id=$2 and language=$3
	`, translationsTable)

	if _, err = tx.Exec(deleteTranslationQuery, recipeId, userId, language); err != nil {
		log.Errorf("unable to delete recipe %s translation to %s: %s", recipeId, language, err)
		return errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}

	getRecipeTranslationsCountQuery := fmt.Sprintf(`
		SELECT COUNT(*)
		FROM %s
		WHERE recipe_id=$1 AND language=$2
	`, recipesTable)

	var translationsCount int

	row := tx.QueryRow(getRecipeTranslationsCountQuery, recipeId, language)
	if err = row.Scan(&translationsCount); err != nil {
		log.Warnf("unable to get recipe %s translations count: %s", recipeId, err)
		return errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}

	if translationsCount == 0 {
		updateRecipeTranslationsQuery := fmt.Sprintf(`
			UPDATE %s
			SET translations=array_remove(translations, $2)
			WHERE recipe_id=$1
		`, recipesTable)

		if _, err = tx.Exec(updateRecipeTranslationsQuery, recipeId, language); err != nil {
			log.Errorf("unable to update recipe %s translations: %s", recipeId, err)
			return errorWithTransactionRollback(tx, fail.GrpcUnknown)
		}
	}

	return commitTransaction(tx)
}
