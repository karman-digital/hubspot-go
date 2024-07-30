package lineItems

import (
	"encoding/json"
	"net/http"

	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (s *LineItemsService) CreateLineItem(body hubspotmodels.PostBody) (hubspotmodels.Result, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return hubspotmodels.Result{}, err
	}
	resp, err := s.SendRequest(http.MethodPost, "/crm/v3/objects/line_items", reqBody)
	if err != nil {
		return hubspotmodels.Result{}, err
	}
	return shared.HandleResponse(resp)
}
