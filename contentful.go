package contentful

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const baseUrl string = "https://cdn.contentful.com/"

type Contentful struct {
	accessToken, spaceId string
}

// client methods =============================================================================
func (c Contentful) GetEntry(entryId string) (entry Entry, err error) {
	var e Entry
	req := c.makeRequest("GET", "entries/"+entryId)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return e, fmt.Errorf("performing request %v", err)
	}

	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&e); err != nil {
		return e, fmt.Errorf("decode entity %v", err)
	}
	return e, nil
}

func (c Contentful) GetEntries() (entries []Entry, err error) {
	return nil, nil
}

func (c Contentful) GetContentTypes() (contentTypes []ContentType, err error) {
	return nil, nil
}
func (c Contentful) GetContentType(contentTypeId string) (contentType ContentType, err error) {
	var ct ContentType
	return ct, nil
}

func (c Contentful) GetSpaces() (spaces []Space, err error) {
	return nil, nil
}
func (c Contentful) GetSpace(spaceId string) (space Space, err error) {
	var s Space
	return s, nil
}

// Create a contentful client
// This is the main entry point
func CreateClient(spaceId, accessToken string) Contentful {
	return Contentful{spaceId, accessToken}
}

// utils methods =================================================================================
func (c Contentful) makeRequest(method, path string) *http.Request {
	fmt.Println("Hello")
	req, err := http.NewRequest(method, baseUrl+"/spaces/"+c.spaceId+"/"+path, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil
	}

	req.Header.Set("Authorization", "Bearer "+c.accessToken)
	req.Header.Set("Content-Type", "application/vnd.contentful.delivery.v1+json")
	req.Header.Set("X-Contentful-User-Agent", "contentful.go/1.0") // hardcoded for now
	return req
}
