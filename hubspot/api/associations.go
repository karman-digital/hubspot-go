package hubspot

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/hashicorp/go-retryablehttp"
	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
)

func (c *credentials) CreateDefaultAssociation(fromObject, toObject string, fromId, toId int) (hubspotmodels.BatchResponse, error) {
	var associationResp hubspotmodels.BatchResponse
	reqUrl := fmt.Sprintf("https://api.hubapi.com/crm/v4/objects/%s/%d/associations/default/%s/%d", fromObject, fromId, toObject, toId)
	req, err := retryablehttp.NewRequest("PUT", reqUrl, nil)
	if err != nil {
		return associationResp, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
	resp, err := c.Client.Do(req)
	if err != nil {
		return associationResp, fmt.Errorf("error making request: %s", err)
	}
	defer resp.Body.Close()
	associationRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return associationResp, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 200 {
		return associationResp, fmt.Errorf("error returned by endpoint: %s", associationRawBody)
	}
	err = json.Unmarshal(associationRawBody, &associationResp)
	if err != nil {
		return associationResp, fmt.Errorf("error parsing body: %s", err)
	}

	return associationResp, nil
}

func (c *credentials) BatchCreateDefaultAssociations(fromObject, toObject string, associations hubspotmodels.BatchCreateDefaultAssociationsBody) (hubspotmodels.BatchResponse, error) {
	var associationResp hubspotmodels.BatchResponse
	reqUrl := fmt.Sprintf("https://api.hubapi.com/crm/v4/associations/%s/%s/batch/associate/default", fromObject, toObject)
	reqBody, err := json.Marshal(associations)
	if err != nil {
		return associationResp, fmt.Errorf("error marshalling post body: %s", err)
	}
	req, err := retryablehttp.NewRequest("POST", reqUrl, reqBody)
	if err != nil {
		return associationResp, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
	resp, err := c.Client.Do(req)
	if err != nil {
		return associationResp, fmt.Errorf("error making request: %s", err)
	}
	defer resp.Body.Close()
	associationRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return associationResp, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 200 {
		return associationResp, fmt.Errorf("error returned by endpoint: %s", associationRawBody)
	}
	err = json.Unmarshal(associationRawBody, &associationResp)
	if err != nil {
		return associationResp, fmt.Errorf("error parsing body: %s", err)
	}
	return associationResp, nil
}
