/*
Package contentful interactes with Conetenful's Delivery API https://contentful.com/developers
*/
package contentful

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const baseUrl string = "https://cdn.contentful.com/"

type Contentful struct {
	spaceId, accessToken string
}

// Creates contentful client
// Example:
//  		client := contentful.CreateClient('SPACE_ID', 'DELIVERY_ACESS_TOKEN')
//
func CreateClient(spaceId, accessToken string) Contentful {
	return Contentful{spaceId, accessToken}
}

// Gets an Entry
// Example:
//  		client := contentful.CreateClient("SPACE_ID", "ACCESS_TOKEN")
//  		entry, _ := client.GetEntry("ENTRY_ID")
//  		fmt.Printf("got entry with id %s", entry.Sys.Id)
// It returns a pointer to an Entry object prefilled with response data
func (c Contentful) GetEntry(entryId string, query map[string]string) (entry Entry, err error) {
	e := Entry{}
	reader, err := c.performeRequest("GET", "spaces/"+c.spaceId+"/entries/"+entryId, query)
	if err != nil {
		return e, err
	}
	body, err := ioutil.ReadAll(reader)
	if err := json.Unmarshal(body, &e); err != nil {
		return e, fmt.Errorf("error while decoding the json response %v", err)
	}
	return e, nil
}

// Gets a collection of entries
// Example:
//  		client := contentful.CreateClient("SPACE_ID", "ACCESS_TOKEN")
//  		entries, _ := client.GetEntries()
//  		fmt.Printf("got entry with id %s", entries.items[0].Sys.Id)
func (c Contentful) GetEntries(query map[string]string) (entriesCollection Collection, err error) {
	ec := Collection{}
	reader, err := c.performeRequest("GET", "spaces/"+c.spaceId+"/entries/", query)
	if err != nil {
		return ec, err
	}
	body, err := ioutil.ReadAll(reader)
	if err := json.Unmarshal(body, &ec); err != nil {
		return ec, fmt.Errorf("error while decoding the json response %v", err)
	}
	return ec, nil
}

// Gets a collection of contentTypes
// Example:
//  		client := contentful.CreateClient("SPACE_ID", "ACCESS_TOKEN")
//  		contentTypes, _ := client.GetContentTypes()
//  		fmt.Printf("got contentType with id %s", contentTypes.items[0].Sys.Id)
func (c Contentful) GetContentTypes(query map[string]string) (contentTypesCollection Collection, err error) {
	ctc := Collection{}
	reader, err := c.performeRequest("GET", "spaces/"+c.spaceId+"/content_types/", query)
	if err != nil {
		return ctc, err
	}
	body, err := ioutil.ReadAll(reader)
	if err := json.Unmarshal(body, &ctc); err != nil {
		return ctc, fmt.Errorf("error while decoding the json response %v", err)
	}
	return ctc, nil
}

// Gets a ContentType
// Example:
//  		client := contentful.CreateClient("SPACE_ID", "ACCESS_TOKEN")
//  		contentType, _ := client.GetContentType("contentTypeId")
//  		fmt.Printf("got entry with id %s", contentType.Sys.Id)
func (c Contentful) GetContentType(contentTypeId string, query map[string]string) (contentType ContentType, err error) {
	ct := ContentType{}
	reader, err := c.performeRequest("GET", "spaces/"+c.spaceId+"/content_types/"+contentTypeId, query)
	if err != nil {
		return ct, err
	}
	body, err := ioutil.ReadAll(reader)
	if err := json.Unmarshal(body, &ct); err != nil {
		return ct, fmt.Errorf("error while decoding the json response %v", err)
	}
	return ct, nil
}

// Gets a Space
// Example:
//  		client := contentful.CreateClient("SPACE_ID", "ACCESS_TOKEN")
//  		space, _ := client.GetSpace("spaceId")
//  		fmt.Printf("got space with id %s", spcae.Sys.Id)
func (c Contentful) GetSpace(spaceId string, query map[string]string) (space Space, err error) {
	s := Space{}
	reader, err := c.performeRequest("GET", "spaces/"+spaceId, query)
	if err != nil {
		return s, err
	}
	body, err := ioutil.ReadAll(reader)
	if err := json.Unmarshal(body, &s); err != nil {
		return s, fmt.Errorf("error while decoding the json response %v", err)
	}
	return s, nil
}

// utils methods =================================================================================
func (c Contentful) makeRequest(method, path string) (request *http.Request, err error) {
	req, err := http.NewRequest(method, baseUrl+"/"+path, nil)
	if err != nil {
		return nil, fmt.Errorf("error making request: ", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.accessToken)
	req.Header.Set("Content-Type", "application/vnd.contentful.delivery.v1+json")
	req.Header.Set("X-Contentful-User-Agent", "contentful.go/1.0") // hardcoded for now
	return req, nil
}
func (c Contentful) performeRequest(method, path string, query map[string]string) (reader io.Reader, err error) {
	req, err := c.makeRequest(method, path)
	if query != nil && len(query) > 0 {
		q := req.URL.Query()
		for k, v := range query {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("performing request %v", err)
	}

	return resp.Body, nil
}
