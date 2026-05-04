package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"github.com/mephistolie/chefbook-backend-recipe/api/model"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
)

func (r *Repository) ConfirmFirebaseDataLoad(ctx context.Context, messageId uuid.UUID) error {
	tx, err := r.handleMessageIdempotently(ctx, messageId)
	if err != nil {
		if isUniqueViolationError(err) {
			return nil
		} else {
			return fail.GrpcUnknown
		}
	}
	return commitTransaction(tx)
}

func (r *Repository) DeleteUserEncryptedRecipes(ctx context.Context, userId uuid.UUID, messageId uuid.UUID) error {
	tx, err := r.handleMessageIdempotently(ctx, messageId)
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

	if _, err := tx.ExecContext(ctx, deleteRecipesQuery, userId); err != nil {
		log.Errorf("unable to delete user %s encrypted recipes: %s", userId, err)
		return errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}

	return commitTransaction(tx)
}

func (r *Repository) DeleteUserData(ctx context.Context, userId uuid.UUID, deleteSharedData bool, messageId uuid.UUID) error {
	tx, err := r.handleMessageIdempotently(ctx, messageId)
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
	`, recipeBookTable)

	if _, err := tx.ExecContext(ctx, deleteRecipeBookQuery, userId); err != nil {
		log.Errorf("unable to delete user %s recipe book: %s", userId, err)
		return errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}

	if deleteSharedData {
		deleteUserRecipesQuery := fmt.Sprintf(`
			DELETE FROM %s
			WHERE owner_id=$1
		`, recipesTable)

		if _, err := tx.ExecContext(ctx, deleteUserRecipesQuery, userId); err != nil {
			log.Errorf("unable to delete user %s recipes: %s", userId, err)
			return errorWithTransactionRollback(tx, fail.GrpcUnknown)
		}

		deleteUserCollectionsQuery := fmt.Sprintf(`
			DELETE FROM %[1]v
			WHERE collection_id IN
			(
				SELECT collection_id
				FROM %[2]v
				WHERE contributor_id=$1 AND role='%[3]v'
			)
		`, collectionsTable, collectionContributorsTable, entity.RoleOwner)

		if _, err := tx.ExecContext(ctx, deleteUserCollectionsQuery, userId); err != nil {
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
		`, recipesTable, recipeBookTable, model.VisibilityPrivate)

		if _, err := tx.ExecContext(ctx, deleteNotUsedRecipesQuery, userId); err != nil {
			log.Errorf("unable to delete not used user %s recipes: %s", userId, err)
			return errorWithTransactionRollback(tx, fail.GrpcUnknown)
		}

		deleteNotUsedCollectionsQuery := fmt.Sprintf(`
			DELETE FROM %[1]v
			WHERE collection_id IN
			(
				SELECT %[1]v.collection_id
				FROM %[1]v
				LEFT JOIN
					%[2]v ON %[2]v.collection_id=%[1]v.collection_id
				WHERE
					%[2]v.contributor_id=$1 AND %[2]v.role='%[4]v' AND NOT EXISTS
					(
						SELECT 1
						FROM %[3]v
						WHERE %[3]v.collection_id=%[1]v.collection_id AND %[3]v.user_id<>$1
					)
			)
			`, collectionsTable, collectionContributorsTable, collectionUsersTable, entity.RoleOwner)

		if _, err := tx.ExecContext(ctx, deleteNotUsedCollectionsQuery, userId); err != nil {
			log.Errorf("unable to delete not used user %s collections: %s", userId, err)
			return errorWithTransactionRollback(tx, fail.GrpcUnknown)
		}
	}

	return commitTransaction(tx)
}

func (r *Repository) handleMessageIdempotently(ctx context.Context, messageId uuid.UUID) (*sql.Tx, error) {
	tx, err := r.startTransaction(ctx)
	if err != nil {
		return nil, err
	}

	addMessageQuery := fmt.Sprintf(`
			INSERT INTO %s (message_id)
			VALUES ($1)
		`, inboxTable)

	if _, err = tx.ExecContext(ctx, addMessageQuery, messageId); err != nil {
		if !isUniqueViolationError(err) {
			log.Error("unable to add message to inbox: ", err)
		}
		return nil, errorWithTransactionRollback(tx, err)
	}

	deleteOutdatedMessagesQuery := fmt.Sprintf(`
			DELETE FROM %[1]v
			WHERE ctid IN
			(
				SELECT ctid
				FROM %[1]v
				ORDER BY timestamp DESC
				OFFSET 1000
			)
		`, inboxTable)

	if _, err = tx.ExecContext(ctx, deleteOutdatedMessagesQuery); err != nil {
		return nil, errorWithTransactionRollback(tx, err)
	}

	return tx, nil
}
