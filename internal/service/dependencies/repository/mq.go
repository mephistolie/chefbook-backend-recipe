package repository

import "github.com/google/uuid"

type MQ interface {
	ConfirmFirebaseDataLoad(messageId uuid.UUID) error
	DeleteUserEncryptedRecipes(userId uuid.UUID, messageId uuid.UUID) error
	DeleteUserData(userId uuid.UUID, deleteSharedData bool, messageId uuid.UUID) error
}
