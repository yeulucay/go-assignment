package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitializeMongo() *mongo.Client {

	// could've been brought from .env
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI("mongodb+srv://challengeUser:WUMglwNBaydH8Yvu@challenge-xzwqd.mongodb.net/"))

	if err != nil {
		panic(err)
	}

	return client
}
