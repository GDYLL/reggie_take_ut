package entity

type ResponseData struct {
	Records []Employee `json:"records"`
	Total   int64      `json:"total"`
}
