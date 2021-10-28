package main

import (
	"fmt"
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

func main() {
	//Mongo setup
	collection, findOptions = middleware.SetMongoClient()

	r = mux.NewRouter()
	r.HandleFunc("/get", Get).Methods(http.MethodGet)
	r.HandleFunc("/set", Set).Methods(http.MethodPost)
	r.HandleFunc("/get/{key}", GetItem).Methods(http.MethodGet)

	fmt.Println(os.Getenv("PORT"))
	http.ListenAndServe(os.Getenv("PORT"), r)

	//query:=bson.M{"eventDateTime":bson.M{"$gte": fromDate, "$lt":toDate}}
	//query:=bson.M{"field":bson.M{"$in":[]string{"value1","value2"}}}
	//bson.M{"$sum": "$counts"}
	//sum := bson.M{"$sum": "$counts.value"}
	//fmt.Println("sum ", sum)

	responseGetItem := service.GetItem("TAKwGc6Jr4i8Z487", collection)
	fmt.Println(responseGetItem)

}

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	reqFetchMongo := model.FetchMongoReqData{
		StartDate: "2015-01-01",
		EndDate:   "2022-02-02",
		MinCount:  2700,
		MaxCount:  3000,
	}
	responseGet := service.Get(reqFetchMongo, collection, findOptions)
	resp := helper.JsonMarshallGet(&responseGet)
	fmt.Println(resp)
	w.Write(resp)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	param := mux.Vars(r)
	key := param["key"]

	responseGetItem := service.GetItem(key, collection)
	resp := helper.JsonMarshallGetItem(&responseGetItem)
	fmt.Println(responseGetItem)
	w.Write(resp)

}

func Set(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	generatedUUID := helper.GenerateUUID()

	responseSet, _ := service.Set(generatedUUID, collection)
	resp := helper.JsonMarshallSet(&responseSet)
	fmt.Println(resp)
	w.Write(resp)
}
