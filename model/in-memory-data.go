package model

import "time"

type InMemDataMongo struct {
	CreatedAt time.Time
	Counts    []int
	InMemData InMemData
}

type InMemData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
