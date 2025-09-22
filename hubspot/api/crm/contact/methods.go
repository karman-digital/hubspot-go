package contact

import (
	"encoding/json"
	"fmt"
	"net/http"

	crmmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm"
	sharedmodels "github.com/karman-digital/hubspot/hubspot/api/models/shared"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (c *ContactService) CreateContact(body crmmodels.PostBody) (crmmodels.Result, error) {
	reqUrl := "/crm/v3/objects/contacts"
	reqBody, err := json.Marshal(body)
	if err != nil {
		return crmmodels.Result{}, fmt.Errorf("error marshalling post body: %s", err)
	}
	resp, err := c.SendRequest(http.MethodPost, reqUrl, reqBody)
	if err != nil {
		return crmmodels.Result{}, fmt.Errorf("error creating request: %s", err)
	}
	return shared.HandleResponse(resp)
}

func (c *ContactService) UpdateContact(id int, patchBody crmmodels.PatchBody) (crmmodels.Result, error) {
	reqUrl := fmt.Sprintf("/crm/v3/objects/contacts/%d", id)
	reqBody, err := json.Marshal(patchBody)
	if err != nil {
		return crmmodels.Result{}, fmt.Errorf("error marshalling patch body: %s", err)
	}
	resp, err := c.SendRequest(http.MethodPatch, reqUrl, reqBody)
	if err != nil {
		return crmmodels.Result{}, fmt.Errorf("error creating request: %s", err)
	}
	return shared.HandleResponse(resp)
}

func (c *ContactService) SearchContacts(body crmmodels.SearchBody) (crmmodels.SearchResponse, error) {
	reqUrl := "/crm/v3/objects/contacts/search"
	reqBody, err := json.Marshal(body)
	if err != nil {
		return crmmodels.SearchResponse{}, fmt.Errorf("error marshalling search body: %s", err)
	}
	resp, err := c.SendRequest(http.MethodPost, reqUrl, reqBody)
	if err != nil {
		return crmmodels.SearchResponse{}, fmt.Errorf("error creating request: %s", err)
	}
	return shared.HandleSearchResponse(resp)
}

func (c *ContactService) GetContact(id int, opts ...sharedmodels.GetOptions) (crmmodels.Result, error) {
	reqUrl := fmt.Sprintf("/crm/v3/objects/contacts/%d", id)
	resp, err := c.SendRequest(http.MethodGet, reqUrl, nil, opts...)
	if err != nil {
		return crmmodels.Result{}, fmt.Errorf("error creating request: %s", err)
	}
	return shared.HandleResponse(resp)
}

func (c *ContactService) GetContactByUniqueProperty(value string, opts ...sharedmodels.GetOptions) (crmmodels.Result, error) {
	if opts[0].IdProperty == "" {
		return crmmodels.Result{}, fmt.Errorf("idProperty must be set for unique property search")
	}
	resp, err := c.SendRequest(http.MethodGet, fmt.Sprintf("/crm/v3/objects/contacts/%s", value), nil, opts...)
	if err != nil {
		return crmmodels.Result{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleResponse(resp)
}

func (c *ContactService) DeleteContact(id int) (err error) {
	resp, err := c.SendRequest(http.MethodDelete, fmt.Sprintf("/crm/v3/objects/contacts/%d", id), nil)
	if err != nil {
		return fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleDeleteResponse(resp)
}
