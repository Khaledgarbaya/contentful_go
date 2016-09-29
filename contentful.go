/*
Package contentful interactes with Conetenful's Delivery API https://contentful.com/developers
*/
package contentful

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const baseUrl string = "https://cdn.contentful.com/"

type Contentful struct {
	spaceId, accessToken string
}

/*
Creates contentful client
Example:
		client := contentful.CreateClient('SPACE_ID', 'DELIVERY_ACESS_TOKEN')
*/
func CreateClient(spaceId, accessToken string) Contentful {
	return Contentful{spaceId, accessToken}
}

/*
Gets an Entry
Example:
		client := contentful.CreateClient('SPACE_ID', 'ACCESS_TOKEN')
		entry, err := client.GetEntry('ENTRY_ID')
		fmt.Printf("got entry with id %s", entry.Sys.Id)
It returns a pointer to an Entry object prefilled with response data
*/
func (c Contentful) GetEntry(entryId string) (entry Entry, err error) {
	e := Entry{}

	req := c.makeRequest("GET", "entries/"+entryId)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return e, fmt.Errorf("performing request %v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &e); err != nil {
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

// utils methods =================================================================================
func (c Contentful) makeRequest(method, path string) *http.Request {
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
