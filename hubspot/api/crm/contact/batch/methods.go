package batchcontact

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/hashicorp/go-retryablehttp"
	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (c *BatchContactService) BatchUpdate(body hubspotmodels.BatchUpdateBody) (hubspotmodels.BatchResponse, error) {
	var contactResp hubspotmodels.BatchResponse
	reqUrl := "https://api.hubapi.com/crm/v3/objects/contacts/batch/update"
	reqBody, err := json.Marshal(body)
	if err != nil {
		return contactResp, fmt.Errorf("error marshalling post body: %s", err)
	}
	req, err := retryablehttp.NewRequest("POST", reqUrl, reqBody)
	if err != nil {
		return contactResp, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.creds.AccessToken))
	resp, err := c.creds.Client.Do(req)
	if err != nil {
		return contactResp, fmt.Errorf("error making request: %s", err)
	}
	defer resp.Body.Close()
	contactRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return contactResp, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 200 && resp.StatusCode != 207 {
		var errorResp hubspotmodels.ErrorResponseBody
		err := json.Unmarshal(contactRawBody, &errorResp)
		if err != nil {
			return contactResp, fmt.Errorf("error parsing error body: %s", err)
		}
		return contactResp, shared.HandleBatchResponseCodes(errorResp, resp.StatusCode)
	}
	err = json.Unmarshal(contactRawBody, &contactResp)
	if err != nil {
		return contactResp, fmt.Errorf("error parsing body: %s", err)
	}
	if resp.StatusCode == 207 {
		return contactResp, shared.ErrBatchCreate
	}
	return contactResp, nil
}

func (c *BatchContactService) BatchCreate(body hubspotmodels.BatchCreateBody) (hubspotmodels.BatchResponse, error) {
	var contactResp hubspotmodels.BatchResponse
	reqUrl := "https://api.hubapi.com/crm/v3/objects/contacts/batch/create"
	reqBody, err := json.Marshal(body)
	if err != nil {
		return contactResp, fmt.Errorf("error marshalling post body: %s", err)
	}
	req, err := retryablehttp.NewRequest("POST", reqUrl, reqBody)
	if err != nil {
		return contactResp, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.creds.AccessToken))
	resp, err := c.creds.Client.Do(req)
	if err != nil {
		return contactResp, fmt.Errorf("error making request: %s", err)
	}
	defer resp.Body.Close()
	contactRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return contactResp, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 201 && resp.StatusCode != 207 {
		var errorResp hubspotmodels.ErrorResponseBody
		err := json.Unmarshal(contactRawBody, &errorResp)
		if err != nil {
			return contactResp, fmt.Errorf("error parsing error body: %s", err)
		}
		return contactResp, shared.HandleBatchResponseCodes(errorResp, resp.StatusCode)
	}
	err = json.Unmarshal(contactRawBody, &contactResp)
	if err != nil {
		return contactResp, fmt.Errorf("error parsing body: %s", err)
	}
	if resp.StatusCode == 207 {
		return contactResp, shared.ErrBatchCreate
	}
	return contactResp, nil
}

func (c *BatchContactService) BatchGet(body hubspotmodels.BatchGetBody) (hubspotmodels.BatchResponse, error) {
	var contactResp hubspotmodels.BatchResponse
	reqUrl := "https://api.hubapi.com/crm/v3/objects/contacts/batch/read"
	reqBody, err := json.Marshal(body)
	if err != nil {
		return contactResp, fmt.Errorf("error marshalling post body: %s", err)
	}
	req, err := retryablehttp.NewRequest("POST", reqUrl, reqBody)
	if err != nil {
		return contactResp, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.creds.AccessToken))
	resp, err := c.creds.Client.Do(req)
	if err != nil {
		return contactResp, fmt.Errorf("error making request: %s", err)
	}
	defer resp.Body.Close()
	contactRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return contactResp, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 200 && resp.StatusCode != 207 {
		var errorResp hubspotmodels.ErrorResponseBody
		err := json.Unmarshal(contactRawBody, &errorResp)
		if err != nil {
			return contactResp, fmt.Errorf("error parsing error body: %s", err)
		}
		return contactResp, shared.HandleBatchResponseCodes(errorResp, resp.StatusCode)
	}
	err = json.Unmarshal(contactRawBody, &contactResp)
	if err != nil {
		return contactResp, fmt.Errorf("error parsing body: %s", err)
	}
	if resp.StatusCode == 207 {
		return contactResp, shared.ErrBatchGet
	}
	return contactResp, nil
}
