package contentful_go

import "fmt"

//Error a custom error based on the API response
type Error struct {
	Request struct {
		URL     string              // url of the request with error
		Headers map[string][]string // the response headers
		Method  int                 // HTTP Method of the request
		Payload string              // payload returned from the server
	}
}

func (e Error) Error() string {
	return fmt.Sprintf("Contentful error")
}
