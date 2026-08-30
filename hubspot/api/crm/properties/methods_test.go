package properties

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	propertiesmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm/properties"
	sharedmodels "github.com/karman-digital/hubspot/hubspot/api/models/shared"
)

type fakeSender struct {
	responses []*http.Response
	methods   []string
	paths     []string
	bodies    [][]byte
}

func (sender *fakeSender) SendRequest(method, path string, body []byte, _ ...sharedmodels.GetOptions) (*http.Response, error) {
	sender.methods = append(sender.methods, method)
	sender.paths = append(sender.paths, path)
	sender.bodies = append(sender.bodies, body)
	response := sender.responses[0]
	sender.responses = sender.responses[1:]
	return response, nil
}

func TestGetPropertyDecodesCompleteEnumerationDefinition(t *testing.T) {
	sender := &fakeSender{responses: []*http.Response{response(http.StatusOK, propertyJSON())}}
	property, err := newPropertiesService(sender).GetProperty("deals", "revenue_line")
	if err != nil {
		t.Fatalf("GetProperty() error = %v", err)
	}
	if property.Name != "revenue_line" || property.Label != "Revenue Line" || property.Description != "Sales revenue account" || property.GroupName != "dealinformation" || property.Type != "enumeration" || property.FieldType != "select" {
		t.Fatalf("definition = %#v", property)
	}
	if property.DisplayOrder != 7 || property.Hidden || !property.FormField || property.Archived || property.CreatedAt != "2026-01-01T00:00:00Z" || property.UpdatedAt != "2026-08-30T00:00:00Z" {
		t.Fatalf("definition metadata = %#v", property)
	}
	if len(property.Options) != 2 || property.Options[0].Value != "account-1" || property.Options[0].Label != "Revenue" || property.Options[0].Hidden || property.Options[0].DisplayOrder != 3 || property.Options[0].Description != "Current" || !property.Options[1].Hidden || property.Options[1].DisplayOrder != 8 {
		t.Fatalf("options = %#v", property.Options)
	}
}

func TestUpdatePropertyOptionsPreservesDefinitionAndOptionFidelity(t *testing.T) {
	sender := &fakeSender{responses: []*http.Response{
		response(http.StatusOK, propertyJSON()),
		response(http.StatusOK, strings.Replace(propertyJSON(), `"label":"Legacy"`, `"label":"Retired"`, 1)),
	}}
	wantOptions := []propertiesmodels.EnumerationOptions{
		{Value: "account-1", Label: "Revenue renamed", Description: "Current", Hidden: false, DisplayOrder: 3},
		{Value: "legacy", Label: "Retired", Description: "Historical", Hidden: true, DisplayOrder: 8},
	}
	updated, err := newPropertiesService(sender).UpdatePropertyOptions("deals", "revenue_line", wantOptions)
	if err != nil {
		t.Fatalf("UpdatePropertyOptions() error = %v", err)
	}
	if len(sender.methods) != 2 || sender.methods[0] != http.MethodGet || sender.methods[1] != http.MethodPatch || sender.paths[1] != "/crm/v3/properties/deals/revenue_line" {
		t.Fatalf("requests = %#v %#v", sender.methods, sender.paths)
	}
	var body propertiesmodels.PropertyUpdateBody
	if err := json.Unmarshal(sender.bodies[1], &body); err != nil {
		t.Fatalf("decode update body: %v", err)
	}
	if body.Label != "Revenue Line" || body.Description != "Sales revenue account" || body.GroupName != "dealinformation" || body.Type != "enumeration" || body.FieldType != "select" || body.DisplayOrder != 7 || body.Hidden || !body.FormField {
		t.Fatalf("update body lost definition fields: %#v", body)
	}
	if len(body.Options) != 2 || body.Options[0] != wantOptions[0] || body.Options[1] != wantOptions[1] {
		t.Fatalf("update options = %#v, want %#v", body.Options, wantOptions)
	}
	if updated.Options[1].Label != "Retired" || !updated.Options[1].Hidden || updated.Options[1].DisplayOrder != 8 {
		t.Fatalf("updated property = %#v", updated)
	}
}

func TestUpdatePropertyOptionsReturnsProviderRejection(t *testing.T) {
	sender := &fakeSender{responses: []*http.Response{
		response(http.StatusOK, propertyJSON()),
		response(http.StatusBadRequest, `{"status":"error","message":"invalid option"}`),
	}}
	if _, err := newPropertiesService(sender).UpdatePropertyOptions("deals", "revenue_line", nil); err == nil || !strings.Contains(err.Error(), "Bad Request") {
		t.Fatalf("UpdatePropertyOptions() error = %v, want provider rejection", err)
	}
}

func TestUpdatePropertyOptionsRejectsMalformedUpdatedProperty(t *testing.T) {
	sender := &fakeSender{responses: []*http.Response{
		response(http.StatusOK, propertyJSON()),
		response(http.StatusOK, `{"options":[`),
	}}
	if _, err := newPropertiesService(sender).UpdatePropertyOptions("deals", "revenue_line", nil); err == nil {
		t.Fatal("UpdatePropertyOptions() error = nil, want decode failure")
	}
}

func propertyJSON() string {
	return `{"name":"revenue_line","label":"Revenue Line","description":"Sales revenue account","groupName":"dealinformation","type":"enumeration","fieldType":"select","formField":true,"hidden":false,"displayOrder":7,"hasUniqueValue":false,"externalOptions":false,"archived":false,"createdAt":"2026-01-01T00:00:00Z","updatedAt":"2026-08-30T00:00:00Z","options":[{"label":"Revenue","value":"account-1","description":"Current","hidden":false,"displayOrder":3},{"label":"Legacy","value":"legacy","description":"Historical","hidden":true,"displayOrder":8}],"modificationMetadata":{"readOnlyOptions":false,"readOnlyValue":false,"readOnlyDefinition":false,"archivable":true}}`
}

func response(status int, body string) *http.Response {
	return &http.Response{StatusCode: status, Status: http.StatusText(status), Body: io.NopCloser(strings.NewReader(body))}
}
