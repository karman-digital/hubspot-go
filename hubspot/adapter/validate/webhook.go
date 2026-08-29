package hsvalidate

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strconv"
	"strings"
	"time"
)

const maximumTimestampAge = 5 * time.Minute

type V3Request struct {
	Method    string
	URI       string
	Body      []byte
	Timestamp string
	Signature string
}

func ValidateV3Request(secret []byte, request V3Request, now func() time.Time) error {
	if now == nil {
		now = time.Now
	}
	if err := validateTimestampAt(request.Timestamp, now().UTC()); err != nil {
		return err
	}
	return validateV3Signature(secret, request)
}

func validateV3Signature(secret []byte, request V3Request) error {
	source := request.Method + decodeSignatureURI(request.URI) + string(request.Body) + request.Timestamp
	hash := hmac.New(sha256.New, secret)
	_, _ = hash.Write([]byte(source))
	expected := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	if !hmac.Equal([]byte(expected), []byte(request.Signature)) {
		return ErrMismatchedSignatures
	}
	return nil
}

func ValidateWebhookSignature(secret []byte, host, requestURI, timestamp, method, signature string, body []byte) error {
	return validateV3Signature(secret, V3Request{
		Method: method, URI: "https://" + host + requestURI, Body: body, Timestamp: timestamp, Signature: signature,
	})
}

func ValidateTimeStamp(timestamp string) error {
	return validateTimestampAt(timestamp, time.Now().UTC())
}

func validateTimestampAt(timestamp string, now time.Time) error {
	milliseconds, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return ErrTimestampInvalid
	}
	requestTime := time.UnixMilli(milliseconds)
	if requestTime.After(now) {
		return ErrTimestampInvalid
	}
	if now.Sub(requestTime) > maximumTimestampAge {
		return ErrTimestampExpired
	}
	return nil
}

func decodeSignatureURI(uri string) string {
	replacer := strings.NewReplacer(
		"%3A", ":", "%3a", ":",
		"%2F", "/", "%2f", "/",
		"%3F", "?", "%3f", "?",
		"%40", "@",
		"%21", "!",
		"%24", "$",
		"%27", "'",
		"%28", "(",
		"%29", ")",
		"%2A", "*", "%2a", "*",
		"%2C", ",", "%2c", ",",
		"%3B", ";", "%3b", ";",
	)
	return replacer.Replace(uri)
}
