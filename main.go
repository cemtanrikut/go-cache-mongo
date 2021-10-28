package main

import (
	"fmt"
	"go-cache-mongo/db"
	"go-cache-mongo/model"
	"go-cache-mongo/service"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	db.MongoClient = db.GetMongoClient("mongodb+srv://challengeUser:WUMglwNBaydH8Yvu@challenge-xzwqd.mongodb.net/getir-case-study?retryWrites=true")

	client := db.MongoClient.Database("getir-case-study")
	collection := client.Collection("records")

	findOptions := options.Find()
	findOptions.SetLimit(5)

	reqFetchMongo := model.FetchMongoReqData{
		StartDate: "2015-01-01",
		EndDate:   "2022-02-02",
		MinCount:  2700,
		MaxCount:  3000,
	}
	fmt.Println(reqFetchMongo)

	//query:=bson.M{"eventDateTime":bson.M{"$gte": fromDate, "$lt":toDate}}
	//query:=bson.M{"field":bson.M{"$in":[]string{"value1","value2"}}}
	//bson.M{"$sum": "$counts"}
	//sum := bson.M{"$sum": "$counts.value"}
	//fmt.Println("sum ", sum)

	response := service.Get(reqFetchMongo, collection, findOptions)
	fmt.Println(response)

}
