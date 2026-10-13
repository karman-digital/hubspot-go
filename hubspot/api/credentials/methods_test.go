package credentials

import (
	"errors"
	"net/http"
	"testing"

	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func TestOAuthResponseErrorPreservesAuthenticationMeaning(t *testing.T) {
	err := oauthResponseError(http.StatusBadRequest, []byte(`{"status":"BAD_REFRESH_TOKEN","error":"invalid_request","error_description":"missing or invalid refresh token"}`))

	if !errors.Is(err, shared.ErrOAuthAuthentication) {
		t.Fatalf("oauthResponseError() = %v, want wrapped ErrOAuthAuthentication", err)
	}
}
