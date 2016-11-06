package contentful_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Khaledgarbaya/contentful"
)

func TestGetEntry(t *testing.T) {
	var entryId string = "nyancat"
	client := createClient()
	entry, err := client.GetEntry(entryId, nil)
	if err != nil {
		t.Error(fmt.Sprintf("expected err == nil got %v", err))
	}
	if !strings.EqualFold(entry.Sys.Id, entryId) {
		t.Error(fmt.Sprintf("Expected %s got %s", entryId, entry.Sys.Id))
		t.Fail()
	}
}
func TestGetSpace(t *testing.T) {
	var spaceId string = "cfexampleapi"
	client := createClient()
	space, err := client.GetSpace(spaceId, nil)
	if err != nil {
		t.Error(fmt.Sprintf("expected err == nil got %v", err))
	}
	if !strings.EqualFold(space.Sys.Id, spaceId) {
		t.Error(fmt.Sprintf("Expected %s got %s", spaceId, space.Sys.Id))
		t.Fail()
	}
}
func TestGetContentType(t *testing.T) {
	var contentTypeId string = "cat"
	client := createClient()
	contentType, err := client.GetContentType(contentTypeId, nil)
	if err != nil {
		t.Error(fmt.Sprintf("expected err == nil got %v", err))
	}
	if !strings.EqualFold(contentType.Sys.Id, contentTypeId) {
		t.Error(fmt.Sprintf("Expected %s got %s", contentTypeId, contentType.Sys.Id))
		t.Fail()
	}
}
func TestGetEntries(t *testing.T) {
	client := createClient()
	entities, err := client.GetEntries(nil)
	if err != nil {
		t.Error(fmt.Sprintf("expected err == nil got %v", err))
	}
	if len(entities.Items) == 0 {
		t.Fail()
	}
}
func TestGetContentTypes(t *testing.T) {
	client := createClient()
	contentTypes, err := client.GetContentTypes(nil)
	if err != nil {
		t.Error(fmt.Sprintf("expected err == nil got %v", err))
	}
	if len(contentTypes.Items) == 0 {
		t.Fail()
	}
}
func createClient() contentful.Contentful {
	return contentful.New("cfexampleapi", "b4c0n73n7fu1")
}
