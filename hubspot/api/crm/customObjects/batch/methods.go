package batchcustomobjects

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/hashicorp/go-retryablehttp"
	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (c *BatchCustomObjectService) BatchUpdate(body hubspotmodels.BatchUpdateBody, objectType string) (hubspotmodels.BatchResponse, error) {
	var dealResp hubspotmodels.BatchResponse
	reqUrl := fmt.Sprintf("https://api.hubapi.com/crm/v3/objects/%s/batch/update", objectType)
	reqBody, err := json.Marshal(body)
	if err != nil {
		return dealResp, fmt.Errorf("error marshalling post body: %s", err)
	}
	req, err := retryablehttp.NewRequest("POST", reqUrl, reqBody)
	if err != nil {
		return dealResp, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken()))
	resp, err := c.Client().Do(req)
	if err != nil {
		return dealResp, fmt.Errorf("error making request: %s", err)
	}
	defer resp.Body.Close()
	contactRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return dealResp, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 200 && resp.StatusCode != 207 {
		var errorResp hubspotmodels.ErrorResponseBody
		err := json.Unmarshal(contactRawBody, &errorResp)
		if err != nil {
			return dealResp, fmt.Errorf("error parsing error body: %s", err)
		}
		return dealResp, shared.HandleBatchResponseCodes(errorResp, resp.StatusCode)
	}
	err = json.Unmarshal(contactRawBody, &dealResp)
	if err != nil {
		return dealResp, fmt.Errorf("error parsing body: %s", err)
	}
	if resp.StatusCode == 207 {
		return dealResp, shared.ErrBatchCreate
	}
	return dealResp, nil
}

func (c *BatchCustomObjectService) BatchCreate(body hubspotmodels.BatchCreateBody, objectType string) (hubspotmodels.BatchResponse, error) {
	var dealResp hubspotmodels.BatchResponse
	reqUrl := fmt.Sprintf("https://api.hubapi.com/crm/v3/objects/%s/batch/create", objectType)
	reqBody, err := json.Marshal(body)
	if err != nil {
		return dealResp, fmt.Errorf("error marshalling post body: %s", err)
	}
	req, err := retryablehttp.NewRequest("POST", reqUrl, reqBody)
	if err != nil {
		return dealResp, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken()))
	resp, err := c.Client().Do(req)
	if err != nil {
		return dealResp, fmt.Errorf("error making request: %s", err)
	}
	defer resp.Body.Close()
	contactRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return dealResp, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 201 && resp.StatusCode != 207 {
		var errorResp hubspotmodels.ErrorResponseBody
		err := json.Unmarshal(contactRawBody, &errorResp)
		if err != nil {
			return dealResp, fmt.Errorf("error parsing error body: %s", err)
		}
		return dealResp, shared.HandleBatchResponseCodes(errorResp, resp.StatusCode)
	}
	err = json.Unmarshal(contactRawBody, &dealResp)
	if err != nil {
		return dealResp, fmt.Errorf("error parsing body: %s", err)
	}
	if resp.StatusCode == 207 {
		return dealResp, shared.ErrBatchCreate
	}
	return dealResp, nil
}

func (c *BatchCustomObjectService) BatchGet(body hubspotmodels.BatchGetBody, objectType string) (hubspotmodels.BatchResponse, error) {
	var dealResp hubspotmodels.BatchResponse
	reqUrl := fmt.Sprintf("https://api.hubapi.com/crm/v3/objects/%s/batch/read", objectType)
	reqBody, err := json.Marshal(body)
	if err != nil {
		return dealResp, fmt.Errorf("error marshalling post body: %s", err)
	}
	req, err := retryablehttp.NewRequest("POST", reqUrl, reqBody)
	if err != nil {
		return dealResp, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken()))
	resp, err := c.Client().Do(req)
	if err != nil {
		return dealResp, fmt.Errorf("error making request: %s", err)
	}
	defer resp.Body.Close()
	contactRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return dealResp, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 200 && resp.StatusCode != 207 {
		var errorResp hubspotmodels.ErrorResponseBody
		err := json.Unmarshal(contactRawBody, &errorResp)
		if err != nil {
			return dealResp, fmt.Errorf("error parsing error body: %s", err)
		}
		return dealResp, shared.HandleBatchResponseCodes(errorResp, resp.StatusCode)
	}
	err = json.Unmarshal(contactRawBody, &dealResp)
	if err != nil {
		return dealResp, fmt.Errorf("error parsing body: %s", err)
	}
	if resp.StatusCode == 207 {
		return dealResp, shared.ErrBatchGet
	}
	return dealResp, nil
}
