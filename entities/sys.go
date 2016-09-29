package entities

type Sys struct {
	Id        string `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Locale    string `json:"locale"`
	Revision  int    `json:"revision"`
	Type      string `json:"type"`
}
