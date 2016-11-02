package contentful

import "github.com/Khaledgarbaya/contentful/entities"

type Collection struct {
	Sys   entities.Sys  `json: "sys"`
	Total int           `json: "total"`
	Skip  int           `json: "skip"`
	Limit int           `json:"limit"`
	Items []interface{} `json: "items"`
}
