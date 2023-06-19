package config

import (
	"github.com/mephistolie/chefbook-backend-common/log"
)

const (
	EnvDev  = "develop"
	EnvProd = "production"
)

type Config struct {
	Environment *string
	Port        *int
	LogsPath    *string

	ProfileService  ProfileService
	TagService      TagService
	CategoryService CategoryService

	Firebase Firebase
	Database Database
	Amqp     Amqp
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

type Amqp struct {
	Host     *string
	Port     *int
	User     *string
	Password *string
	VHost    *string
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
		"Profile Service Address: %v\n"+
		"Tag Service Address: %v\n"+
		"Category Service Address: %v\n\n"+
		"Database host: %v\n"+
		"Database port: %v\n"+
		"Database name: %v\n\n"+
		"MQ host: %v\n"+
		"MQ port: %v\n"+
		"MQ vhost: %v\n\n",
		*c.Environment, *c.Port, *c.LogsPath,
		*c.ProfileService.Addr, *c.TagService.Addr, *c.CategoryService.Addr,
		*c.Database.Host, *c.Database.Port, *c.Database.DBName,
		*c.Amqp.Host, *c.Amqp.Port, *c.Amqp.VHost,
	)
}
