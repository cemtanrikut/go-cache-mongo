package helper

import (
	"encoding/json"
	"go-cache-mongo/model"
)

func JsonMarshallGet(resp *model.FetchMongoRespData) []byte {
	response, err := json.Marshal(&resp)
	if err != nil {
		panic(err)
	}
	return response
}

func JsonMarshallSet(resp *model.InMemDataMongo) []byte {
	response, err := json.Marshal(&resp)
	if err != nil {
		panic(err)
	}
	return response
}

func JsonMarshallGetItem(resp *[]model.InMemData) []byte {
	response, err := json.Marshal(&resp)
	if err != nil {
		panic(err)
	}
	return response
}
