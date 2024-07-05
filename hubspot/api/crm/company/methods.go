package company

import (
	"encoding/json"
	"fmt"
	"net/http"

	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (c *CompanyService) CreateCompany(body hubspotmodels.PostBody) (hubspotmodels.Result, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return hubspotmodels.Result{}, fmt.Errorf("error marshalling post body: %s", err)
	}
	resp, err := c.SendRequest(http.MethodPost, "/crm/v3/objects/companies", reqBody)
	if err != nil {
		return hubspotmodels.Result{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleResponse(resp)
}

func (c *CompanyService) UpdateCompany(id int, body hubspotmodels.PatchBody) (hubspotmodels.Result, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return hubspotmodels.Result{}, fmt.Errorf("error marshalling patch body: %s", err)
	}
	resp, err := c.SendRequest(http.MethodPatch, fmt.Sprintf("/crm/v3/objects/companies/%d", id), reqBody)
	if err != nil {
		return hubspotmodels.Result{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleResponse(resp)
}

func (c *CompanyService) GetCompany(id int, opts ...hubspotmodels.GetOptions) (hubspotmodels.Result, error) {
	resp, err := c.SendRequest(http.MethodGet, fmt.Sprintf("/crm/v3/objects/companies/%d", id), nil, opts...)
	if err != nil {
		return hubspotmodels.Result{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleResponse(resp)
}

func (c *CompanyService) SearchCompanies(body hubspotmodels.SearchBody) (hubspotmodels.SearchResponse, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return hubspotmodels.SearchResponse{}, fmt.Errorf("error marshalling post body: %s", err)
	}
	resp, err := c.SendRequest(http.MethodPost, "/crm/v3/objects/companies/search", reqBody)
	if err != nil {
		return hubspotmodels.SearchResponse{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleSearchResponse(resp)
}
