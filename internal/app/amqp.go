package app

import (
	"context"
	auth "github.com/mephistolie/chefbook-backend-auth/api/mq"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/mq/config"
	mqConsumer "github.com/mephistolie/chefbook-backend-common/mq/consumer"
	mqApi "github.com/mephistolie/chefbook-backend-common/mq/dependencies"
	mqPublisher "github.com/mephistolie/chefbook-backend-common/mq/publisher"
	encryption "github.com/mephistolie/chefbook-backend-encryption/api/mq"
	api "github.com/mephistolie/chefbook-backend-recipe/api/mq"
	mqRecipeApi "github.com/mephistolie/chefbook-backend-recipe/api/mq"
	"github.com/mephistolie/chefbook-backend-recipe/internal/repository/postgres"
	amqp "github.com/wagslane/go-rabbitmq"
)

const queueProfiles = "recipe.profiles"
const queueVaults = "recipe.vaults"

var supportedMsgTypes = []string{
	auth.MsgTypeProfileFirebaseImport,
	auth.MsgTypeProfileDeleted,
	encryption.MsgTypeVaultDeleted,
}

func NewMqPublisher(
	cfg config.Amqp,
	repository *postgres.Repository,
) (*mqPublisher.Publisher, error) {
	var publisher *mqPublisher.Publisher = nil
	var err error

	if len(*cfg.Host) > 0 {
		publisher, err = mqPublisher.New(mqRecipeApi.AppId, cfg, repository)
		if err != nil {
			return nil, err
		}
		if err = publisher.Start(
			amqp.WithPublisherOptionsExchangeName(api.ExchangeRecipes),
			amqp.WithPublisherOptionsExchangeKind("fanout"),
			amqp.WithPublisherOptionsExchangeDurable,
			amqp.WithPublisherOptionsExchangeDeclare,
		); err != nil {
			return nil, err
		}

		log.Log(context.Background(), log.Event{
			Event:     "mq.publisher.initialized",
			Message:   "mq publisher initialized",
			Component: log.ComponentAMQP,
		})
	}

	return publisher, nil
}

func NewMqSubscriber(
	cfg config.Amqp,
	service mqApi.Inbox,
) (*mqConsumer.Consumer, error) {
	var consumer *mqConsumer.Consumer = nil
	var err error

	if len(*cfg.Host) > 0 {

		consumer, err = mqConsumer.New(cfg, service, supportedMsgTypes)
		if err != nil {
			return nil, err
		}
		if err = consumer.Start(
			mqConsumer.Params{
				QueueName: queueProfiles,
				Options: []func(*amqp.ConsumerOptions){
					amqp.WithConsumerOptionsQueueQuorum,
					amqp.WithConsumerOptionsQueueDurable,
					amqp.WithConsumerOptionsExchangeName(auth.ExchangeProfiles),
					amqp.WithConsumerOptionsExchangeKind("fanout"),
					amqp.WithConsumerOptionsExchangeDurable,
					amqp.WithConsumerOptionsExchangeDeclare,
					amqp.WithConsumerOptionsRoutingKey(""),
				},
			},
			mqConsumer.Params{
				QueueName: queueVaults,
				Options: []func(*amqp.ConsumerOptions){
					amqp.WithConsumerOptionsQueueQuorum,
					amqp.WithConsumerOptionsQueueDurable,
					amqp.WithConsumerOptionsExchangeName(encryption.ExchangeEncryption),
					amqp.WithConsumerOptionsExchangeKind("fanout"),
					amqp.WithConsumerOptionsExchangeDurable,
					amqp.WithConsumerOptionsExchangeDeclare,
					amqp.WithConsumerOptionsRoutingKey(""),
				},
			},
		); err != nil {
			return nil, err
		}

		log.Log(context.Background(), log.Event{
			Event:     "mq.consumer.initialized",
			Message:   "mq consumer initialized",
			Component: log.ComponentAMQP,
		})
	}

	return consumer, nil
}
