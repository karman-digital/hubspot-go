package company

import (
	"encoding/json"
	"fmt"
	"net/http"

	crmmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm"
	sharedmodels "github.com/karman-digital/hubspot/hubspot/api/models/shared"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (c *CompanyService) CreateCompany(body crmmodels.PostBody) (crmmodels.Result, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return crmmodels.Result{}, fmt.Errorf("error marshalling post body: %s", err)
	}
	resp, err := c.SendRequest(http.MethodPost, "/crm/v3/objects/companies", reqBody)
	if err != nil {
		return crmmodels.Result{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleResponse(resp)
}

func (c *CompanyService) UpdateCompany(id int, body crmmodels.PatchBody) (crmmodels.Result, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return crmmodels.Result{}, fmt.Errorf("error marshalling patch body: %s", err)
	}
	resp, err := c.SendRequest(http.MethodPatch, fmt.Sprintf("/crm/v3/objects/companies/%d", id), reqBody)
	if err != nil {
		return crmmodels.Result{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleResponse(resp)
}

func (c *CompanyService) GetCompany(id int, opts ...sharedmodels.GetOptions) (crmmodels.Result, error) {
	resp, err := c.SendRequest(http.MethodGet, fmt.Sprintf("/crm/v3/objects/companies/%d", id), nil, opts...)
	if err != nil {
		return crmmodels.Result{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleResponse(resp)
}

func (c *CompanyService) SearchCompanies(body crmmodels.SearchBody) (crmmodels.SearchResponse, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return crmmodels.SearchResponse{}, fmt.Errorf("error marshalling post body: %s", err)
	}
	resp, err := c.SendRequest(http.MethodPost, "/crm/v3/objects/companies/search", reqBody)
	if err != nil {
		return crmmodels.SearchResponse{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleSearchResponse(resp)
}

func (c *CompanyService) GetCompanyByUniqueProperty(value string, opts ...sharedmodels.GetOptions) (crmmodels.Result, error) {
	if opts[0].IdProperty == "" {
		return crmmodels.Result{}, fmt.Errorf("idProperty must be set for unique property search")
	}
	resp, err := c.SendRequest(http.MethodGet, fmt.Sprintf("/crm/v3/objects/companies/%s", value), nil, opts...)
	if err != nil {
		return crmmodels.Result{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleResponse(resp)
}

func (c *CompanyService) DeleteCompany(id int) (err error) {
	resp, err := c.SendRequest(http.MethodDelete, fmt.Sprintf("/crm/v3/objects/companies/%d", id), nil)
	if err != nil {
		return fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleDeleteResponse(resp)
}

func (c *CompanyService) GetCompanies(opts ...sharedmodels.GetOptions) (crmmodels.ListResponse, error) {
	resp, err := c.SendRequest(http.MethodGet, "/crm/v3/objects/companies", nil, opts...)
	if err != nil {
		return crmmodels.ListResponse{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleListResponse(resp)
}
