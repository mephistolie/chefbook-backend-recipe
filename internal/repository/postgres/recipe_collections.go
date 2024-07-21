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
	SELECT collection_id
	FROM %[2]v
	LEFT JOIN
		%[3]v ON %[3]v.collection_id=%[2]v.collection_id
	LEFT JOIN
		%[4]v ON %[4]v.collection_id=%[2]v.collection_id
	LEFT JOIN
		%[5]v ON %[5]v.collection_id=%[2]v.collection_id
	WHERE
		%[2]v.recipe_id=%[1]v.recipe_id AND %[5]v.user_id=$2 AND (%[3]v.visibility<>'%[6]v' OR %[2]v.contributor_id=$2 AND role='%[7]v')
`, recipesTable, recipesCollectionsTable, collectionsTable, collectionContributorsTable, collectionUsersTable,
	model.VisibilityPrivate, entity.RoleOwner)

func (r *Repository) SetRecipeCollections(recipeId, userId uuid.UUID, collections []uuid.UUID) error {
	tx, err := r.startTransaction()
	if err != nil {
		return err
	}

	editableCollections, err := r.getEditableCollections(tx, userId)
	if err != nil {
		return errorWithTransactionRollback(tx, fail.GrpcNotFound)
	}

	clearCollectionsQuery := fmt.Sprintf(`
		DELETE FROM %s
		WHERE recipe_id=$1 AND collection_id=ANY($2)
	`, recipeUsersTable)

	if _, err := tx.Exec(clearCollectionsQuery, recipeId, editableCollections); err != nil {
		log.Errorf("unable to clear recipe %s collections for user %s: %s", recipeId, userId, err)
		return errorWithTransactionRollback(tx, fail.GrpcNotFound)
	}

	editableCollectionsMap := make(map[uuid.UUID]bool)
	for _, collectionId := range editableCollections {
		editableCollectionsMap[collectionId] = true
	}

	setRecipeCollectionsQuery := fmt.Sprintf(`
		INSERT INTO %[1]v (recipe_id, collection_id)
		VALUES `, collectionContributorsTable)

	var args []interface{}
	for _, collectionId := range collections {
		if _, ok := editableCollectionsMap[collectionId]; !ok {
			continue
		}
		setRecipeCollectionsQuery += "(?, ?),"
		args = append(args, recipeId, collectionId)
	}
	setRecipeCollectionsQuery = setRecipeCollectionsQuery[0 : len(setRecipeCollectionsQuery)-1]
	setRecipeCollectionsQuery += fmt.Sprint(" ON CONFLICT DO NOTHING")

	if _, err := tx.Exec(setRecipeCollectionsQuery, args); err != nil {
		log.Errorf("unable to set recipe %s collections for user %s: %s", recipeId, userId, err)
		return errorWithTransactionRollback(tx, fail.GrpcNotFound)
	}

	return nil
}

func (r *Repository) getEditableCollections(tx *sql.Tx, userId uuid.UUID) ([]uuid.UUID, error) {
	var getEditableCollectionsQuery = fmt.Sprintf(`
		SELECT collection_id
		FROM %[1]v
		LEFT JOIN
			%[2]v ON %[2]v.collection_id=%[1]v.collection_id
		WHERE %[1]v.contributor_id=$1 AND %[2]v.user_id=$1
	`, collectionContributorsTable, collectionUsersTable)

	rows, err := tx.Query(getEditableCollectionsQuery, userId)
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
