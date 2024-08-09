package batchlineitems

import (
	"encoding/json"
	"fmt"
	"net/http"

	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (b *BatchLineItemService) BatchUpdate(body hubspotmodels.BatchUpdateBody) (hubspotmodels.BatchResponse, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return hubspotmodels.BatchResponse{}, fmt.Errorf("error marshalling post body: %s", err)
	}
	resp, err := b.SendRequest(http.MethodPost, "/crm/v3/objects/line_items/batch/update", reqBody)
	if err != nil {
		return hubspotmodels.BatchResponse{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleBatchResponse(resp, http.MethodPatch)
}

func (b *BatchLineItemService) BatchCreate(body hubspotmodels.BatchCreateBody) (hubspotmodels.BatchResponse, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return hubspotmodels.BatchResponse{}, fmt.Errorf("error marshalling post body: %s", err)
	}
	resp, err := b.SendRequest(http.MethodPost, "/crm/v3/objects/line_items/batch/create", reqBody)
	if err != nil {
		return hubspotmodels.BatchResponse{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleBatchResponse(resp, http.MethodPost)
}

func (b *BatchLineItemService) BatchGet(body hubspotmodels.BatchGetBody) (hubspotmodels.BatchResponse, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return hubspotmodels.BatchResponse{}, fmt.Errorf("error marshalling post body: %s", err)
	}
	resp, err := b.SendRequest(http.MethodPost, "/crm/v3/objects/line_items/batch/read", reqBody)
	if err != nil {
		return hubspotmodels.BatchResponse{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleBatchResponse(resp, http.MethodGet)
}
