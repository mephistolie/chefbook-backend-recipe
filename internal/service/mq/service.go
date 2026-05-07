package mq

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	auth "github.com/mephistolie/chefbook-backend-auth/api/mq"
	"github.com/mephistolie/chefbook-backend-common/firebase"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/mq/model"
	encryption "github.com/mephistolie/chefbook-backend-encryption/api/mq"
	"github.com/mephistolie/chefbook-backend-recipe/internal/service/dependencies/repository"
)

type Service struct {
	mqRepo         repository.MQ
	recipeRepo     repository.Recipe
	collectionRepo repository.Collection
	firebase       *firebase.Client
}

func NewService(
	mqRepo repository.MQ,
	recipeRepo repository.Recipe,
	collectionRepo repository.Collection,
	firebase *firebase.Client,
) *Service {
	return &Service{
		mqRepo:         mqRepo,
		recipeRepo:     recipeRepo,
		collectionRepo: collectionRepo,
		firebase:       firebase,
	}
}

func (s *Service) HandleMessage(msg model.MessageData) error {
	ctx := context.Background()
	log.Log(ctx, log.Event{
		Event:     "mq.message.processing",
		Message:   "processing message",
		Component: log.ComponentAMQP,
		MessageID: msg.Id.String(),
		Payload: map[string]any{
			"message_type": msg.Type,
		},
	})
	switch msg.Type {
	case auth.MsgTypeProfileFirebaseImport:
		return s.handleFirebaseImportMsg(ctx, msg.Id, msg.Body)
	case auth.MsgTypeProfileDeleted:
		return s.handleProfileDeletedMsg(ctx, msg.Id, msg.Body)
	case encryption.MsgTypeVaultDeleted:
		return s.handleVaultDeletedMsg(ctx, msg.Id, msg.Body)
	default:
		log.LogWarn(ctx, log.Event{
			Event:     "mq.message.unsupported_type",
			Message:   "got unsupported message type",
			Component: log.ComponentAMQP,
			MessageID: msg.Id.String(),
			Payload: map[string]any{
				"message_type": msg.Type,
			},
		})
		return errors.New("not implemented")
	}
}

func (s *Service) handleFirebaseImportMsg(ctx context.Context, messageId uuid.UUID, data []byte) error {
	var body auth.MsgBodyProfileFirebaseImport
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}

	userId, err := uuid.Parse(body.UserId)
	if err != nil {
		return err
	}

	log.Log(ctx, log.Event{
		Event:     "profile.firebase_import.message.processing",
		Message:   "processing firebase profile import message",
		Component: log.ComponentAMQP,
		MessageID: messageId.String(),
		UserID:    body.UserId,
		Payload: map[string]any{
			"firebase_id": body.FirebaseId,
		},
	})
	return s.ImportFirebaseRecipes(ctx, userId, body.FirebaseId, messageId)
}

func (s *Service) handleProfileDeletedMsg(ctx context.Context, messageId uuid.UUID, data []byte) error {
	var body auth.MsgBodyProfileDeleted
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}

	userId, err := uuid.Parse(body.UserId)
	if err != nil {
		return err
	}

	log.Log(ctx, log.Event{
		Event:     "profile.deleted.message.processing",
		Message:   "processing profile deleted message",
		Component: log.ComponentAMQP,
		MessageID: messageId.String(),
		UserID:    body.UserId,
	})
	return s.mqRepo.DeleteUserData(ctx, userId, body.DeleteSharedData, messageId)
}

func (s *Service) handleVaultDeletedMsg(ctx context.Context, messageId uuid.UUID, data []byte) error {
	var body encryption.MsgBodyVaultDeleted
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}

	log.Log(ctx, log.Event{
		Event:     "vault.deleted.message.processing",
		Message:   "processing vault deleted message",
		Component: log.ComponentAMQP,
		MessageID: messageId.String(),
		UserID:    body.UserId.String(),
	})
	return s.mqRepo.DeleteUserEncryptedRecipes(ctx, body.UserId, messageId)
}
