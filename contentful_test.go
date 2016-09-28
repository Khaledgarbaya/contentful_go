package contentful_test

import (
	"fmt"
	"testing"

	"github.com/Khaledgarbaya/contentful"
)

func TestGetEntry(*testing.T) {
	client := contentful.CreateClient("hh70p2vwgvr9", "62f1aa605a38cc00c526ee5085404d8af80305c64227868157462725c59a3cd9")
	entry, err := client.GetEntry("4DyrC6MPp6Ws8UmQEQIGUc")
	if err != nil {
		fmt.Errorf("expected err == nil got %v", err)
	}
	fmt.Println(entry)
}
