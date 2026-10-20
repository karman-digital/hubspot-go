// Package transcripts implements the native call and meeting transcript beta.
// Contract: https://docs.google.com/document/d/12RTCStyS-0YoorxrYoMn6rDljP9MyeUO9uOAeX2Sf-0/edit
package transcripts

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	models "github.com/karman-digital/hubspot/hubspot/api/models/crm/transcripts"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

const basePath = "/crm/extensions/calling/2027-03-beta/transcripts"

type Service struct{ *credentials.Credentials }

func NewService(creds *credentials.Credentials) *Service { return &Service{Credentials: creds} }

func (s *Service) Create(body models.CreateRequest) (models.CreateResponse, error) {
	var result models.CreateResponse
	raw, err := json.Marshal(body)
	if err != nil {
		return result, err
	}
	// This endpoint has no documented idempotency key. A retry following a lost
	// response or a server failure could attach a second transcript.
	creds := *s.Credentials
	client := *s.Client()
	client.RetryMax = 0
	creds.SetClient(&client)
	resp, err := creds.SendRequest(http.MethodPost, basePath, raw)
	if err != nil {
		return result, fmt.Errorf("create transcript: %w", err)
	}
	err = decodeResponse(resp, &result)
	return result, err
}
func (s *Service) Get(id string) (models.Transcript, error) {
	var result models.Transcript
	resp, err := s.SendRequest(http.MethodGet, basePath+"/"+url.PathEscape(id), nil)
	if err != nil {
		return result, fmt.Errorf("get transcript: %w", err)
	}
	err = decodeResponse(resp, &result)
	return result, err
}
func (s *Service) GetBest(objectType, objectID string) (models.BestResponse, error) {
	var result models.BestResponse
	resp, err := s.SendRequest(http.MethodGet, basePath+"/objects/"+url.PathEscape(objectType)+"/"+url.PathEscape(objectID), nil)
	if err != nil {
		return result, fmt.Errorf("get best transcript: %w", err)
	}
	err = decodeResponse(resp, &result)
	return result, err
}
func (s *Service) Delete(id string) error {
	resp, err := s.SendRequest(http.MethodDelete, basePath+"/"+url.PathEscape(id), nil)
	if err != nil {
		return fmt.Errorf("delete transcript: %w", err)
	}
	defer resp.Body.Close()
	return shared.HandleDeleteResponse(resp)
}
func decodeResponse(resp *http.Response, result any) error {
	defer resp.Body.Close()
	raw, err := shared.HandleBasicResponseCode(resp)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(raw, result); err != nil {
		return fmt.Errorf("decode transcript response: %w", err)
	}
	return nil
}
