package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"time"
)

func (r *Repository) GetCollectionKey(ctx context.Context, collectionId uuid.UUID) (uuid.UUID, time.Time, error) {
	tx, err := r.startTransaction(ctx)
	if err != nil {
		return uuid.UUID{}, time.Time{}, err
	}

	var key uuid.UUID
	var expiresAt time.Time

	createKeyQuery := fmt.Sprintf(`
			WITH s AS
			(
				SELECT key, expires_at
				FROM %[1]v
				WHERE collection_id=$1
			), i AS
			(
				INSERT INTO %[1]v (collection_id, expires_at)
				SELECT $1, $2
				WHERE NOT EXISTS (SELECT 1 FROM s)
				RETURNING key, expires_at
			)
			SELECT key, expires_at FROM i
			UNION ALL
			SELECT key, expires_at FROM s
		`, collectionKeysTable)

	row := tx.QueryRowContext(ctx, createKeyQuery, collectionId, time.Now().Add(r.keyTtl))
	if err := row.Scan(&key, &expiresAt); err != nil {
		log.Errorf("unable to create collection %s key: %s", collectionId, err)
		return uuid.UUID{}, time.Time{}, errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}
	if expiresAt.Unix() < time.Now().Unix() {
		log.Debugf("key for collection %s is outdated; updating...", collectionId.String())
		return r.updateCollectionKey(ctx, tx, collectionId)
	}

	return key, expiresAt, commitTransaction(tx)
}

func (r *Repository) updateCollectionKey(ctx context.Context, tx *sql.Tx, collectionId uuid.UUID) (uuid.UUID, time.Time, error) {
	key := uuid.New()
	expiresAt := time.Now().Add(r.keyTtl)

	updateKeyQuery := fmt.Sprintf(`
			UPDATE %s
			SET key=$1, expires_at=$2
			WHERE collection_id=$3
		`, collectionKeysTable)

	if _, err := tx.ExecContext(ctx, updateKeyQuery, key, expiresAt, collectionId); err != nil {
		log.Errorf("unable to update collection %s key: %s", collectionId, err)
		return uuid.UUID{}, time.Time{}, errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}

	return key, expiresAt, commitTransaction(tx)
}

func (r *Repository) IsCollectionKeyValid(ctx context.Context, collectionId, key uuid.UUID) (bool, error) {
	valid := false

	query := fmt.Sprintf(`
			SELECT EXISTS
			(
				SELECT 1
				FROM %s
				WHERE collection_id=$1 AND key=$2 AND expires_at>=$3
			)
		`, collectionKeysTable)

	currentTime := time.Now()
	if err := r.db.GetContext(ctx, &valid, query, collectionId, key, currentTime); err != nil {
		log.Errorf("unable to validate collection %s key: %s", collectionId, err)
		return false, fail.GrpcUnknown
	}
	return true, nil
}

func (r *Repository) AddCollectionContributor(ctx context.Context, collectionId, contributorId uuid.UUID, role string) error {
	query := fmt.Sprintf(`
		INSERT INTO %[1]v (collection_id, contributor_id, role)
		VALUES ($1, $2, $3)
	`, collectionContributorsTable)

	if _, err := r.db.ExecContext(ctx, query, collectionId, contributorId, role); err != nil {
		log.Errorf("unable to add collection %s contributor %s: %s", collectionId, contributorId, err)
		return fail.GrpcUnknown
	}

	return nil
}

func (r *Repository) RemoveCollectionContributors(ctx context.Context, collectionId uuid.UUID, contributorIds []uuid.UUID) error {
	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE collection_id=$1 AND contributor_id=ANY($2)
	`, collectionContributorsTable)

	if _, err := r.db.ExecContext(ctx, query, collectionId, contributorIds); err != nil {
		log.Errorf("unable to remove collection %s contributors: %s", collectionId, err)
		return fail.GrpcUnknown
	}

	return nil
}
