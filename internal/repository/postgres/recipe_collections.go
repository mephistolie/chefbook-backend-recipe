package postgres

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"github.com/mephistolie/chefbook-backend-recipe/api/model"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
)

var getRecipeCollectionIdsSubquery = fmt.Sprintf(`
	SELECT COALESCE(
		jsonb_agg(%[2]v.collection_id)
		FILTER (WHERE %[2]v.collection_id IS NOT NULL),
		'[]'
	)
	FROM %[2]v
	LEFT JOIN
		%[3]v ON %[3]v.collection_id=%[2]v.collection_id
	LEFT JOIN
		%[4]v ON %[4]v.collection_id=%[2]v.collection_id
	LEFT JOIN
		%[5]v ON %[5]v.collection_id=%[2]v.collection_id
	WHERE
		%[2]v.recipe_id=%[1]v.recipe_id AND %[5]v.user_id=$1 AND (%[3]v.visibility<>'%[6]v' OR %[4]v.contributor_id=$1 AND role='%[7]v')
`, recipesTable, recipesCollectionsTable, collectionsTable, collectionContributorsTable, collectionUsersTable,
	model.VisibilityPrivate, entity.RoleOwner)

func (r *Repository) AddRecipeToCollection(recipeId, collectionId, userId uuid.UUID) error {
	tx, err := r.startTransaction()
	if err != nil {
		return err
	}

	if !r.checkCollectionAccessible(tx, collectionId, userId) {
		return errorWithTransactionRollback(tx, fail.GrpcAccessDenied)
	}

	query := fmt.Sprintf(`
		INSERT INTO %s (recipe_id, collection_id)
		VALUES ($1, $2)
		ON CONFLICT (recipe_id, collection_id) DO NOTHING
	`, recipesCollectionsTable)

	if _, err := r.db.Exec(query, recipeId, collectionId); err != nil {
		log.Errorf("unable to add recipe %s to collection %s: %s", recipeId, collectionId, err)
		return errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}

	return commitTransaction(tx)
}

func (r *Repository) RemoveRecipeFromCollection(recipeId, collectionId, userId uuid.UUID) error {
	tx, err := r.startTransaction()
	if err != nil {
		return err
	}

	if !r.checkCollectionAccessible(tx, collectionId, userId) {
		return errorWithTransactionRollback(tx, fail.GrpcAccessDenied)
	}

	query := fmt.Sprintf(`
		DELETE FROM %[1]v
		WHERE recipe_id=$1 AND collection_id=$2
	`, recipesCollectionsTable)

	if _, err := r.db.Exec(query, recipeId, collectionId); err != nil {
		log.Errorf("unable to remove recipe %s from collection %s: %s", recipeId, collectionId, err)
		return errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}

	return commitTransaction(tx)
}

func (r *Repository) SetRecipeCollections(recipeId, userId uuid.UUID, collections []uuid.UUID) error {
	tx, err := r.startTransaction()
	if err != nil {
		return err
	}

	editableCollections, err := r.getEditableCollections(tx, userId)
	if err != nil {
		return errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}

	if len(editableCollections) == 0 {
		_ = tx.Rollback()
		return nil
	}

	clearCollectionsQuery := fmt.Sprintf(`
		DELETE FROM %s
		WHERE recipe_id=$1 AND collection_id=ANY($2)
	`, recipesCollectionsTable)

	if _, err := tx.Exec(clearCollectionsQuery, recipeId, editableCollections); err != nil {
		log.Errorf("unable to clear recipe %s collections for user %s: %s", recipeId, userId, err)
		return errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}

	editableCollectionsMap := make(map[uuid.UUID]bool)
	for _, collectionId := range editableCollections {
		editableCollectionsMap[collectionId] = true
	}

	setRecipeCollectionsQuery := fmt.Sprintf(`
		INSERT INTO %[1]v (recipe_id, collection_id)
		VALUES `, recipesCollectionsTable)

	var args []interface{}
	for _, collectionId := range collections {
		if _, ok := editableCollectionsMap[collectionId]; !ok {
			continue
		}
		currentIndex := len(args) + 1
		setRecipeCollectionsQuery += fmt.Sprintf("($%[1]v, $%[2]v),", currentIndex, currentIndex+1)
		args = append(args, recipeId, collectionId)
	}

	if len(args) == 0 {
		return commitTransaction(tx)
	}

	setRecipeCollectionsQuery = setRecipeCollectionsQuery[0 : len(setRecipeCollectionsQuery)-1]
	setRecipeCollectionsQuery += fmt.Sprint(" ON CONFLICT (recipe_id, collection_id) DO NOTHING")

	if _, err = tx.Exec(setRecipeCollectionsQuery, args...); err != nil {
		log.Errorf("unable to set recipe %s collections for user %s: %s", recipeId, userId, err)
		return errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}

	return commitTransaction(tx)
}

func (r *Repository) getEditableCollections(tx *sql.Tx, userId uuid.UUID) ([]uuid.UUID, error) {
	query := fmt.Sprintf(`
		SELECT %[1]v.collection_id
		FROM %[1]v
		LEFT JOIN
			%[2]v ON %[2]v.collection_id=%[1]v.collection_id
		WHERE %[1]v.contributor_id=$1 AND %[2]v.user_id=$1
	`, collectionContributorsTable, collectionUsersTable)

	rows, err := tx.Query(query, userId)
	if err != nil {
		log.Errorf("unable to get editable collections for user %s: %s", userId, err)
		return []uuid.UUID{}, fail.GrpcUnknown
	}

	var editableCollectionIds []uuid.UUID
	for rows.Next() {
		var collectionId uuid.UUID
		if err = rows.Scan(&collectionId); err != nil {
			log.Warnf("unable to parse collection id: %s", err)
			continue
		}
		editableCollectionIds = append(editableCollectionIds, collectionId)
	}

	return editableCollectionIds, nil
}

func (r *Repository) checkCollectionAccessible(tx *sql.Tx, collectionId, userId uuid.UUID) bool {
	var hasAccess bool

	query := fmt.Sprintf(`
		SELECT EXISTS
		(
			SELECT 1
			FROM %[1]v
			WHERE %[1]v.collection_id=$1 AND %[1]v.contributor_id=$2
		)
	`, collectionContributorsTable)

	row := tx.QueryRow(query, collectionId, userId)
	if err := row.Scan(&hasAccess); err != nil {
		log.Errorf("unable to check is user %s has access to collection %s: %s", userId, collectionId, err)
		return false
	}
	return hasAccess
}
