package contentful

import "github.com/Khaledgarbaya/contentful.go/contentful/entities"

type Entry struct {
	Sys    entities.Sys           `json:"sys"`
	Fields map[string]interface{} `json:"fields"`
}
