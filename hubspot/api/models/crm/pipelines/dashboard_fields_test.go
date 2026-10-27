package pipelinesmodels

import (
	"encoding/json"
	"testing"
)

func TestStageRetainsLabel(t *testing.T) {
	var stage Stage
	if err := json.Unmarshal([]byte(`{"id":"qualified","label":"Qualified"}`), &stage); err != nil {
		t.Fatal(err)
	}
	if stage.Label != "Qualified" {
		t.Fatalf("label = %q", stage.Label)
	}
}
