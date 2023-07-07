package model

import "time"

type Record struct {
	Key        string    `json:"key"`
	CreatedAt  time.Time `json:"createdAt"`
	TotalCount int       `json:"totalCount"`
}

type RecordsRequestPayload struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	MinCount  int    `json:"minCount"`
	MaxCount  int    `json:"maxCount"`
}

type RecordsResponsePayload struct {
	Code    int      `json:"code"`
	Message string   `json:"msg"`
	Records []Record `json:"records"`
}
