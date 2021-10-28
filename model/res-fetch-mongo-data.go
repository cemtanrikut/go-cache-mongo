package model

import "time"

type FetchMongoRespData struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Records []int  `json:"records"`
}

type Record struct {
	Key        string    `json:"key"`
	CreatedAt  time.Time `json:"createdAt"`
	TotalCount int       `json:"totalCount"`
}
