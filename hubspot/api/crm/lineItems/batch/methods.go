package batchlineitems

import (
	"encoding/json"
	"fmt"
	"net/http"

	crmmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (b *BatchLineItemService) BatchUpdate(body crmmodels.BatchUpdateBody) (crmmodels.BatchResponse, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return crmmodels.BatchResponse{}, fmt.Errorf("error marshalling post body: %s", err)
	}
	resp, err := b.SendRequest(http.MethodPost, "/crm/v3/objects/line_items/batch/update", reqBody)
	if err != nil {
		return crmmodels.BatchResponse{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleBatchResponse(resp, http.MethodPatch)
}

func (b *BatchLineItemService) BatchCreate(body crmmodels.BatchCreateBody) (crmmodels.BatchResponse, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return crmmodels.BatchResponse{}, fmt.Errorf("error marshalling post body: %s", err)
	}
	resp, err := b.SendRequest(http.MethodPost, "/crm/v3/objects/line_items/batch/create", reqBody)
	if err != nil {
		return crmmodels.BatchResponse{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleBatchResponse(resp, http.MethodPost)
}

func (b *BatchLineItemService) BatchGet(body crmmodels.BatchGetBody) (crmmodels.BatchResponse, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return crmmodels.BatchResponse{}, fmt.Errorf("error marshalling post body: %s", err)
	}
	resp, err := b.SendRequest(http.MethodPost, "/crm/v3/objects/line_items/batch/read", reqBody)
	if err != nil {
		return crmmodels.BatchResponse{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleBatchResponse(resp, http.MethodGet)
}

func (b *BatchLineItemService) BatchDelete(body crmmodels.BatchDeleteBody) error {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("error marshalling post body: %s", err)
	}
	resp, err := b.SendRequest(http.MethodPost, "/crm/v3/objects/line_items/batch/archive", reqBody)
	if err != nil {
		return fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleDeleteResponse(resp)
}
