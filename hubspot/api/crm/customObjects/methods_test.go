package customObjects

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	crmmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm"
	sharedmodels "github.com/karman-digital/hubspot/hubspot/api/models/shared"
	"github.com/karman-digital/hubspot/hubspot/interfaces"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (function roundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return function(request)
}

func TestCustomObjectInterfaceExposesBatchUpsert(t *testing.T) {
	client := retryablehttp.NewClient()
	client.Logger = nil
	client.HTTPClient.Transport = roundTripFunc(func(request *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: http.StatusOK, Header: make(http.Header), Body: io.NopCloser(strings.NewReader(`{"status":"COMPLETE","results":[]}`))}, nil
	})
	creds := credentials.NewHubspotOauthCredentials("", "", "", "token", "")
	creds.SetClient(client)
	var service interfaces.CustomObject = NewCustomObjectService(creds)
	if _, err := service.BatchUpsert(crmmodels.BatchUpsertBody{}, "2-123"); err != nil {
		t.Fatal(err)
	}
}

func TestGetCustomObjectsHonoursLimit(t *testing.T) {
	client := retryablehttp.NewClient()
	client.Logger = nil
	client.HTTPClient.Transport = roundTripFunc(func(request *http.Request) (*http.Response, error) {
		if got := request.URL.Query().Get("limit"); got != "100" {
			t.Fatalf("limit = %q, want 100", got)
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(strings.NewReader(`{"results":[]}`)),
		}, nil
	})
	creds := credentials.NewHubspotOauthCredentials("", "", "", "token", "")
	creds.SetClient(client)

	_, err := NewCustomObjectService(creds).GetCustomObjects("2-123", sharedmodels.GetOptions{Limit: 100})
	if err != nil {
		t.Fatalf("GetCustomObjects() error = %v", err)
	}
}
