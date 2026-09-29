package communicationpreferences

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	communicationmodels "github.com/karman-digital/hubspot/hubspot/api/models/communicationpreferences"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

const (
	communicationChannelEmail = "EMAIL"
	statusSubscribed          = "SUBSCRIBED"
	statusUnsubscribed        = "UNSUBSCRIBED"
)

func (c *CommunicationPreferencesService) GetCommunicationPreferences() (communicationmodels.CommunicationPreferencesResponse, error) {
	var result communicationmodels.CommunicationPreferencesResponse
	resp, err := c.SendRequest(http.MethodGet, "/communication-preferences/2026-09/definitions", nil)
	if err != nil {
		return result, fmt.Errorf("get communication preference definitions: %w", err)
	}
	defer resp.Body.Close()

	rawBody, err := shared.HandleBasicResponseCode(resp)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(rawBody, &result); err != nil {
		return result, fmt.Errorf("decode communication preference definitions: %w", err)
	}
	return result, nil
}

func (c *CommunicationPreferencesService) UnsubscribeFromCommunicationPreference(contactEmail string, subscriptionID int, legalOptions ...communicationmodels.CommunicationLegalBasis) error {
	body := communicationmodels.CommunicationPreferencesPostBody{
		Channel:        communicationChannelEmail,
		StatusState:    statusUnsubscribed,
		SubscriptionID: int64(subscriptionID),
	}
	if len(legalOptions) > 0 {
		body.CommunicationLegalBasis = legalOptions[0]
	}
	_, err := c.SetCommunicationPreferenceStatus(contactEmail, body)
	return err
}

func (c *CommunicationPreferencesService) SubscribeToCommunicationPreference(contactEmail string, subscriptionID int, legalOptions ...communicationmodels.CommunicationLegalBasis) error {
	body := communicationmodels.CommunicationPreferencesPostBody{
		Channel:        communicationChannelEmail,
		StatusState:    statusSubscribed,
		SubscriptionID: int64(subscriptionID),
	}
	if len(legalOptions) > 0 {
		body.CommunicationLegalBasis = legalOptions[0]
	}
	_, err := c.SetCommunicationPreferenceStatus(contactEmail, body)
	return err
}

// SetCommunicationPreferenceStatus exposes the 2026-09 status contract directly.
// The subscribe and unsubscribe helpers remain convenience wrappers for callers
// that only need success or failure.
func (c *CommunicationPreferencesService) SetCommunicationPreferenceStatus(contactEmail string, body communicationmodels.CommunicationPreferencesPostBody) (communicationmodels.CommunicationPreferenceStatusResponse, error) {
	var result communicationmodels.CommunicationPreferenceStatusResponse
	reqBody, err := json.Marshal(body)
	if err != nil {
		return result, fmt.Errorf("marshal communication preference status: %w", err)
	}

	path := fmt.Sprintf("/communication-preferences/2026-09/statuses/%s", url.PathEscape(contactEmail))
	resp, err := c.SendRequest(http.MethodPost, path, reqBody)
	if err != nil {
		return result, fmt.Errorf("set communication preference status: %w", err)
	}
	defer resp.Body.Close()

	rawBody, err := shared.HandleBasicResponseCode(resp)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(rawBody, &result); err != nil {
		return result, fmt.Errorf("decode communication preference status: %w", err)
	}
	return result, nil
}

func (c *CommunicationPreferencesService) GetCommunicationPreferenceStatus(contactEmail string) (communicationmodels.CommunicationPreferenceStatusResponse, error) {
	var result communicationmodels.CommunicationPreferenceStatusResponse
	path := fmt.Sprintf("/communication-preferences/2026-09/statuses/%s?channel=%s", url.PathEscape(contactEmail), communicationChannelEmail)
	resp, err := c.SendRequest(http.MethodGet, path, nil)
	if err != nil {
		return result, fmt.Errorf("get communication preference status: %w", err)
	}
	defer resp.Body.Close()

	rawBody, err := shared.HandleBasicResponseCode(resp)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(rawBody, &result); err != nil {
		return result, fmt.Errorf("decode communication preference status: %w", err)
	}
	return result, nil
}
