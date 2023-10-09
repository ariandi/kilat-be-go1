package db

import (
	"context"
	util "github.com/ariandi/kilat-be-go1/utils"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var _ *mongo.Client

// Connect database
func Connect(config util.Config) (*mongo.Client, error) {

	clientOptions := options.Client().ApplyURI(config.MongoUrl)
	if config.MongoIsPassword == "1" {
		credential := options.Credential{
			Username: config.MongoUser,
			Password: config.MongoPassword,
		}
		clientOptions.SetAuth(credential)
	}
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("⛒ Connection Failed to Database")
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("⛒ Connection Failed to Database")
		log.Fatal(err)
	}
	logrus.Info("========================")
	logrus.Info("      DB CONNECTED      ")
	logrus.Info("========================")
	return client, nil
}
