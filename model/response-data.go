package model

type ResponseData struct {
	Code    int               `json:"code"`
	Msg     string            `json:"msg"`
	Records map[Record]Record `json:"records"`
}

type Record struct {
	Key        string `json:"key"`
	CreatedAt  string `json:"createdAt"`
	TotalCount int    `json:"totalCount"`
}
