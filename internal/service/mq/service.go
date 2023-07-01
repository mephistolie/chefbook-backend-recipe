package mq

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	auth "github.com/mephistolie/chefbook-backend-auth/api/mq"
	"github.com/mephistolie/chefbook-backend-common/firebase"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/mq/model"
	"github.com/mephistolie/chefbook-backend-recipe/internal/repository/grpc"
	"github.com/mephistolie/chefbook-backend-recipe/internal/service/dependencies/repository"
)

type Service struct {
	repo     repository.Recipe
	grpc     *grpc.Repository
	firebase *firebase.Client
}

func NewService(
	repo repository.Recipe,
	grpc *grpc.Repository,
	firebase *firebase.Client,
) *Service {
	return &Service{
		repo:     repo,
		grpc:     grpc,
		firebase: firebase,
	}
}

func (s *Service) HandleMessage(msg model.MessageData) error {
	log.Infof("processing message %s with type %s", msg.Id, msg.Type)
	switch msg.Type {
	case auth.MsgTypeProfileFirebaseImport:
		return s.handleFirebaseImportMsg(msg.Id, msg.Body)
	case auth.MsgTypeProfileDeleted:
		return s.handleProfileDeletedMsg(msg.Id, msg.Body)
	default:
		log.Warnf("got unsupported message type %s for message %s", msg.Type, msg.Id)
		return errors.New("not implemented")
	}
}

func (s *Service) handleFirebaseImportMsg(messageId uuid.UUID, data []byte) error {
	var body auth.MsgBodyProfileFirebaseImport
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}

	userId, err := uuid.Parse(body.UserId)
	if err != nil {
		return err
	}

	log.Infof("import firebase profile %s for user %s...", body.FirebaseId, body.UserId)
	return s.ImportFirebaseRecipes(userId, body.FirebaseId, messageId)
}

func (s *Service) handleProfileDeletedMsg(messageId uuid.UUID, data []byte) error {
	var body auth.MsgBodyProfileDeleted
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}

	userId, err := uuid.Parse(body.UserId)
	if err != nil {
		return err
	}

	log.Infof("deleting user %s...", body.UserId)
	return s.DeleteUserRecipes(userId, body.DeleteSharedData, messageId)
}
