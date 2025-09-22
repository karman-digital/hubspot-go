package lineItems

import (
	"encoding/json"
	"fmt"
	"net/http"

	crmmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm"
	sharedmodels "github.com/karman-digital/hubspot/hubspot/api/models/shared"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (s *LineItemsService) CreateLineItem(body crmmodels.PostBody) (crmmodels.Result, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return crmmodels.Result{}, err
	}
	resp, err := s.SendRequest(http.MethodPost, "/crm/v3/objects/line_items", reqBody)
	if err != nil {
		return crmmodels.Result{}, err
	}
	return shared.HandleCreateResponse(resp)
}

func (s *LineItemsService) UpdateLineItem(id int, patchBody crmmodels.PatchBody) (crmmodels.Result, error) {
	reqBody, err := json.Marshal(patchBody)
	if err != nil {
		return crmmodels.Result{}, err
	}
	resp, err := s.SendRequest(http.MethodPatch, fmt.Sprintf("/crm/v3/objects/line_items/%d", id), reqBody)
	if err != nil {
		return crmmodels.Result{}, err
	}
	return shared.HandleResponse(resp)
}

func (s *LineItemsService) GetLineItem(id int, opts ...sharedmodels.GetOptions) (crmmodels.Result, error) {
	resp, err := s.SendRequest(http.MethodGet, fmt.Sprintf("/crm/v3/objects/line_items/%d", id), nil, opts...)
	if err != nil {
		return crmmodels.Result{}, err
	}
	return shared.HandleResponse(resp)
}
