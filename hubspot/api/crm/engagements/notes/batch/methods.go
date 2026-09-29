package batchnotes

import (
	"encoding/json"
	"fmt"
	"net/http"

	crmmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (c *BatchNotesService) BatchUpdate(body crmmodels.BatchUpdateBody) (crmmodels.BatchResponse, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return crmmodels.BatchResponse{}, fmt.Errorf("error marshalling post body: %s", err)
	}
	resp, err := c.SendRequest(http.MethodPost, "/crm/objects/2026-09/notes/batch/update", reqBody)
	if err != nil {
		return crmmodels.BatchResponse{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleBatchResponse(resp, http.MethodPatch)
}

func (c *BatchNotesService) BatchCreate(body crmmodels.BatchCreateBody) (crmmodels.BatchResponse, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return crmmodels.BatchResponse{}, fmt.Errorf("error marshalling post body: %s", err)
	}
	resp, err := c.SendRequest(http.MethodPost, "/crm/objects/2026-09/notes/batch/create", reqBody)
	if err != nil {
		return crmmodels.BatchResponse{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleBatchResponse(resp, http.MethodPost)
}

func (c *BatchNotesService) BatchGet(body crmmodels.BatchGetBody) (crmmodels.BatchResponse, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return crmmodels.BatchResponse{}, fmt.Errorf("error marshalling post body: %s", err)
	}
	resp, err := c.SendRequest(http.MethodPost, "/crm/objects/2026-09/notes/batch/read", reqBody)
	if err != nil {
		return crmmodels.BatchResponse{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleBatchResponse(resp, http.MethodGet)
}

func (c *BatchNotesService) BatchDelete(body crmmodels.BatchDeleteBody) error {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("error marshalling post body: %s", err)
	}
	resp, err := c.SendRequest(http.MethodPost, "/crm/objects/2026-09/notes/batch/archive", reqBody)
	if err != nil {
		return fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleDeleteResponse(resp)
}
