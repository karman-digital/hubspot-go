package batchemails

import (
	"encoding/json"
	"fmt"
	"net/http"

	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (c *BatchEmailsService) BatchUpdate(body hubspotmodels.BatchUpdateBody) (hubspotmodels.BatchResponse, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return hubspotmodels.BatchResponse{}, fmt.Errorf("error marshalling post body: %s", err)
	}
	resp, err := c.SendRequest(http.MethodPost, "/crm/v3/objects/emails/batch/update", reqBody)
	if err != nil {
		return hubspotmodels.BatchResponse{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleBatchResponse(resp, http.MethodPatch)
}

func (c *BatchEmailsService) BatchCreate(body hubspotmodels.BatchCreateBody) (hubspotmodels.BatchResponse, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return hubspotmodels.BatchResponse{}, fmt.Errorf("error marshalling post body: %s", err)
	}
	resp, err := c.SendRequest(http.MethodPost, "/crm/v3/objects/emails/batch/create", reqBody)
	if err != nil {
		return hubspotmodels.BatchResponse{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleBatchResponse(resp, http.MethodPost)
}

func (c *BatchEmailsService) BatchGet(body hubspotmodels.BatchGetBody) (hubspotmodels.BatchResponse, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return hubspotmodels.BatchResponse{}, fmt.Errorf("error marshalling post body: %s", err)
	}
	resp, err := c.SendRequest(http.MethodPost, "/crm/v3/objects/emails/batch/read", reqBody)
	if err != nil {
		return hubspotmodels.BatchResponse{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleBatchResponse(resp, http.MethodGet)
}

func (c *BatchEmailsService) BatchDelete(body hubspotmodels.BatchDeleteBody) error {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("error marshalling post body: %s", err)
	}
	resp, err := c.SendRequest(http.MethodPost, "/crm/v3/objects/emails/batch/archive", reqBody)
	if err != nil {
		return fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleDeleteResponse(resp)
}
