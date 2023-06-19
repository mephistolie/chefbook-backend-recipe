package postgres

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
)

func (r *Repository) ConfirmFirebaseDataLoad(messageId uuid.UUID) error {
	tx, err := r.handleMessageIdempotently(messageId)
	if err != nil {
		if isUniqueViolationError(err) {
			return nil
		} else {
			return fail.GrpcUnknown
		}
	}
	return commitTransaction(tx)
}

func (r *Repository) DeleteUserRecipes(userId uuid.UUID, deleteSharedData bool, messageId uuid.UUID) error {
	tx, err := r.handleMessageIdempotently(messageId)
	if err != nil {
		if isUniqueViolationError(err) {
			return nil
		} else {
			return fail.GrpcUnknown
		}
	}

	deleteRecipeBookQuery := fmt.Sprintf(`
			DELETE FROM %s
			WHERE user_id=$1
		`, usersTable)

	if _, err := tx.Exec(deleteRecipeBookQuery, userId); err != nil {
		log.Errorf("unable to delete user %s recipe book: %s", userId, err)
		return errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}

	if deleteSharedData {
		deleteUserRecipesQuery := fmt.Sprintf(`
			DELETE FROM %s
			WHERE owner_id=$1
		`, recipesTable)

		if _, err := tx.Exec(deleteUserRecipesQuery, userId); err != nil {
			log.Errorf("unable to delete user %s recipes: %s", userId, err)
			return errorWithTransactionRollback(tx, fail.GrpcUnknown)
		}
	} else {
		deleteNotUsedRecipesQuery := fmt.Sprintf(`
			DELETE FROM %[1]v
			WHERE owner_id=$1 AND recipe_id NOT IN
			(
				SELECT recipe_id
				FROM %[2]v
				WHERE recipe_id IN
				(
					SELECT recipe_id
					FROM %[1]v
					WHERE owner_id=$1 AND visibility<>'private'
				)
			)
		`, recipesTable, usersTable)

		if _, err := tx.Exec(deleteNotUsedRecipesQuery, userId); err != nil {
			log.Errorf("unable to delete not used user %s recipes: %s", userId, err)
			return errorWithTransactionRollback(tx, fail.GrpcUnknown)
		}
	}

	return commitTransaction(tx)
}

func (r *Repository) handleMessageIdempotently(messageId uuid.UUID) (*sql.Tx, error) {
	tx, err := r.startTransaction()
	if err != nil {
		return nil, err
	}

	addMessageQuery := fmt.Sprintf(`
			INSERT INTO %s (message_id)
			VALUES ($1)
		`, inboxTable)

	if _, err = tx.Exec(addMessageQuery, messageId); err != nil {
		if !isUniqueViolationError(err) {
			log.Error("unable to add message to inbox: ", err)
		}
		return nil, errorWithTransactionRollback(tx, err)
	}

	deleteOutdatedMessagesQuery := fmt.Sprintf(`
			DELETE FROM %[1]v
			WHERE ctid IN
			(
				SELECT ctid IN
				FROM %[1]v
				ORDER BY timestamp DESC
				OFFSET 1000
			)
		`, inboxTable)

	if _, err = tx.Exec(deleteOutdatedMessagesQuery); err != nil {
		return nil, errorWithTransactionRollback(tx, err)
	}

	return tx, nil
}
