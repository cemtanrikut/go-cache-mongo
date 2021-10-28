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

func Get(fetchModel model.FetchMongoReqData, collection *mongo.Collection, findOptions *options.FindOptions) {
	sDate, _ := helper.FormatTime(fetchModel.StartDate)
	eDate, _ := helper.FormatTime(fetchModel.EndDate)

	var testArray []model.TestData

	cur, err := collection.Find(context.TODO(), bson.M{"createdAt": bson.M{"$gte": sDate, "$lt": eDate}}, findOptions)

	if err != nil {
		fmt.Println("bulunamadı", err)
	} else {
		for cur.Next(context.TODO()) {
			var elem model.TestData
			err := cur.Decode(&elem)
			if err != nil {
				fmt.Println("for err")
			}

			testArray = append(testArray, elem)

		}
		cur.Close(context.TODO())
		fmt.Println(testArray)

		//cache.Cache.Items()
	}
}
