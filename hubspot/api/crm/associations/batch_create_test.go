package associations

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	associationsmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm/associations"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func TestBatchCreateDefaultAssociationsPreservesPairResultsAndCompletion(t *testing.T) {
	for _, test := range []struct {
		name       string
		httpStatus int
		status     string
		errorsJSON string
		wantError  bool
	}{
		{name: "complete reordered pairs", httpStatus: 200, status: "COMPLETE", errorsJSON: `[]`},
		{name: "partial response", httpStatus: 207, status: "COMPLETE", errorsJSON: `[{"category":"VALIDATION_ERROR","context":{"fromObjectId":["r1"]}}]`, wantError: true},
		{name: "embedded errors", httpStatus: 200, status: "COMPLETE", errorsJSON: `[{"category":"VALIDATION_ERROR","context":{"fromObjectId":["r1"]}}]`},
		{name: "pending", httpStatus: 200, status: "PENDING", errorsJSON: `[]`},
	} {
		t.Run(test.name, func(t *testing.T) {
			body := associationsmodels.BatchCreateDefaultAssociationsBody{Inputs: []associationsmodels.AssociationPair{
				{From: associationsmodels.AssociationId{Id: "r1"}, To: associationsmodels.AssociationId{Id: "c1"}},
				{From: associationsmodels.AssociationId{Id: "r2"}, To: associationsmodels.AssociationId{Id: "c2"}},
			}}
			client := retryablehttp.NewClient()
			client.Logger = nil
			client.HTTPClient.Transport = roundTripFunc(func(request *http.Request) (*http.Response, error) {
				if request.Method != http.MethodPost || request.URL.Path != "/crm/v4/associations/2-123/contacts/batch/associate/default" {
					t.Fatalf("unexpected request: %s %s", request.Method, request.URL.Path)
				}
				var actual associationsmodels.BatchCreateDefaultAssociationsBody
				if err := json.NewDecoder(request.Body).Decode(&actual); err != nil || !reflect.DeepEqual(actual, body) {
					t.Fatalf("request body = %#v, error = %v", actual, err)
				}
				count := 0
				if test.errorsJSON != `[]` {
					count = 1
				}
				response := fmt.Sprintf(`{"status":%q,"numErrors":%d,"errors":%s,"results":[{"from":{"id":"r2"},"to":{"id":"c2"},"associationSpec":{"associationCategory":"USER_DEFINED","associationTypeId":175}},{"from":{"id":"r1"},"to":{"id":"c1"},"associationSpec":{"associationCategory":"USER_DEFINED","associationTypeId":175}}]}`, test.status, count, test.errorsJSON)
				return &http.Response{StatusCode: test.httpStatus, Header: make(http.Header), Body: io.NopCloser(strings.NewReader(response)), Request: request}, nil
			})
			creds := credentials.NewHubspotOauthCredentials("", "", "", "token", "")
			creds.SetClient(client)
			result, err := NewAssociationService(creds).BatchCreateDefaultAssociations("2-123", "contacts", body)
			if test.wantError != errors.Is(err, shared.ErrBatchCreate) || (!test.wantError && err != nil) {
				t.Fatalf("error = %v", err)
			}
			if result.Status != test.status || len(result.Results) != 2 {
				t.Fatalf("response = %#v", result)
			}
			if result.Results[0].From.Id != "r2" || result.Results[0].To.Id != "c2" || result.Results[1].From.Id != "r1" || result.Results[1].To.Id != "c1" || result.Results[0].AssociationSpec.AssociationTypeId != 175 || result.Results[0].AssociationSpec.AssociationCategory != "USER_DEFINED" {
				t.Fatalf("pair results = %#v", result.Results)
			}
			if test.errorsJSON != `[]` && (result.NumErrors != 1 || len(result.Errors) != 1 || result.Errors[0].Context["fromObjectId"][0] != "r1") {
				t.Fatalf("errors lost: %#v", result)
			}
		})
	}
}
