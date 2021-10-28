package service

import (
	"go-cache-mongo/controller"
	"go-cache-mongo/model"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Get(fetchModel model.FetchMongoReqData, collection *mongo.Collection, findOptions *options.FindOptions) model.FetchMongoRespData {
	response := controller.Get(fetchModel, collection, findOptions)
	return response
}
