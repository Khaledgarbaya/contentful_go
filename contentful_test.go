package contentful_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Khaledgarbaya/contentful"
)

func TestGetEntry(t *testing.T) {
	var entryId string = "nyancat"
	client := contentful.CreateClient("cfexampleapi", "b4c0n73n7fu1")
	entry, err := client.GetEntry(entryId)
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
	client := contentful.CreateClient("cfexampleapi", "b4c0n73n7fu1")
	space, err := client.GetSpace(spaceId)
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
	client := contentful.CreateClient("cfexampleapi", "b4c0n73n7fu1")
	contentType, err := client.GetContentType(contentTypeId)
	if err != nil {
		t.Error(fmt.Sprintf("expected err == nil got %v", err))
	}
	if !strings.EqualFold(contentType.Sys.Id, contentTypeId) {
		t.Error(fmt.Sprintf("Expected %s got %s", contentTypeId, contentType.Sys.Id))
		t.Fail()
	}
}
