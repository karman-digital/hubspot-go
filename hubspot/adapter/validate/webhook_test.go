package hsvalidate

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strconv"
	"testing"
	"time"
)

func TestValidateV3RequestPreservesRawBodyAndExactQueryEncoding(t *testing.T) {
	now := time.Date(2026, 8, 29, 12, 0, 0, 0, time.UTC)
	timestamp := "1788004800000"
	request := V3Request{
		Method: "POST",
		URI:    "https://workspace.example/v1/company/123?userEmail=alex%40example.com&return=%2Fcompanies%2F123&literal=%26",
		Body:   []byte("{\n  \"target_status\": \"prospect\"\n}"), Timestamp: timestamp,
	}
	decodedURI := "https://workspace.example/v1/company/123?userEmail=alex@example.com&return=/companies/123&literal=%26"
	request.Signature = signature([]byte("secret"), request.Method+decodedURI+string(request.Body)+timestamp)
	if err := ValidateV3Request([]byte("secret"), request, func() time.Time { return now }); err != nil {
		t.Fatalf("ValidateV3Request() error = %v", err)
	}

	request.Body = []byte(`{"target_status":"prospect"}`)
	if err := ValidateV3Request([]byte("secret"), request, func() time.Time { return now }); err != ErrMismatchedSignatures {
		t.Fatalf("ValidateV3Request(compacted body) error = %v", err)
	}
}

func TestValidateV3RequestRejectsExpiredAndFutureTimestamps(t *testing.T) {
	now := time.Date(2026, 8, 29, 12, 0, 0, 0, time.UTC)
	for _, test := range []struct {
		name      string
		timestamp string
		want      error
	}{
		{"expired", "1788004499999", ErrTimestampExpired},
		{"future", "1788004800001", ErrTimestampInvalid},
		{"malformed", "nope", ErrTimestampInvalid},
	} {
		t.Run(test.name, func(t *testing.T) {
			request := V3Request{Method: "GET", URI: "https://workspace.example/v1/company/123", Timestamp: test.timestamp, Signature: "invalid"}
			if err := ValidateV3Request([]byte("secret"), request, func() time.Time { return now }); err != test.want {
				t.Fatalf("ValidateV3Request() error = %v, want %v", err, test.want)
			}
		})
	}
}

func TestValidateWebhookSignatureCompatibilityUsesRawRequestURI(t *testing.T) {
	timestamp := time.Now().UTC().UnixMilli()
	requestTimestamp := strconv.FormatInt(timestamp, 10)
	body := []byte("{ \"exact\": true }")
	source := "POSThttps://workspace.example/path?b=2&a=1" + string(body) + requestTimestamp
	if err := ValidateWebhookSignature([]byte("secret"), "workspace.example", "/path?b=2&a=1", requestTimestamp, "POST", signature([]byte("secret"), source), body); err != nil {
		t.Fatalf("ValidateWebhookSignature() error = %v", err)
	}
}

func signature(secret []byte, source string) string {
	hash := hmac.New(sha256.New, secret)
	_, _ = hash.Write([]byte(source))
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}
