package controller

import (
	"context"
	"fmt"
	"go-cache-mongo/cache"
	"go-cache-mongo/helper"
	"go-cache-mongo/model"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

func Fetch(fetchModel model.FetchMongoReqData, collection *mongo.Collection) model.FetchMongoRespData {
	sDate, _ := helper.FormatTime(fetchModel.StartDate)
	eDate, _ := helper.FormatTime(fetchModel.EndDate)

	var records []model.Record
	var response model.FetchMongoRespData

	var recDatas []model.RecordsData

	//cur, err := collection.Find(context.TODO(), bson.M{"createdAt": bson.M{"$gte": sDate, "$lt": eDate}})

	m := bson.M{
		"createdAt": bson.M{
			"$gte": sDate,
			"$lt":  eDate,
		},
		//"key":        "key",
		//"totalCount": "totalCount", //bson.M{"$sum": "count"},
	}

	//m'den gelen key ile db'ye sorgu at (Yeni model oluşturulması gerek). m'deki key ile yeni sorgudaki key eşleştir. counts'u topla, m deki totalCount'a yaz

	//cur, err := collection.Aggregate(context.TODO(), mongo.Pipeline{m})
	cur, err := collection.Find(context.TODO(), m)

	if err != nil {
		fmt.Println("Can't find data ", err)
		response = model.FetchMongoRespData{
			Code:    0,
			Msg:     err.Error(),
			Records: nil,
		}
	} else {
		for cur.Next(context.TODO()) {
			var elem model.RecordsData
			err := cur.Decode(&elem)
			if err != nil {
				fmt.Println("for err")
			}

			recDatas = append(recDatas, elem)

		}
		cur.Close(context.TODO())

		fmt.Println(recDatas)

		var rec model.Record
		var sumCount int = 0
		for _, item := range recDatas {
			for _, cnt := range item.Counts {
				sumCount += cnt
			}
			if sumCount > fetchModel.MinCount && sumCount < fetchModel.MaxCount {
				rec = model.Record{
					Key:        item.Key,
					CreatedAt:  item.CreatedAt,
					TotalCount: sumCount,
				}

				records = append(records, rec)
			}
		}

		response = model.FetchMongoRespData{
			Code:    0,
			Msg:     "Success",
			Records: records,
		}

	}
	return response
}

func Get(key string, collection *mongo.Collection, c *cache.Cache) []model.InMemData {
	var responseArray []model.InMemData

	cur, err := collection.Find(context.TODO(), bson.M{"key": key})
	fmt.Println(cur)
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
			//c.Set(elem, elem.Value)
		}
	}
	cur.Close(context.TODO())

	return responseArray

}

func Set(data model.InMemData, collection *mongo.Collection, c *cache.Cache) (model.InMemData, error) {
	response := data
	var req model.InMemData

	_, err := collection.InsertOne(context.TODO(), response)
	if err != nil {
		fmt.Println("Insert error ", err)
		return req, err
	}
	c.Set(data, data.Value)

	return response, nil
}
