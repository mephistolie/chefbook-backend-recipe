package config

import (
	"context"
	"github.com/mephistolie/chefbook-backend-common/log"
	amqpConfig "github.com/mephistolie/chefbook-backend-common/mq/config"
	"time"
)

const (
	EnvDev  = "develop"
	EnvProd = "production"
)

type Config struct {
	Environment *string
	Port        *int
	LogsPath    *string

	Recipes      Recipes
	Subscription Subscription

	ProfileService    Service
	TagService        Service
	EncryptionService Service

	Firebase Firebase
	Database Database
	S3       S3
	Amqp     amqpConfig.Amqp
}

type Recipes struct {
	KeyTtl *time.Duration
}

type Subscription struct {
	CheckSubscription *bool

	MaxPicturesFree    *int
	MaxPicturesPremium *int

	PictureMaxSizeFree    *int64
	PictureMaxSizePremium *int64
}

type Service struct {
	Addr *string
}

type Firebase struct {
	Credentials *string
}

type Database struct {
	Host     *string
	Port     *int
	User     *string
	Password *string
	DBName   *string
}

type S3 struct {
	Host            *string
	AccessKeyId     *string
	SecretAccessKey *string
	Bucket          *string
	Region          *string
}

func (c Config) Validate() error {
	if *c.Environment != EnvProd {
		*c.Environment = EnvDev
	}
	return nil
}

func (c Config) Print() {
	log.Log(context.Background(), log.Event{
		Event:     "config.loaded",
		Message:   "service configuration loaded",
		Component: "config",
	})
}
