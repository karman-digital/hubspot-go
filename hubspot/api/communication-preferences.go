package hubspot

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/hashicorp/go-retryablehttp"
	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
)

func (c *credentials) GetCommunicationPreferences() (hubspotmodels.CommunicationPreferencesResponse, error) {
	var communicationPreferencesResp hubspotmodels.CommunicationPreferencesResponse
	reqUrl := "https://api.hubapi.com/communication-preferences/v3/definitions"
	req, err := retryablehttp.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return communicationPreferencesResp, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
	resp, err := c.Client.Do(req)
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
