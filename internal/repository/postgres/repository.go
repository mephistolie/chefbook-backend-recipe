package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"github.com/mephistolie/chefbook-backend-recipe/internal/config"
	"time"
)

const (
	recipesTable               = "recipes"
	translationsTable          = "translations"
	recipePicturesUploadsTable = "recipe_pictures_uploads"
	recipeBookTable            = "recipe_book"
	favouritesTable            = "favourites"
	scoresTable                = "scores"

	collectionsTable            = "collections"
	collectionContributorsTable = "collections_contributors"
	collectionKeysTable         = "collections_keys"
	collectionUsersTable        = "collections_users"
	recipesCollectionsTable     = "recipes_collections"

	inboxTable  = "inbox"
	outboxTable = "outbox"

	errUniqueViolation = "23505"
)

type Repository struct {
	db     *sqlx.DB
	keyTtl time.Duration
}

func Connect(cfg config.Database) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx",
		fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=require",
			*cfg.Host, *cfg.Port, *cfg.User, *cfg.DBName, *cfg.Password))
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewRepository(db *sqlx.DB, cfg config.Recipes) *Repository {
	return &Repository{
		db:     db,
		keyTtl: *cfg.KeyTtl,
	}
}

func (r *Repository) startTransaction() (*sql.Tx, error) {
	tx, err := r.db.Begin()
	if err != nil {
		log.Error("unable to begin transaction: ", err)
		return nil, fail.GrpcUnknown
	}
	return tx, nil
}

func errorWithTransactionRollback(tx *sql.Tx, err error) error {
	_ = tx.Rollback()
	return err
}

func commitTransaction(tx *sql.Tx) error {
	if err := tx.Commit(); err != nil {
		log.Error("unable to commit transaction: ", err)
		_ = tx.Rollback()
		return fail.GrpcUnknown
	}
	return nil
}

func isUniqueViolationError(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code == errUniqueViolation
	}
	return false
}
