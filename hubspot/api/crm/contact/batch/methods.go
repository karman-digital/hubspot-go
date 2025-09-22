package batchcontact

import (
	"encoding/json"
	"fmt"
	"net/http"

	crmmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (c *BatchContactService) BatchUpdate(body crmmodels.BatchUpdateBody) (crmmodels.BatchResponse, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return crmmodels.BatchResponse{}, fmt.Errorf("error marshalling post body: %s", err)
	}
	resp, err := c.SendRequest(http.MethodPost, "/crm/v3/objects/contacts/batch/update", reqBody)
	if err != nil {
		return crmmodels.BatchResponse{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleBatchResponse(resp, http.MethodPatch)
}

func (c *BatchContactService) BatchCreate(body crmmodels.BatchCreateBody) (crmmodels.BatchResponse, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return crmmodels.BatchResponse{}, fmt.Errorf("error marshalling post body: %s", err)
	}
	resp, err := c.SendRequest(http.MethodPost, "/crm/v3/objects/contacts/batch/create", reqBody)
	if err != nil {
		return crmmodels.BatchResponse{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleBatchResponse(resp, http.MethodPost)
}

func (c *BatchContactService) BatchGet(body crmmodels.BatchGetBody) (crmmodels.BatchResponse, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return crmmodels.BatchResponse{}, fmt.Errorf("error marshalling post body: %s", err)
	}
	resp, err := c.SendRequest(http.MethodPost, "/crm/v3/objects/contacts/batch/read", reqBody)
	if err != nil {
		return crmmodels.BatchResponse{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleBatchResponse(resp, http.MethodGet)
}

func (c *BatchContactService) BatchDelete(body crmmodels.BatchDeleteBody) error {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("error marshalling post body: %s", err)
	}
	resp, err := c.SendRequest(http.MethodPost, "/crm/v3/objects/contacts/batch/archive", reqBody)
	if err != nil {
		return fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleDeleteResponse(resp)
}
