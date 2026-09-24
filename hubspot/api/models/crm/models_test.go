package crmmodels

import (
	"encoding/json"
	"testing"
	"time"
)

func TestResultDecodesPropertyHistory(t *testing.T) {
	var result Result
	err := json.Unmarshal([]byte(`{
		"id":"123",
		"properties":{"end_date":"2026-09-30"},
		"propertiesWithHistory":{"end_date":[
			{"value":"2026-08-31","timestamp":"2026-07-01T09:30:00Z","sourceType":"CRM_UI"},
			{"value":"2026-09-30","timestamp":"2026-08-20T10:00:00Z","sourceType":"CRM_UI"}
		]}
	}`), &result)
	if err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}
	history := result.PropertiesWithHistory["end_date"]
	if len(history) != 2 {
		t.Fatalf("history = %#v", history)
	}
	if history[0].Value != "2026-08-31" || history[0].SourceType != "CRM_UI" || !history[0].Timestamp.Equal(time.Date(2026, 7, 1, 9, 30, 0, 0, time.UTC)) {
		t.Fatalf("first history = %#v", history[0])
	}
}
