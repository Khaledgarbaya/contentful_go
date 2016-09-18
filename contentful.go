package contentful
import(
//	"fmt"
	"log"
	"net/http"
//	"net/url"
)
// Constants
// const contentfulClient http.Client = http.Client{}
const baseUrl string = "https://cdn.contentful.com/"

type Contentful struct{
	accessToken, spaceId string
}

// client methods =============================================================================
func (c Contentful) GetEntry(entryId string) string{
	req := c.makeRequest("GET", "entries/"+entryId)

 // For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()
  
	return resp.Body
}

// main entry point =============================================================================
func CreateClient(spaceId, accessToken string) Contentful {
  return Contentful{spaceId, accessToken} 
}

// utils methods =================================================================================
func (c Contentful)makeRequest(method, path string) *http.Request{
	req,err := http.NewRequest(method, baseUrl+path, nil)
	
	if err != nil {
    log.Fatal("NewRequest: ", err)
		return nil
	}

	req.Header.Set("Authorization", "Bearer "+c.accessToken)
	req.Header.Set("Content-Type", "application/vnd.contentful.delivery.v1+json")
	req.Header.Set("X-Contentful-User-Agent", "contentful.go/1.0") // hardcoded for now
	return req
}
