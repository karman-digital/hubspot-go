package associations

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	crmmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm"
	associationsmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm/associations"
	sharedmodels "github.com/karman-digital/hubspot/hubspot/api/models/shared"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (c *AssociationService) CreateDefaultAssociation(fromObject, toObject string, fromId, toId int) (crmmodels.BatchResponse, error) {
	var associationResp crmmodels.BatchResponse
	resp, err := c.SendRequest(
		http.MethodPut,
		fmt.Sprintf("/crm/v4/objects/%s/%d/associations/default/%s/%d", fromObject, fromId, toObject, toId),
		nil,
	)
	if err != nil {
		return associationResp, err
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

func (c *AssociationService) GetAssociations(fromObject, toObject string, id int) (associationsmodels.AssociationGetResponse, error) {
	var association associationsmodels.AssociationGetResponse
	resp, err := c.SendRequest(
		http.MethodGet,
		fmt.Sprintf("/crm/v4/objects/%s/%d/associations/%s", fromObject, id, toObject),
		nil,
	)
	if err != nil {
		return association, err
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

func (c *AssociationService) BatchCreateDefaultAssociations(fromObject, toObject string, associations associationsmodels.BatchCreateDefaultAssociationsBody) (crmmodels.BatchResponse, error) {
	var associationResp crmmodels.BatchResponse
	reqBody, err := json.Marshal(associations)
	if err != nil {
		return associationResp, fmt.Errorf("error marshalling post body: %s", err)
	}
	resp, err := c.SendRequest(
		http.MethodPost,
		fmt.Sprintf("/crm/v4/associations/%s/%s/batch/associate/default", fromObject, toObject),
		reqBody,
	)
	if err != nil {
		return associationResp, err
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

func (c *AssociationService) BatchGetAssociations(fromObject, toObject string, body associationsmodels.BatchGetAssociationsBody) (associationsmodels.BatchAssociationGetResponse, error) {
	var batchResp associationsmodels.BatchAssociationGetResponse
	reqBody, err := json.Marshal(body)
	if err != nil {
		return batchResp, fmt.Errorf("error marshalling post body: %s", err)
	}
	resp, err := c.SendRequest(
		http.MethodPost,
		fmt.Sprintf("/crm/v4/associations/%s/%s/batch/read", fromObject, toObject),
		reqBody,
	)
	if err != nil {
		return batchResp, err
	}
	defer resp.Body.Close()
	contactRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return batchResp, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 200 && resp.StatusCode != 207 {
		var errorResp sharedmodels.ErrorResponseBody
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

func (c *AssociationService) BatchCreateAssociations(fromObject, toObject string, body associationsmodels.BatchCreateAssociationsBody) (crmmodels.BatchResponse, error) {
	var batchResp crmmodels.BatchResponse
	reqBody, err := json.Marshal(body)
	if err != nil {
		return batchResp, fmt.Errorf("error marshalling post body: %s", err)
	}
	resp, err := c.SendRequest(
		http.MethodPost,
		fmt.Sprintf("/crm/v4/associations/%s/%s/batch/create", fromObject, toObject),
		reqBody,
	)
	if err != nil {
		return batchResp, err
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

func (c *AssociationService) CreateAssociation(fromObject, toObject, fromObjectType, toObjectType string, body []associationsmodels.AssociationType) (crmmodels.Result, error) {
	var associationResp crmmodels.Result
	reqBody, err := json.Marshal(body)
	if err != nil {
		return associationResp, fmt.Errorf("error marshalling post body: %s", err)
	}
	resp, err := c.SendRequest(
		http.MethodPut,
		fmt.Sprintf("/crm/v4/objects/%s/%s/associations/%s/%s", fromObjectType, fromObject, toObjectType, toObject),
		reqBody,
	)
	if err != nil {
		return associationResp, err
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

func (c *AssociationService) BatchArchiveAssociationLabels(fromObject, toObject string, body associationsmodels.BatchCreateAssociationsBody) error {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("error marshalling post body: %s", err)
	}
	resp, err := c.SendRequest(
		http.MethodPost,
		fmt.Sprintf("/crm/v4/associations/%s/%s/batch/labels/archive", fromObject, toObject),
		reqBody,
	)
	if err != nil {
		return err
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
