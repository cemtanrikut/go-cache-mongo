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

	http.ListenAndServe(":"+os.Getenv("PORT"), r)

}

//Fetch the mongodb items
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

//Get items where key includes db any item and set it to cache
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

//Set item to db and add cache
func Set(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	generatedUUID := helper.GenerateUUID()

	responseSet, err := service.Set(generatedUUID, collection, c)
	if err != nil {
		er, _ := json.Marshal(err)
		w.Write(er)
	} else {
		resp := helper.JsonMarshallSet(&responseSet)
		fmt.Println(resp, err)
		w.Write(resp)
	}
}
