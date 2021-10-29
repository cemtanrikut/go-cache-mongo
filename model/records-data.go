package model

import "time"

type RecordsData struct {
	Key       string    `json:"key"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"createdAt"`
	Counts    []int     `json:"counts"`
}
