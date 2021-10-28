package main

import (
	"fmt"
	"go-cache-mongo/middleware"
	"go-cache-mongo/model"
	"go-cache-mongo/service"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var r *mux.Router

func main() {
	collection, findOptions := middleware.SetMongoClient()

	//query:=bson.M{"eventDateTime":bson.M{"$gte": fromDate, "$lt":toDate}}
	//query:=bson.M{"field":bson.M{"$in":[]string{"value1","value2"}}}
	//bson.M{"$sum": "$counts"}
	//sum := bson.M{"$sum": "$counts.value"}
	//fmt.Println("sum ", sum)

	reqFetchMongo := model.FetchMongoReqData{
		StartDate: "2015-01-01",
		EndDate:   "2022-02-02",
		MinCount:  2700,
		MaxCount:  3000,
	}

	responseGet := service.Get(reqFetchMongo, collection, findOptions)
	fmt.Println(responseGet)

	responseGetItem := service.GetItem("TAKwGc6Jr4i8Z487", collection)
	fmt.Println(responseGetItem)

	uuid := uuid.New().String()
	uuidWithoutHyphens := strings.Replace(uuid, "-", "", -1)
	fmt.Println("UUID ", uuidWithoutHyphens)

	responseSet, _ := service.Set(uuidWithoutHyphens, collection)
	fmt.Println(responseSet)

}
