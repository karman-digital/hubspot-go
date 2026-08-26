package pipelines

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (function roundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return function(request)
}

func TestGetPipelinesReadsDealStageClosureMetadata(t *testing.T) {
	client := retryablehttp.NewClient()
	client.Logger = nil
	client.HTTPClient.Transport = roundTripFunc(func(request *http.Request) (*http.Response, error) {
		if got := request.URL.Path; got != "/crm/v3/pipelines/deals" {
			t.Fatalf("path = %q", got)
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(strings.NewReader(`{"results":[{"id":"default","stages":[{"id":"closedwon","metadata":{"isClosed":"true"}}]}]}`)),
		}, nil
	})
	creds := credentials.NewHubspotOauthCredentials("", "", "", "token", "")
	creds.SetClient(client)

	response, err := NewPipelineService(creds).GetPipelines("deals")
	if err != nil {
		t.Fatalf("GetPipelines() error = %v", err)
	}
	if got := response.Results[0].Stages[0].Metadata.IsClosed; got != "true" {
		t.Fatalf("isClosed = %q", got)
	}
}
