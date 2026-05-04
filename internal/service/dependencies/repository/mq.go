package repository

import (
	"context"

	"github.com/google/uuid"
)

type MQ interface {
	ConfirmFirebaseDataLoad(ctx context.Context, messageId uuid.UUID) error
	DeleteUserEncryptedRecipes(ctx context.Context, userId uuid.UUID, messageId uuid.UUID) error
	DeleteUserData(ctx context.Context, userId uuid.UUID, deleteSharedData bool, messageId uuid.UUID) error
}
