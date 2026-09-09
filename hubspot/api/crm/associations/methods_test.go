package associations

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	associationsmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm/associations"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (function roundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return function(request)
}

func TestBatchGetAllAssociationsFollowsEachSourceCursorAndPreservesZeroAndMultipleTargets(t *testing.T) {
	requests := make([]associationsmodels.BatchGetAssociationsBody, 0)
	service := associationTestService(t, func(requestNumber int, request associationsmodels.BatchGetAssociationsBody) (int, string) {
		requests = append(requests, request)
		switch requestNumber {
		case 1:
			return http.StatusOK, `{"results":[{"from":{"id":"engagement-2"},"to":[]},{"from":{"id":"engagement-1"},"to":[{"toObjectId":11}],"paging":{"next":{"after":"engagement-1-page-2"}}},{"from":{"id":"engagement-3"},"to":[{"toObjectId":31},{"toObjectId":32}]}]}`
		case 2:
			return http.StatusOK, `{"results":[{"from":{"id":"engagement-1"},"to":[{"toObjectId":12}]}]}`
		default:
			t.Fatalf("unexpected request %d", requestNumber)
			return 0, ""
		}
	})

	results, err := service.BatchGetAllAssociations("emails", "contacts", []string{"engagement-1", "engagement-2", "engagement-3"})
	if err != nil {
		t.Fatal(err)
	}
	if len(requests) != 2 {
		t.Fatalf("requests = %d, want 2", len(requests))
	}
	if len(requests[0].Inputs) != 3 || len(requests[1].Inputs) != 1 || requests[1].Inputs[0].Id != "engagement-1" || requests[1].Inputs[0].After != "engagement-1-page-2" {
		t.Fatalf("requests = %#v", requests)
	}
	if len(results) != 3 || results[0].From.ID != "engagement-1" || results[1].From.ID != "engagement-2" || results[2].From.ID != "engagement-3" {
		t.Fatalf("results = %#v", results)
	}
	if len(results[0].To) != 2 || results[0].To[0].ToObjectId != 11 || results[0].To[1].ToObjectId != 12 {
		t.Fatalf("first source targets = %#v", results[0].To)
	}
	if len(results[1].To) != 0 || len(results[2].To) != 2 {
		t.Fatalf("results = %#v", results)
	}
	if results[0].Paging.Next.After != "" {
		t.Fatalf("completed result retained paging = %#v", results[0].Paging)
	}
}

func TestBatchGetAllAssociationsRejectsInvalidInputsBeforeRequesting(t *testing.T) {
	tooMany := make([]string, 1001)
	for index := range tooMany {
		tooMany[index] = strconv.Itoa(index + 1)
	}

	for name, sourceIDs := range map[string][]string{
		"blank":      {"engagement-1", ""},
		"duplicate":  {"engagement-1", "engagement-1"},
		"over limit": tooMany,
	} {
		t.Run(name, func(t *testing.T) {
			service := associationTestService(t, func(int, associationsmodels.BatchGetAssociationsBody) (int, string) {
				t.Fatal("unexpected HubSpot request")
				return 0, ""
			})

			results, err := service.BatchGetAllAssociations("emails", "contacts", sourceIDs)
			if err == nil || results != nil {
				t.Fatalf("results = %#v, error = %v", results, err)
			}
		})
	}
}

func TestBatchGetAllAssociationsAcceptsTheAPIRequestLimit(t *testing.T) {
	sourceIDs := make([]string, 1000)
	for index := range sourceIDs {
		sourceIDs[index] = strconv.Itoa(index + 1)
	}
	service := associationTestService(t, func(requestNumber int, request associationsmodels.BatchGetAssociationsBody) (int, string) {
		if requestNumber != 1 || len(request.Inputs) != 1000 {
			t.Fatalf("request %d inputs = %d", requestNumber, len(request.Inputs))
		}
		response := associationsmodels.BatchAssociationGetResponse{Results: make([]associationsmodels.BatchAssociationResult, len(request.Inputs))}
		for index, input := range request.Inputs {
			response.Results[index] = associationsmodels.BatchAssociationResult{From: associationsmodels.From{ID: input.Id}, To: []associationsmodels.ToItem{}}
		}
		encoded, err := json.Marshal(response)
		if err != nil {
			t.Fatal(err)
		}
		return http.StatusOK, string(encoded)
	})

	results, err := service.BatchGetAllAssociations("emails", "contacts", sourceIDs)
	if err != nil || len(results) != 1000 {
		t.Fatalf("results = %d, error = %v", len(results), err)
	}
}

func TestBatchGetAllAssociationsRejectsIncompleteOrAmbiguousResponses(t *testing.T) {
	tests := map[string]string{
		"duplicate source": `{"results":[{"from":{"id":"engagement-1"},"to":[]},{"from":{"id":"engagement-1"},"to":[]}]}`,
		"unknown source":   `{"results":[{"from":{"id":"engagement-1"},"to":[]},{"from":{"id":"unknown"},"to":[]}]}`,
		"missing source":   `{"results":[{"from":{"id":"engagement-1"},"to":[]}]}`,
		"duplicate target": `{"results":[{"from":{"id":"engagement-1"},"to":[{"toObjectId":11},{"toObjectId":11}]},{"from":{"id":"engagement-2"},"to":[]}]}`,
		"reported errors":  `{"numErrors":1,"results":[{"from":{"id":"engagement-1"},"to":[]},{"from":{"id":"engagement-2"},"to":[]}]}`,
	}

	for name, response := range tests {
		t.Run(name, func(t *testing.T) {
			service := associationTestService(t, func(int, associationsmodels.BatchGetAssociationsBody) (int, string) {
				return http.StatusOK, response
			})

			results, err := service.BatchGetAllAssociations("emails", "contacts", []string{"engagement-1", "engagement-2"})
			if err == nil || results != nil {
				t.Fatalf("results = %#v, error = %v", results, err)
			}
		})
	}
}

func TestBatchGetAllAssociationsRejectsRepeatedCursorAndDuplicateTargetAcrossPages(t *testing.T) {
	for name, secondPage := range map[string]string{
		"repeated cursor":  `{"results":[{"from":{"id":"engagement-1"},"to":[{"toObjectId":12}],"paging":{"next":{"after":"page-2"}}}]}`,
		"duplicate target": `{"results":[{"from":{"id":"engagement-1"},"to":[{"toObjectId":11}]}]}`,
	} {
		t.Run(name, func(t *testing.T) {
			service := associationTestService(t, func(requestNumber int, _ associationsmodels.BatchGetAssociationsBody) (int, string) {
				if requestNumber == 1 {
					return http.StatusOK, `{"results":[{"from":{"id":"engagement-1"},"to":[{"toObjectId":11}],"paging":{"next":{"after":"page-2"}}}]}`
				}
				return http.StatusOK, secondPage
			})

			results, err := service.BatchGetAllAssociations("emails", "contacts", []string{"engagement-1"})
			if err == nil || results != nil {
				t.Fatalf("results = %#v, error = %v", results, err)
			}
		})
	}
}

func TestBatchGetAllAssociationsReturnsNoPartialResultsAfterLater207(t *testing.T) {
	service := associationTestService(t, func(requestNumber int, _ associationsmodels.BatchGetAssociationsBody) (int, string) {
		if requestNumber == 1 {
			return http.StatusOK, `{"results":[{"from":{"id":"engagement-1"},"to":[{"toObjectId":11}],"paging":{"next":{"after":"page-2"}}}]}`
		}
		return http.StatusMultiStatus, `{"numErrors":1,"results":[{"from":{"id":"engagement-1"},"to":[]}],"errors":[{"message":"failed"}]}`
	})

	results, err := service.BatchGetAllAssociations("emails", "contacts", []string{"engagement-1"})
	if err == nil || results != nil {
		t.Fatalf("results = %#v, error = %v", results, err)
	}
}

func associationTestService(t *testing.T, respond func(int, associationsmodels.BatchGetAssociationsBody) (int, string)) *AssociationService {
	t.Helper()
	requestNumber := 0
	client := retryablehttp.NewClient()
	client.Logger = nil
	client.HTTPClient.Transport = roundTripFunc(func(request *http.Request) (*http.Response, error) {
		requestNumber++
		var body associationsmodels.BatchGetAssociationsBody
		if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
			t.Fatalf("decode request %d: %v", requestNumber, err)
		}
		status, response := respond(requestNumber, body)
		return &http.Response{
			StatusCode: status,
			Status:     fmt.Sprintf("%d", status),
			Header:     make(http.Header),
			Body:       io.NopCloser(strings.NewReader(response)),
			Request:    request,
		}, nil
	})
	creds := credentials.NewHubspotOauthCredentials("", "", "", "token", "")
	creds.SetClient(client)
	return NewAssociationService(creds)
}
