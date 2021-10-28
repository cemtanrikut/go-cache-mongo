package middleware

import (
	"go-cache-mongo/db"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetMongoClient() (*mongo.Collection, *options.FindOptions) {
	db.MongoClient = db.GetMongoClient("mongodb+srv://challengeUser:WUMglwNBaydH8Yvu@challenge-xzwqd.mongodb.net/getir-case-study?retryWrites=true")

	client := db.MongoClient.Database("getir-case-study")
	collection := client.Collection("records")

	findOptions := options.Find()
	findOptions.SetLimit(5)

	return collection, findOptions
}
