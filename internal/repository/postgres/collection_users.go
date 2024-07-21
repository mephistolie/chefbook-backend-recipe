package postgres

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
)

func (r *Repository) SaveCollectionToRecipeBook(collectionId, userId uuid.UUID) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (collection_id, user_id)
		VALUES ($1, $2)
	`, collectionUsersTable)

	if _, err := r.db.Exec(query, collectionId, userId); err != nil {
		if isUniqueViolationError(err) {
			return nil
		}
		log.Errorf("unable to add collection %s to user %s recipe book: %s", collectionId, userId, err)
		return fail.GrpcUnknown
	}

	return nil
}

func (r *Repository) RemoveCollectionFromRecipeBook(collectionId, userId uuid.UUID) error {
	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE collection_id=$1 AND user_id=$2
	`, collectionUsersTable)

	if _, err := r.db.Exec(query, collectionId, userId); err != nil {
		log.Errorf("unable to remove collection %s from user %s recipe book: %s", collectionId, userId, err)
		return fail.GrpcUnknown
	}

	return nil
}
