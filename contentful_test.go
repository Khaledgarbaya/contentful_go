package contentful_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Khaledgarbaya/contentful"
)

func TestGetEntry(t *testing.T) {
	var entryId string = "4DyrC6MPp6Ws8UmQEQIGUc"
	client := contentful.CreateClient("hh70p2vwgvr9", "62f1aa605a38cc00c526ee5085404d8af80305c64227868157462725c59a3cd9")
	entry, err := client.GetEntry(entryId)
	if err != nil {
		t.Error(fmt.Sprintf("expected err == nil got %v", err))
	}
	if !strings.EqualFold(entry.Id, entryId) {
		t.Error(fmt.Sprintf("Expected %s got %s", entryId, entry.Id))
		t.Fail()
	}
}
