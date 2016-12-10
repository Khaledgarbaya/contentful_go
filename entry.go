package contentful_go

import "github.com/Khaledgarbaya/contentful_go/entities"

// Entry a complete definition of an entry
type Entry struct {
	Sys    entities.Sys           `json:"sys"`
	Fields map[string]interface{} `json:"fields"`
}
