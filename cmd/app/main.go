package main

import (
	"flag"
	amqpConfig "github.com/mephistolie/chefbook-backend-common/mq/config"
	"github.com/mephistolie/chefbook-backend-recipe/internal/app"
	"github.com/mephistolie/chefbook-backend-recipe/internal/config"
	"github.com/peterbourgon/ff/v3"
	"os"
)

func main() {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	cfg := config.Config{
		Environment: fs.String("environment", "debug", "service environment"),
		Port:        fs.Int("port", 8080, "service port"),
		LogsPath:    fs.String("logs-path", "", "logs file path"),

		Subscription: config.Subscription{
			CheckSubscription:     fs.Bool("check-subscription", true, "enable free subscription limits"),
			MaxPicturesFree:       fs.Int("max-pictures-free", 5, "max pictures count per recipe for free subscription"),
			MaxPicturesPremium:    fs.Int("max-pictures-premium", 15, "max pictures count per recipe for premium subscription"),
			PictureMaxSizeFree:    fs.Int64("picture-max-size-free", 1024*768, "max picture size for free subscription"),
			PictureMaxSizePremium: fs.Int64("picture-max-size-premium", 1024*1536, "max picture size for premium subscription"),
		},

		ProfileService: config.ProfileService{
			Addr: fs.String("profile-addr", "", "profile service address"),
		},
		TagService: config.TagService{
			Addr: fs.String("tag-addr", "", "category service address"),
		},
		CategoryService: config.CategoryService{
			Addr: fs.String("category-addr", "", "category service address"),
		},

		Firebase: config.Firebase{
			Credentials: fs.String("firebase-credentials", "", "Firebase credentials JSON; leave empty to disable"),
		},

		Database: config.Database{
			Host:     fs.String("db-host", "localhost", "database host"),
			Port:     fs.Int("db-port", 5432, "database port"),
			User:     fs.String("db-user", "", "database user name"),
			Password: fs.String("db-password", "", "database user password"),
			DBName:   fs.String("db-name", "", "service database name"),
		},

		S3: config.S3{
			Host:            fs.String("s3-host", "", "S3 host"),
			AccessKeyId:     fs.String("s3-access-key-id", "", "S3 access key ID"),
			SecretAccessKey: fs.String("s3-secret-access-key", "", "S3 access key ID"),
			Bucket:          fs.String("s3-bucket", "images", "S3 bucket"),
			Region:          fs.String("s3-region", "us-east-1", "S3 region"),
		},

		Amqp: amqpConfig.Amqp{
			Host:     fs.String("amqp-host", "", "message broker host; leave empty to disable"),
			Port:     fs.Int("amqp-port", 5672, "message broker port"),
			User:     fs.String("amqp-user", "guest", "message broker user name"),
			Password: fs.String("amqp-password", "guest", "message broker user password"),
			VHost:    fs.String("amqp-vhost", "", "message broker virtual host"),
		},
	}
	if err := ff.Parse(fs, os.Args[1:], ff.WithEnvVars()); err != nil {
		panic(err)
	}

	err := cfg.Validate()
	if err != nil {
		panic(err)
	}

	app.Run(&cfg)
}
