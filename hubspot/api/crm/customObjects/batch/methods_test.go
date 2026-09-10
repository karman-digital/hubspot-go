package batchcustomobjects

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	crmmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (function roundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return function(request)
}

func TestBatchUpsertUsesUniqueProperty(t *testing.T) {
	client := retryablehttp.NewClient()
	client.Logger = nil
	client.HTTPClient.Transport = roundTripFunc(func(request *http.Request) (*http.Response, error) {
		if request.Method != http.MethodPost || request.URL.Path != "/crm/v3/objects/2-123/batch/upsert" {
			t.Fatalf("request = %s %s", request.Method, request.URL.Path)
		}
		if request.Header.Get("Authorization") != "Bearer token" || request.Header.Get("Content-Type") != "application/json" {
			t.Fatal("missing request headers")
		}
		var body struct {
			Inputs []struct {
				ID         string            `json:"id"`
				IDProperty string            `json:"idProperty"`
				Properties map[string]string `json:"properties"`
			} `json:"inputs"`
		}
		if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		if len(body.Inputs) != 1 || body.Inputs[0].ID != "contact:owner" || body.Inputs[0].IDProperty != "external_key" || body.Inputs[0].Properties["name"] != "Example" {
			t.Fatalf("body = %+v", body)
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader(`{"status":"COMPLETE","results":[{"id":"456","properties":{"external_key":"contact:owner"}}]}`)), Header: make(http.Header)}, nil
	})
	creds := credentials.NewHubspotOauthCredentials("", "", "", "token", "")
	creds.SetClient(client)
	result, err := NewBatchCustomObjectService(creds).BatchUpsert(crmmodels.BatchUpsertBody{
		Inputs: []crmmodels.BatchUpsertInput{{ID: "contact:owner", IDProperty: "external_key", Properties: map[string]any{"name": "Example"}}},
	}, "2-123")
	if err != nil {
		t.Fatal(err)
	}
	if result.Status != "COMPLETE" || len(result.Results) != 1 || result.Results[0].ID != "456" {
		t.Fatalf("result = %+v", result)
	}
}

func TestBatchUpsertExposesIncompleteResponses(t *testing.T) {
	for _, test := range []struct {
		name      string
		code      int
		body      string
		wantError bool
		status    string
		numErrors int
		errors    int
		results   int
	}{
		{name: "partial", code: http.StatusMultiStatus, body: `{"status":"COMPLETE","numErrors":1,"errors":[{"message":"invalid value"}],"results":[{"id":"456"}]}`, wantError: true, status: "COMPLETE", numErrors: 1, errors: 1, results: 1},
		{name: "embedded errors", code: http.StatusOK, body: `{"status":"COMPLETE","numErrors":1,"errors":[{"message":"invalid value"}],"results":[]}`, status: "COMPLETE", numErrors: 1, errors: 1},
		{name: "pending", code: http.StatusOK, body: `{"status":"PENDING","results":[]}`, status: "PENDING"},
		{name: "missing results", code: http.StatusOK, body: `{"status":"COMPLETE","results":[]}`, status: "COMPLETE"},
		{name: "malformed", code: http.StatusOK, body: `{`, wantError: true},
		{name: "forbidden", code: http.StatusForbidden, body: `{"message":"missing scopes"}`, wantError: true},
	} {
		t.Run(test.name, func(t *testing.T) {
			client := retryablehttp.NewClient()
			client.Logger = nil
			client.HTTPClient.Transport = roundTripFunc(func(request *http.Request) (*http.Response, error) {
				return &http.Response{StatusCode: test.code, Body: io.NopCloser(strings.NewReader(test.body)), Header: make(http.Header)}, nil
			})
			creds := credentials.NewHubspotOauthCredentials("", "", "", "token", "")
			creds.SetClient(client)
			response, err := NewBatchCustomObjectService(creds).BatchUpsert(crmmodels.BatchUpsertBody{
				Inputs: []crmmodels.BatchUpsertInput{{ID: "key", IDProperty: "external_key", Properties: map[string]any{"name": "Example"}}},
			}, "2-123")
			if (err != nil) != test.wantError {
				t.Fatalf("error = %v, want error %v", err, test.wantError)
			}
			if response.Status != test.status || response.NumErrors != test.numErrors || len(response.Errors) != test.errors || len(response.Results) != test.results {
				t.Fatalf("response = %+v", response)
			}
		})
	}
}
