package transcripts

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"errors"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	models "github.com/karman-digital/hubspot/hubspot/api/models/crm/transcripts"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

type transportFunc func(*http.Request) (*http.Response, error)

func (f transportFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }
func testService(t *testing.T, fn transportFunc) *Service {
	t.Helper()
	c := credentials.NewHubspotOauthCredentials("", "", "", "test-token", "")
	client := retryablehttp.NewClient()
	client.Logger = nil
	client.HTTPClient.Transport = fn
	c.SetClient(client)
	return NewService(c)
}
func reply(status int, body string) *http.Response {
	return &http.Response{StatusCode: status, Body: io.NopCloser(strings.NewReader(body)), Header: make(http.Header)}
}

func TestMeetingTranscriptUsesTheBetaObjectContract(t *testing.T) {
	service := testService(t, func(r *http.Request) (*http.Response, error) {
		if r.Method != "POST" || r.URL.Path != "/crm/extensions/calling/2027-03-beta/transcripts" || r.Header.Get("Authorization") != "Bearer test-token" {
			t.Fatalf("request=%s %s", r.Method, r.URL)
		}
		var body map[string]json.RawMessage
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		if string(body["objectId"]) != `"42"` || string(body["objectType"]) != `"MEETING_EVENT"` || body["transcriptUtterances"] == nil || body["engagementId"] != nil || body["transcriptCreateUtterances"] != nil {
			t.Fatalf("wrong transcript contract: %v", body)
		}
		return reply(201, `{"id":"73"}`), nil
	})
	result, err := service.Create(models.CreateRequest{ObjectID: "42", ObjectType: models.ObjectTypeMeeting, Utterances: []models.Utterance{{Text: "Hello", Speaker: models.Speaker{ID: "speaker-1", Name: "Guest"}, StartTimeMillis: 0, EndTimeMillis: 1500}}})
	if err != nil || result.ID != "73" {
		t.Fatalf("result=%+v error=%v", result, err)
	}
}

func TestTranscriptReadsAndCleanupPreserveBetaResponses(t *testing.T) {
	service := testService(t, func(r *http.Request) (*http.Response, error) {
		switch r.Method + " " + r.URL.Path {
		case "GET /crm/extensions/calling/2027-03-beta/transcripts/73":
			return reply(200, `{"id":"73","objectId":"42","objectType":"MEETING_EVENT","transcriptSource":"PUBLIC_API_PROVIDED","transcriptUtterances":[{"id":"u1","text":"Hello","speaker":{"id":"speaker-1","name":"Guest"},"startTimeMillis":0,"endTimeMillis":1500}]}`), nil
		case "GET /crm/extensions/calling/2027-03-beta/transcripts/objects/MEETING_EVENT/42":
			return reply(200, `{"status":"PROCESSING","requestedAt":"2026-10-20T14:00:00Z"}`), nil
		case "DELETE /crm/extensions/calling/2027-03-beta/transcripts/73":
			return reply(204, ""), nil
		default:
			t.Fatalf("request=%s %s", r.Method, r.URL)
			return nil, nil
		}
	})
	got, err := service.Get("73")
	if err != nil || got.ObjectID != "42" || got.ObjectType != models.ObjectTypeMeeting || got.Source != "PUBLIC_API_PROVIDED" || len(got.Utterances) != 1 {
		t.Fatalf("result=%+v error=%v", got, err)
	}
	best, err := service.GetBest(models.ObjectTypeMeeting, "42")
	if err != nil || best.Status != "PROCESSING" || best.Result != nil {
		t.Fatalf("best=%+v error=%v", best, err)
	}
	if err := service.Delete("73"); err != nil {
		t.Fatal(err)
	}
}

func TestTranscriptCreateDoesNotRetryNonIdempotentWrites(t *testing.T) {
	calls := 0
	service := testService(t, func(r *http.Request) (*http.Response, error) {
		calls++
		return reply(503, `{"message":"temporarily unavailable","category":"INTERNAL_ERROR"}`), nil
	})
	_, err := service.Create(models.CreateRequest{})
	if err == nil || calls != 1 {
		t.Fatalf("calls=%d error=%v", calls, err)
	}
}

func TestMissingTranscriptRemainsDistinguishableFromOtherFailures(t *testing.T) {
	service := testService(t, func(r *http.Request) (*http.Response, error) { return reply(404, `{"message":"not found"}`), nil })
	_, err := service.Get("73")
	if !errors.Is(err, shared.ErrResourceNotFound) {
		t.Fatalf("error=%v", err)
	}
}
