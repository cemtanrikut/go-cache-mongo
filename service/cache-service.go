package service

import (
	"go-cache-mongo/cache"
	"go-cache-mongo/controller"
	"go-cache-mongo/model"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Get(fetchModel model.FetchMongoReqData, collection *mongo.Collection, findOptions *options.FindOptions, c *cache.Cache) {
	controller.Get(fetchModel, collection, findOptions, c)
}
