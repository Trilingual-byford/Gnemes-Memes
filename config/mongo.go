package config

import (
	"context"
	"github.com/kataras/golog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DB(logger *golog.Logger) (*mongo.Client, error) {
	mongoClientOpt := options.Client().ApplyURI("mongodb://gnemes:123456@localhost/")
	logger.Info("initial db connection")
	client, err := mongo.Connect(context.TODO(), mongoClientOpt)
	if client != nil {
		var pingErr = client.Ping(context.TODO(), nil)
		if pingErr != nil {
			logger.Error("mongodb ping failed", pingErr)
		}
	}
	return client, err
}
