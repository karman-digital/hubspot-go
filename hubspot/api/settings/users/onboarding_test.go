package users

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	"io"
	"net/http"
	"strings"
	"testing"
)

type lookupTransport func(*http.Request) (*http.Response, error)

func (f lookupTransport) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }
func TestGetByEmailDistinguishesMissingFromPermissionFailure(t *testing.T) {
	for _, tc := range []struct {
		status           int
		body             string
		found, wantError bool
	}{{200, `{"id":"123","email":"rangel.metodiev@karman.digital"}`, true, false}, {404, `{}`, false, false}, {403, `{"message":"missing scope"}`, false, true}, {200, `broken`, false, true}} {
		t.Run(http.StatusText(tc.status)+tc.body, func(t *testing.T) {
			c := credentials.NewHubspotOauthCredentials("", "", "", "test", "")
			c.Client().HTTPClient.Transport = lookupTransport(func(r *http.Request) (*http.Response, error) {
				if r.Method != "GET" || r.URL.Query().Get("idProperty") != "EMAIL" || !strings.HasSuffix(r.URL.Path, "/rangel.metodiev@karman.digital") {
					t.Fatalf("request=%s", r.URL)
				}
				return &http.Response{StatusCode: tc.status, Body: io.NopCloser(strings.NewReader(tc.body)), Header: http.Header{}}, nil
			})
			_, found, err := NewUserService(c).GetByEmail("rangel.metodiev@karman.digital")
			if found != tc.found || (err != nil) != tc.wantError {
				t.Fatalf("found=%t err=%v", found, err)
			}
		})
	}
}
