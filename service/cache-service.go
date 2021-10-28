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

func GetItem(key string, collection *mongo.Collection) []model.InMemData {
	response := controller.GetItem(key, collection)
	return response
}

func Set(key string, collection *mongo.Collection) (model.InMemData, error) {
	respData := model.InMemData{
		Key:   key,
		Value: "Getir",
	}
	response, err := controller.Set(respData, collection)
	return response, err
}
