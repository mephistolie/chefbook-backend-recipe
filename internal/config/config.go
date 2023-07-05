package config

import (
	"github.com/mephistolie/chefbook-backend-common/log"
	amqpConfig "github.com/mephistolie/chefbook-backend-common/mq/config"
)

const (
	EnvDev  = "develop"
	EnvProd = "production"
)

type Config struct {
	Environment *string
	Port        *int
	LogsPath    *string

	Subscription Subscription

	ProfileService  ProfileService
	TagService      TagService
	CategoryService CategoryService

	Firebase Firebase
	Database Database
	S3       S3
	Amqp     amqpConfig.Amqp
}

type Subscription struct {
	CheckSubscription *bool

	MaxPicturesFree    *int
	MaxPicturesPremium *int

	PictureMaxSizeFree    *int64
	PictureMaxSizePremium *int64
}

type ProfileService struct {
	Addr *string
}

type TagService struct {
	Addr *string
}

type CategoryService struct {
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
	log.Infof("RECIPE SERVICE CONFIGURATION\n"+
		"Environment: %v\n"+
		"Port: %v\n"+
		"Logs path: %v\n\n"+
		"Check subscription: %v\n"+
		"Max recipe pictures for free subscription: %v\n"+
		"Max recipe pictures for premium subscription: %v\n"+
		"Max recipe picture size for free subscription: %vB\n"+
		"Max recipe picture size for premium subscription: %vB\n\n"+
		"Profile Service Address: %v\n"+
		"Tag Service Address: %v\n"+
		"Category Service Address: %v\n\n"+
		"Database host: %v\n"+
		"Database port: %v\n"+
		"Database name: %v\n\n"+
		"S3 host: %v\n"+
		"S3 bucket: %v\n"+
		"S3 region: %v\n\n"+
		"MQ host: %v\n"+
		"MQ port: %v\n"+
		"MQ vhost: %v\n\n",
		*c.Environment, *c.Port, *c.LogsPath,
		*c.Subscription.CheckSubscription, *c.Subscription.MaxPicturesFree, *c.Subscription.MaxPicturesPremium,
		*c.Subscription.PictureMaxSizeFree, *c.Subscription.PictureMaxSizePremium,
		*c.ProfileService.Addr, *c.TagService.Addr, *c.CategoryService.Addr,
		*c.Database.Host, *c.Database.Port, *c.Database.DBName,
		*c.S3.Host, *c.S3.Bucket, *c.S3.Region,
		*c.Amqp.Host, *c.Amqp.Port, *c.Amqp.VHost,
	)
}
