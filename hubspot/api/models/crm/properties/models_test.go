package propertiesmodels

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestEnumerationOptionsSerializesZeroDisplayOrder(t *testing.T) {
	body, err := json.Marshal(EnumerationOptions{Label: "First", Value: "first", DisplayOrder: 0})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(body), `"displayOrder":0`) {
		t.Fatalf("body = %s", body)
	}
}
