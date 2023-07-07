package model

type MemoryRequestPayload struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type MemoryErrorResponsePayload struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}
