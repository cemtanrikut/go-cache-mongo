package service

import (
	"go-cache-mongo/cache"
	"go-cache-mongo/controller"
	"go-cache-mongo/model"

	"go.mongodb.org/mongo-driver/mongo"
)

func Fetch(fetchModel model.FetchMongoReqData, collection *mongo.Collection) model.FetchMongoRespData {
	response := controller.Fetch(fetchModel, collection)
	return response
}

func Get(key string, collection *mongo.Collection, c *cache.Cache) []model.InMemData {
	response := controller.Get(key, collection, c)
	return response
}

func Set(key string, collection *mongo.Collection, c *cache.Cache) (model.InMemData, error) {
	respData := model.InMemData{
		Key:   key,
		Value: "Getir",
	}
	response, err := controller.Set(respData, collection, c)
	return response, err
}
