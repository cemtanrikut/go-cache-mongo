package main

import (
	"encoding/json"
	"fmt"
	"go-cache-mongo/cache"
	"go-cache-mongo/helper"
	"go-cache-mongo/middleware"
	"go-cache-mongo/model"
	"go-cache-mongo/service"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var r *mux.Router
var collection *mongo.Collection
var findOptions *options.FindOptions
var c *cache.Cache

func main() {
	//Mongo setup
	collection, findOptions = middleware.SetMongoClient()

	r = mux.NewRouter()
	r.HandleFunc("/mongodb/fetch", Fetch).Methods(http.MethodPost)
	r.HandleFunc("/set", Set).Methods(http.MethodPost)
	r.HandleFunc("/get/{key}", Get).Methods(http.MethodGet)

	//fmt.Println(os.Getenv("PORT"))
	http.ListenAndServe(os.Getenv("PORT"), r)

	//query:=bson.M{"eventDateTime":bson.M{"$gte": fromDate, "$lt":toDate}}
	//query:=bson.M{"field":bson.M{"$in":[]string{"value1","value2"}}}
	//bson.M{"$sum": "$counts"}
	//sum := bson.M{"$sum": "$counts.value"}
	//fmt.Println("sum ", sum)

	//responseGetItem := service.GetItem("TAKwGc6Jr4i8Z487", collection, c)
	//fmt.Println(responseGetItem)

}

func Fetch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var mData model.FetchMongoReqData
	err := json.NewDecoder(r.Body).Decode(&mData)
	if err != nil {
		fmt.Println("parsing err", err)
	}

	reqFetchMongo := model.FetchMongoReqData{
		StartDate: mData.StartDate,
		EndDate:   mData.EndDate,
		MinCount:  mData.MinCount,
		MaxCount:  mData.MaxCount,
	}
	responseGet := service.Fetch(reqFetchMongo, collection)
	resp := helper.JsonMarshallGet(&responseGet)
	fmt.Println(resp)
	w.Write(resp)
}

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	param := mux.Vars(r)
	key := param["key"]

	fmt.Println("key ", key)

	responseGetItem := service.Get(key, collection, c)
	resp := helper.JsonMarshallGetItem(&responseGetItem)
	fmt.Println(responseGetItem)
	fmt.Println(resp)
	w.Write(resp)

}

func Set(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	generatedUUID := helper.GenerateUUID()

	responseSet, _ := service.Set(generatedUUID, collection, c)
	resp := helper.JsonMarshallSet(&responseSet)
	fmt.Println(resp)
	w.Write(resp)
}
