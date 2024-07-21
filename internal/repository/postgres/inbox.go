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

func (r *Repository) DeleteUserEncryptedRecipes(userId uuid.UUID, messageId uuid.UUID) error {
	tx, err := r.handleMessageIdempotently(messageId)
	if err != nil {
		if isUniqueViolationError(err) {
			return nil
		} else {
			return fail.GrpcUnknown
		}
	}

	deleteRecipesQuery := fmt.Sprintf(`
		DELETE FROM %s
		WHERE owner_id=$1 AND encrypted=true
	`, recipesTable)

	if _, err := tx.Exec(deleteRecipesQuery, userId); err != nil {
		log.Errorf("unable to delete user %s encrypted recipes: %s", userId, err)
		return errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}

	return commitTransaction(tx)
}

func (r *Repository) DeleteUserData(userId uuid.UUID, deleteSharedData bool, messageId uuid.UUID) error {
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
	`, recipeUsersTable)

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

		deleteUserCollectionsQuery := fmt.Sprintf(`
			DELETE FROM %[1]v
			RIGHT JOIN
				%[2]v ON %[2]v.recipe_id=%[1]v.recipe_id
			WHERE %[2]v.user_id=$1 AND %[2]v.role='%[3]v'
		`, collectionsTable, collectionContributorsTable, entity.RoleOwner)

		if _, err := tx.Exec(deleteUserCollectionsQuery, userId); err != nil {
			log.Errorf("unable to delete user %s collections: %s", userId, err)
			return errorWithTransactionRollback(tx, fail.GrpcUnknown)
		}
	} else {
		deleteNotUsedRecipesQuery := fmt.Sprintf(`
			DELETE FROM %[1]v
			WHERE owner_id=$1 AND recipe_id NOT IN
			(
				SELECT recipe_id
				FROM %[1]v
				LEFT JOIN
					%[2]v ON %[2]v.recipe_id=%[1]v.recipe_id
				WHERE
					%[1]v.owner_id=$1 AND %[1]v.visibility<>'%[3]v' AND user_id<>$1
			)
		`, recipesTable, recipeUsersTable, model.VisibilityPrivate)

		if _, err := tx.Exec(deleteNotUsedRecipesQuery, userId); err != nil {
			log.Errorf("unable to delete not used user %s recipes: %s", userId, err)
			return errorWithTransactionRollback(tx, fail.GrpcUnknown)
		}

		deleteNotUsedCollectionsQuery := fmt.Sprintf(`
			DELETE FROM %[1]v
			WHERE collection_id
			(
				SELECT %[1]v.collection_id
				FROM %[1]v
				LEFT JOIN
					%[2]v ON %[2]v.collection_id=%[1]v.collection_id
				LEFT JOIN
					%[3]v ON %[3]v.collection_id=%[1]v.collection_id
				HAVING COUNT(%[2]v.user_id) > 0
				WHERE
					%[2]v.contributor_id=$1 AND %[2]v.role='%[4]v' AND %[1]v.visibility<>'%[4]v'
			)
		`, collectionsTable, collectionContributorsTable, collectionUsersTable, entity.RoleOwner, model.VisibilityPrivate)

		if _, err := tx.Exec(deleteNotUsedCollectionsQuery, userId); err != nil {
			log.Errorf("unable to delete not used user %s collections: %s", userId, err)
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
