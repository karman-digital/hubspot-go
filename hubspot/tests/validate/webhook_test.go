package hsvalidate_test

import (
	"fmt"
	"testing"
	"time"

	hsvalidate "github.com/karman-digital/hubspot/hubspot/adapter/validate"
	"github.com/stretchr/testify/assert"
)

type hubspotWebhookSignatureTest struct {
	name      string
	secret    []byte
	host      string
	urlPath   string
	timestamp string
	method    string
	signature string
	body      []byte
	want      error
}

var hubspotWebhookSignatureTests = []hubspotWebhookSignatureTest{
	{
		name:      "Valid signature",
		secret:    []byte("secret"),
		host:      "www.example.com",
		urlPath:   "/webhook",
		timestamp: "123456789",
		method:    "POST",
		signature: "Ty19dgGUvCxJM6t4zq3TOpxl61jOyacj9pYGrCFVoIY=",
		body:      []byte(`{"test":"test"}`),
		want:      nil,
	},
	{
		name:      "Invalid signature",
		secret:    []byte("secre"),
		host:      "www.example.com",
		urlPath:   "/webhook",
		timestamp: "123456789",
		method:    "POST",
		signature: "Ty19dgGUvCxJM6t4zq3TOpxl61jOyacj9pYGrCFVoIY=",
		body:      []byte(`{"test":"test"}`),
		want:      assert.AnError,
	},
	{
		name:      "Valid",
		secret:    []byte("secret"),
		host:      "www.example.com",
		urlPath:   "/webhook?test=test",
		timestamp: "123456789",
		method:    "GET",
		signature: "xAG+TDhMcea0lr5jWvPFe1vS+nixsuWPyYdlLC3RHC8=",
		body:      nil,
		want:      nil,
	},
}

func TestValidateHubspotPostWebhookSignature(t *testing.T) {
	for _, tt := range hubspotWebhookSignatureTests {
		t.Run(tt.name, func(t *testing.T) {
			got := hsvalidate.ValidateWebhookSignature(tt.secret, tt.host, tt.urlPath, tt.timestamp, tt.method, tt.signature, tt.body)
			if tt.want != nil {
				assert.Error(t, got)
			} else {
				assert.NoError(t, got)
			}
		})
	}
}

type validateTimeStampTest struct {
	name           string
	input          string
	expectedOutput error
}

var validateTimeStampTests = []validateTimeStampTest{
	{
		name:           "Valid Timestamp",
		input:          fmt.Sprintf("%d", time.Now().UnixMilli()),
		expectedOutput: nil,
	},
	{
		name:           "Invalid Timestamp - more than 5 minutes ago",
		input:          fmt.Sprintf("%d", time.Now().Add(-6*time.Minute).UnixMilli()),
		expectedOutput: assert.AnError,
	},
	{
		name:           "Invalid Timestamp - future timestamp",
		input:          fmt.Sprintf("%d", time.Now().Add(6*time.Minute).UnixMilli()),
		expectedOutput: assert.AnError,
	},
	{
		name:           "Invalid Timestamp - empty string",
		input:          "",
		expectedOutput: assert.AnError,
	},
}

func TestValidateTimestamp(t *testing.T) {
	for _, test := range validateTimeStampTests {
		t.Run(test.name, func(t *testing.T) {
			err := hsvalidate.ValidateTimeStamp(test.input)
			if test.expectedOutput != nil {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
