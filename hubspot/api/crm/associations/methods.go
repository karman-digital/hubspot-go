package associations

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/hashicorp/go-retryablehttp"
	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (c *AssociationService) CreateDefaultAssociation(fromObject, toObject string, fromId, toId int) (hubspotmodels.BatchResponse, error) {
	var associationResp hubspotmodels.BatchResponse
	reqUrl := fmt.Sprintf("https://api.hubapi.com/crm/v4/objects/%s/%d/associations/default/%s/%d", fromObject, fromId, toObject, toId)
	req, err := retryablehttp.NewRequest("PUT", reqUrl, nil)
	if err != nil {
		return associationResp, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken()))
	resp, err := c.Client().Do(req)
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

func (c *AssociationService) GetAssociations(fromObject, toObject string, id int) (hubspotmodels.AssociationGetResponse, error) {
	var association hubspotmodels.AssociationGetResponse
	reqUrl := fmt.Sprintf("https://api.hubapi.com/crm/v4/objects/%s/%d/associations/%s", fromObject, id, toObject)
	req, err := retryablehttp.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return association, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken()))
	resp, err := c.Client().Do(req)
	if err != nil {
		return association, fmt.Errorf("error making request: %s", err)
	}
	defer resp.Body.Close()
	associationRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return association, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 200 {
		return association, fmt.Errorf("error returned by endpoint: %s", associationRawBody)
	}
	err = json.Unmarshal(associationRawBody, &association)
	if err != nil {
		return association, fmt.Errorf("error parsing body: %s", err)
	}
	return association, nil
}

func (c *AssociationService) BatchCreateDefaultAssociations(fromObject, toObject string, associations hubspotmodels.BatchCreateDefaultAssociationsBody) (hubspotmodels.BatchResponse, error) {
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
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken()))
	resp, err := c.Client().Do(req)
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

func (c *AssociationService) BatchGetAssociations(fromObject, toObject string, body hubspotmodels.BatchGetAssociationsBody) (hubspotmodels.BatchAssociationGetResponse, error) {
	var batchResp hubspotmodels.BatchAssociationGetResponse
	reqUrl := fmt.Sprintf("https://api.hubapi.com/crm/v4/associations/%s/%s/batch/read", fromObject, toObject)
	reqBody, err := json.Marshal(body)
	if err != nil {
		return batchResp, fmt.Errorf("error marshalling post body: %s", err)
	}
	req, err := retryablehttp.NewRequest("POST", reqUrl, reqBody)
	if err != nil {
		return batchResp, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken()))
	resp, err := c.Client().Do(req)
	if err != nil {
		return batchResp, fmt.Errorf("error making request: %s", err)
	}
	defer resp.Body.Close()
	contactRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return batchResp, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 200 && resp.StatusCode != 207 {
		var errorResp hubspotmodels.ErrorResponseBody
		err := json.Unmarshal(contactRawBody, &errorResp)
		if err != nil {
			return batchResp, fmt.Errorf("error parsing error body: %s", err)
		}
		return batchResp, shared.HandleBatchResponseCodes(errorResp, resp.StatusCode)
	}
	err = json.Unmarshal(contactRawBody, &batchResp)
	if err != nil {
		return batchResp, fmt.Errorf("error parsing body: %s", err)
	}
	if resp.StatusCode == 207 {
		return batchResp, shared.ErrBatchGet
	}
	return batchResp, nil
}

func (c *AssociationService) BatchCreateAssociations(fromObject, toObject string, body hubspotmodels.BatchCreateAssociationsBody) (hubspotmodels.BatchResponse, error) {
	var batchResp hubspotmodels.BatchResponse
	reqUrl := fmt.Sprintf("https://api.hubapi.com/crm/v4/associations/%s/%s/batch/create", fromObject, toObject)
	reqBody, err := json.Marshal(body)
	if err != nil {
		return batchResp, fmt.Errorf("error marshalling post body: %s", err)
	}
	req, err := retryablehttp.NewRequest("POST", reqUrl, reqBody)
	if err != nil {
		return batchResp, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken()))
	resp, err := c.Client().Do(req)
	if err != nil {
		return batchResp, fmt.Errorf("error making request: %s", err)
	}
	defer resp.Body.Close()
	associationRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return batchResp, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 201 {
		return batchResp, fmt.Errorf("error returned by endpoint: %s", associationRawBody)
	}
	err = json.Unmarshal(associationRawBody, &batchResp)
	if err != nil {
		return batchResp, fmt.Errorf("error parsing body: %s", err)
	}
	return batchResp, nil
}

func (c *AssociationService) CreateAssociation(fromObject, toObject, fromObjectType, toObjectType string, body []hubspotmodels.AssociationType) (hubspotmodels.Result, error) {
	var associationResp hubspotmodels.Result
	reqUrl := fmt.Sprintf("https://api.hubapi.com/crm/v4/objects/%s/%s/associations/%s/%s", fromObjectType, fromObject, toObjectType, toObject)
	reqBody, err := json.Marshal(body)
	if err != nil {
		return associationResp, fmt.Errorf("error marshalling post body: %s", err)
	}
	req, err := retryablehttp.NewRequest("PUT", reqUrl, reqBody)
	if err != nil {
		return associationResp, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken()))
	resp, err := c.Client().Do(req)
	if err != nil {
		return associationResp, fmt.Errorf("error making request: %s", err)
	}
	defer resp.Body.Close()
	associationRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return associationResp, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 201 {
		return associationResp, fmt.Errorf("error returned by endpoint: %s", associationRawBody)
	}
	err = json.Unmarshal(associationRawBody, &associationResp)
	if err != nil {
		return associationResp, fmt.Errorf("error parsing body: %s", err)
	}
	return associationResp, nil
}

func (c *AssociationService) BatchArchiveAssociationLabels(fromObject, toObject string, body hubspotmodels.BatchCreateAssociationsBody) error {
	reqUrl := fmt.Sprintf("https://api.hubapi.com/crm/v4/associations/%s/%s/batch/labels/archive", fromObject, toObject)
	reqBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("error marshalling post body: %s", err)
	}
	req, err := retryablehttp.NewRequest("POST", reqUrl, reqBody)
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
	associationRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("error returned by endpoint: %s", associationRawBody)
	}
	return nil
}
