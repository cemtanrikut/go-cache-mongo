package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoClient(uri string) *mongo.Client {

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.NewClient(clientOptions)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	err = client.Connect(context.Background())

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return client

}

var MongoClient *mongo.Client
