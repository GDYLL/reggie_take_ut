package entity

type ResponseData struct {
	Records interface{} `json:"records"`
	Total   int64       `json:"total"`
}
