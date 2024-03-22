package communicationpreferences

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (c *CommunicationPreferencesService) GetCommunicationPreferences() (hubspotmodels.CommunicationPreferencesResponse, error) {
	var communicationPreferencesResp hubspotmodels.CommunicationPreferencesResponse
	reqUrl := "https://api.hubapi.com/communication-preferences/v3/definitions"
	req, err := retryablehttp.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return communicationPreferencesResp, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken()))
	resp, err := c.Client().Do(req)
	if err != nil {
		return communicationPreferencesResp, fmt.Errorf("error making request: %s", err)
	}
	defer resp.Body.Close()
	communicationPreferencesRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return communicationPreferencesResp, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 200 {
		return communicationPreferencesResp, fmt.Errorf("error returned by endpoint: %s", communicationPreferencesRawBody)
	}
	err = json.Unmarshal(communicationPreferencesRawBody, &communicationPreferencesResp)
	if err != nil {
		return communicationPreferencesResp, fmt.Errorf("error parsing body: %s", err)
	}
	return communicationPreferencesResp, nil
}

func (c *CommunicationPreferencesService) UnsubscribeFromCommunicationPreference(contactEmail string, subscriptionId int, legalOptions ...hubspotmodels.CommunicationLegalBasis) error {
	reqUrl := "https://api.hubapi.com/communication-preferences/v3/unsubscribe"
	reqBody := hubspotmodels.CommunicationPreferencesPostBody{
		EmailAddress:   contactEmail,
		SubscriptionId: fmt.Sprintf("%d", subscriptionId),
	}
	if len(legalOptions) > 0 {
		reqBody.CommunicationLegalBasis = legalOptions[0]
	}
	reqBodyJson, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("error marshalling post body: %s", err)
	}
	req, err := retryablehttp.NewRequest("POST", reqUrl, reqBodyJson)
	if err != nil {
		return fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken()))
	resp, err := c.Client().Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %s", err)
	}
	defer resp.Body.Close()
	unsubscribeRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusBadRequest {
			var errBody hubspotmodels.ErrorResponseBody
			err = json.Unmarshal(unsubscribeRawBody, &errBody)
			if err != nil {
				return fmt.Errorf("error parsing error body: %s", err)
			}
			if errBody.Category == "VALIDATION_ERROR" {
				return shared.ErrSubscriptionAlreadySubscribed
			} else {
				return fmt.Errorf("error returned by endpoint: %s", unsubscribeRawBody)
			}
		}
		return fmt.Errorf("error returned by endpoint: %s", unsubscribeRawBody)
	}
	return nil
}

func (c *CommunicationPreferencesService) SubscribeToCommunicationPreference(contactEmail string, subscriptionId int, legalOptions ...hubspotmodels.CommunicationLegalBasis) error {
	reqUrl := "https://api.hubapi.com/communication-preferences/v3/subscribe"
	reqBody := hubspotmodels.CommunicationPreferencesPostBody{
		EmailAddress:   contactEmail,
		SubscriptionId: fmt.Sprintf("%d", subscriptionId),
	}
	if len(legalOptions) > 0 {
		reqBody.CommunicationLegalBasis = legalOptions[0]
	}
	reqBodyJson, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("error marshalling post body: %s", err)
	}
	req, err := retryablehttp.NewRequest("POST", reqUrl, reqBodyJson)
	if err != nil {
		return fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken()))
	resp, err := c.Client().Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %s", err)
	}
	defer resp.Body.Close()
	unsubscribeRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusBadRequest {
			var errBody hubspotmodels.ErrorResponseBody
			err = json.Unmarshal(unsubscribeRawBody, &errBody)
			if err != nil {
				return fmt.Errorf("error parsing error body: %s", err)
			}
			if errBody.Category == "VALIDATION_ERROR" {
				return shared.ErrSubscriptionAlreadySubscribed
			} else {
				return fmt.Errorf("error returned by endpoint: %s", unsubscribeRawBody)
			}
		}
		return fmt.Errorf("error returned by endpoint: %s", unsubscribeRawBody)
	}
	return nil
}

func (c *CommunicationPreferencesService) GetCommunicationPreferenceStatus(contactEmail string) (hubspotmodels.CommunicationPreferenceStatusResponse, error) {
	var communicationPreferenceStatus hubspotmodels.CommunicationPreferenceStatusResponse
	reqUrl := fmt.Sprintf("https://api.hubapi.com/communication-preferences/v3/status/email/%s", contactEmail)
	req, err := retryablehttp.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return communicationPreferenceStatus, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken()))
	resp, err := c.Client().Do(req)
	if err != nil {
		return communicationPreferenceStatus, fmt.Errorf("error making request: %s", err)
	}
	defer resp.Body.Close()
	communicationPreferenceStatusRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return communicationPreferenceStatus, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 200 {
		return communicationPreferenceStatus, fmt.Errorf("error returned by endpoint: %s", communicationPreferenceStatusRawBody)
	}
	err = json.Unmarshal(communicationPreferenceStatusRawBody, &communicationPreferenceStatus)
	if err != nil {
		return communicationPreferenceStatus, fmt.Errorf("error parsing body: %s", err)
	}
	return communicationPreferenceStatus, nil
}
