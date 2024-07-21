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

var selectCollectionsQuery = fmt.Sprintf(`
	SELECT
		%[1]v.collection_id,
		%[1]v.name,
		%[1]v.visibility,
		ARRAY(
			SELECT json_agg(json_build_object('contributor_id', %[2]v.contributor_id, 'role', %[2]v.role))
			FROM %[2]v
			WHERE %[2]v.collection_id=%[2]v.collection_id
		) as contributors,
		(
			SELECT COUNT(*)
			FROM %[4]v
			WHERE %[4]v.collection_id=%[1]v.collection_id
		) AS recipes_count
	FROM %[3]v
	LEFT JOIN
		%[1]v ON %[1]v.collection_id=%[1]v.collection_id
	LEFT JOIN
		%[3]v ON %[3]v.collection_id=%[1]v.collection_id
`, collectionsTable, collectionContributorsTable, collectionUsersTable, recipesCollectionsTable)

func (r *Repository) GetCollections(userId, requesterId uuid.UUID) []entity.Collection {
	query := fmt.Sprintf(`
		%[1]v
		WHERE
			%[4]v.user_id=$1 AND (%[2]v.visibility<>'%[5]v' OR %[3]v.contributor_id=$2)
	`, selectCollectionsQuery, collectionsTable, collectionContributorsTable, collectionUsersTable, model.VisibilityPrivate)

	rows, err := r.db.Query(query, userId, requesterId)
	if err != nil {
		log.Errorf("unable to get user %s collections: %s", userId, err)
		return []entity.Collection{}
	}

	var collections []entity.Collection
	for rows.Next() {
		var collection dto.Collection
		m := pgtype.NewMap()
		err = rows.Scan(
			&collection.Id, &collection.Name,
			&collection.Visibility, m.SQLScanner(&collection.Contributors),
			&collection.RecipesCount,
		)
		if err != nil {
			log.Errorf("unable to parse user %s collection: %s", userId, err)
			continue
		}
		collections = append(collections, collection.Entity())
	}

	return collections
}

func (r *Repository) GetCollectionsMap(collectionIds []uuid.UUID) map[uuid.UUID]entity.CollectionInfo {
	query := fmt.Sprintf(`
		SELECT
			%[1]v.collection_id,
			%[1]v.name
		FROM %[1]v
		WHERE collection_id=ANY($1)
	`, collectionsTable, collectionContributorsTable, model.VisibilityPrivate)

	rows, err := r.db.Query(query, collectionIds)
	if err != nil {
		log.Errorf("unable to get collections: %s", err)
		return map[uuid.UUID]entity.CollectionInfo{}
	}

	collections := make(map[uuid.UUID]entity.CollectionInfo)
	for rows.Next() {
		var id uuid.UUID
		var collection entity.CollectionInfo
		err = rows.Scan(&id, &collection.Name)
		if err != nil {
			log.Errorf("unable to parse collection: %s", err)
			continue
		}
		collections[id] = collection
	}

	return collections
}

func (r *Repository) CreateCollection(input entity.CollectionInput) (uuid.UUID, error) {
	tx, err := r.startTransaction()
	if err != nil {
		return uuid.UUID{}, err
	}

	createCollectionQuery := fmt.Sprintf(`
		INSERT INTO %s (collection_id, visibility, name)
		VALUES ($1, $2, $3)
		RETURNING collection_id
	`, collectionsTable)

	if _, err = tx.Exec(createCollectionQuery, input.Id, input.Visibility, input.Name); err != nil {
		log.Errorf("unable to add input %s: %s", input.Id, err)
		return uuid.UUID{}, errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}

	addOwnerQuery := fmt.Sprintf(`
		INSERT INTO %[1]v (collection_id, contributor_id, role)
		VALUES ($1, $2, $3)
	`, collectionContributorsTable)

	if _, err = tx.Exec(addOwnerQuery, input.Id, input.UserId, entity.RoleOwner); err != nil {
		log.Errorf("unable to add owner for collection %s: %s", input.Id, err)
		return uuid.UUID{}, errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}

	saveCollectionForOwnerQuery := fmt.Sprintf(`
		INSERT INTO %s (collection_id, user_id)
		VALUES ($1, $2)
	`, collectionUsersTable)

	if _, err = tx.Exec(saveCollectionForOwnerQuery, input.Id, input.UserId); err != nil {
		log.Errorf("unable to save collection %s for owner %s: %s", input.Id, input.UserId, err)
		return uuid.UUID{}, errorWithTransactionRollback(tx, fail.GrpcUnknown)
	}

	return input.Id, commitTransaction(tx)
}

func (r *Repository) GetCollection(collectionId, userId uuid.UUID) (entity.Collection, error) {
	var collection dto.Collection

	query := fmt.Sprintf(`
		%[1]v
		WHERE
			%[2]v.collection_id=$1 AND (%[2]v.visibility<>'%[4]v' OR %[3]v.contributor_id=$2)
	`, selectCollectionsQuery, collectionsTable, collectionContributorsTable, model.VisibilityPrivate)

	row := r.db.QueryRow(query, collectionId, userId)
	m := pgtype.NewMap()

	if err := row.Scan(
		&collection.Id, &collection.Name,
		&collection.Visibility, m.SQLScanner(&collection.Contributors),
		&collection.RecipesCount,
	); err != nil {
		log.Errorf("unable to get collection %s: %s", collectionId, err)
		return entity.Collection{}, fail.GrpcAccessDenied
	}

	return collection.Entity(), nil
}

func (r *Repository) UpdateCollection(collection entity.CollectionInput) error {
	query := fmt.Sprintf(`
		UPDATE %[1]v
		LEFT JOIN
			%[2]v ON %[2]v.collection_id=%[1]v.collection_id
		SET %[1]v.visibility=$3, %[1]v.name=$4
		WHERE %[1]v.collection_id=$1 AND %[2]v.contributor_id=$2 AND %[2]v.role='%[3]v'
	`, collectionsTable, collectionContributorsTable, entity.RoleOwner)

	result, err := r.db.Exec(query, collection.Id, collection.UserId, collection.Visibility, collection.Name)
	if err != nil {
		log.Errorf("unable to update collection %s: %s", collection.Id, err)
		return fail.GrpcUnknown
	}
	if rows, err := result.RowsAffected(); err != nil || rows == 0 {
		log.Errorf("user %s isn't owner of collection %s: %s", collection.UserId, collection.Id, err)
		return fail.GrpcAccessDenied
	}

	return nil
}

func (r *Repository) DeleteCollection(collectionId, userId uuid.UUID) error {
	query := fmt.Sprintf(`
		DELETE FROM %[1]v
		LEFT JOIN
			%[2]v ON %[2]v.collection_id=%[1]v.collection_id
		WHERE %[1]v.collection_id=$1 AND %[2]v.contributor_id=$2 AND %[2]v.role='%[3]v'
	`, collectionsTable, collectionContributorsTable, entity.RoleOwner)

	if _, err := r.db.Exec(query, collectionId, userId); err != nil {
		log.Errorf("unable to delete collection %s: %s", collectionId, err)
		return fail.GrpcUnknown
	}

	return nil
}
