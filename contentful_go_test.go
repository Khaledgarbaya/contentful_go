package contentful_go

import (
	"strings"
	"testing"
)

func TestGetEntry(t *testing.T) {
	entryID := "nyancat"
	client := createClient()
	entry, err := client.GetEntry(entryID, nil)
	if err != nil {
		t.Errorf("expected err == nil got %v", err)
	}
	if !strings.EqualFold(entry.Sys.Id, entryID) {
		t.Errorf("Expected %s got %s", entryID, entry.Sys.Id)
		t.Fail()
	}
}
func TestGetSpace(t *testing.T) {
	spaceID := "cfexampleapi"
	client := createClient()
	space, err := client.GetSpace(nil)
	if err != nil {
		t.Errorf("expected err == nil got %v", err)
	}
	if !strings.EqualFold(space.Sys.Id, spaceID) {
		t.Errorf("Expected %s got %s", spaceID, space.Sys.Id)
		t.Fail()
	}
}
func TestGetContentType(t *testing.T) {
	contentTypeID := "cat"
	client := createClient()
	contentType, err := client.GetContentType(contentTypeID, nil)
	if err != nil {
		t.Errorf("expected err == nil got %v", err)
	}
	if !strings.EqualFold(contentType.Sys.Id, contentTypeID) {
		t.Errorf("Expected %s got %s", contentTypeID, contentType.Sys.Id)
		t.Fail()
	}
}
func TestGetEntries(t *testing.T) {
	client := createClient()
	entities, err := client.GetEntries(nil)
	if err != nil {
		t.Errorf("expected err == nil got %v", err)
	}
	if len(entities.Items) == 0 {
		t.Fail()
	}
}
func TestGetContentTypes(t *testing.T) {
	client := createClient()
	contentTypes, err := client.GetContentTypes(nil)
	if err != nil {
		t.Errorf("expected err == nil got %v", err)
	}
	if len(contentTypes.Items) == 0 {
		t.Fail()
	}
}
func createClient() Contentful {
	return New("cfexampleapi", "b4c0n73n7fu1")
}
