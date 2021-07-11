package config

import (
	"context"
	"errors"
	"github.com/kataras/golog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type DB string

const (
	POST DB = "gnemes_post"
	USER DB = "gnemes_user"
)

func GetGnemesDBClient(db DB, logger *golog.Logger) (*mongo.Client, error) {
	//mongoClientOpt := options.Client().ApplyURI("mongodb://gnemes:123456@localhost:27017/gnemes?authSource=admin&authMechanism=SCRAM-SHA-256&ssl=false")
	//mongoClientOpt := options.Client().ApplyURI("mongodb://localhost:27017")
	//mongo --username "myTestDBUser" --password --authenticationDatabase test --authenticationMechanism SCRAM-SHA-256
	var mongoClientOpt *options.ClientOptions
	switch db {
	case POST:
		mongoClientOpt = options.Client().ApplyURI("mongodb://admin:password@127.0.0.1:27017/")
	case USER:
		mongoClientOpt = options.Client().ApplyURI("mongodb://admin:password@127.0.0.1:27017/")
	default:
		logger.Error("DB initial Failed,No Db was signed")
		return nil, errors.New("DB initialFailed")
	}

	logger.Info("initial db connection")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, mongoClientOpt)
	if client != nil {
		var pingErr = client.Ping(ctx, nil)
		if pingErr != nil {
			logger.Error("mongodb ping failed ", pingErr)
		} else {
			logger.Info("mongodb connect successfully")
		}
	}
	return client, err
}
