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

	var records []model.Record

	cur, err := collection.Find(context.TODO(), bson.M{"createdAt": bson.M{"$gte": sDate, "$lt": eDate}}, findOptions)

	if err != nil {
		fmt.Println("Can't find data ", err)
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

	}
}
