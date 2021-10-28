package controller

import (
	"context"
	"fmt"
	"go-cache-mongo/helper"
	"go-cache-mongo/model"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func Get(fetchModel model.FetchMongoReqData, collection *mongo.Collection, findOptions *options.FindOptions) model.FetchMongoRespData {
	sDate, _ := helper.FormatTime(fetchModel.StartDate)
	eDate, _ := helper.FormatTime(fetchModel.EndDate)

	var records []model.Record
	var response model.FetchMongoRespData

	cur, err := collection.Find(context.TODO(), bson.M{"createdAt": bson.M{"$gte": sDate, "$lt": eDate}}, findOptions)

	if err != nil {
		fmt.Println("Can't find data ", err)
		response = model.FetchMongoRespData{
			Code:    0,
			Msg:     err.Error(),
			Records: nil,
		}
	} else {
		for cur.Next(context.TODO()) {
			var elem model.Record
			err := cur.Decode(&elem)
			if err != nil {
				fmt.Println("for err")
			}

			records = append(records, elem)

		}
		cur.Close(context.TODO())

		fmt.Println(records)

		response = model.FetchMongoRespData{
			Code:    0,
			Msg:     "Success",
			Records: records,
		}

	}
	return response
}

func GetItem(key string, collection *mongo.Collection) []model.InMemData {
	var responseArray []model.InMemData

	cur, err := collection.Find(context.TODO(), bson.M{"key": key})

	if err != nil {
		fmt.Println("Can't find data ", err)
	} else {
		for cur.Next(context.TODO()) {
			var elem model.InMemData
			err := cur.Decode(&elem)
			if err != nil {
				fmt.Println("for err")
			}

			responseArray = append(responseArray, elem)

		}
		cur.Close(context.TODO())
	}

	return responseArray

}

func Set(data model.InMemData, collection *mongo.Collection) (model.InMemData, error) {
	response := data
	var req model.InMemData

	_, err := collection.InsertOne(context.TODO(), response)
	if err != nil {
		fmt.Println("Insert error ", err)
		return req, err
	}
	return response, nil
}
