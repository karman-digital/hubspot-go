package deals

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/hashicorp/go-retryablehttp"
	crmmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm"
	sharedmodels "github.com/karman-digital/hubspot/hubspot/api/models/shared"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (c *DealService) CreateDeal(body crmmodels.PostBody) (crmmodels.Result, error) {
	var respStruct crmmodels.Result
	reqUrl := "https://api.hubapi.com/crm/v3/objects/deals"
	reqBody, err := json.Marshal(body)
	if err != nil {
		return respStruct, fmt.Errorf("error marshalling post body: %s", err)
	}
	req, err := retryablehttp.NewRequest("POST", reqUrl, reqBody)
	if err != nil {
		return respStruct, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken()))
	resp, err := c.Client().Do(req)
	if err != nil {
		return respStruct, fmt.Errorf("error making request: %s", err)
	}
	defer resp.Body.Close()
	contactRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return respStruct, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != http.StatusCreated {
		return respStruct, fmt.Errorf("error returned by endpoint: %s", contactRawBody)
	}
	err = json.Unmarshal(contactRawBody, &respStruct)
	if err != nil {
		return respStruct, fmt.Errorf("error parsing body: %s", err)
	}
	return respStruct, nil
}

func (c *DealService) UpdateDeal(id int, patchBody crmmodels.PatchBody) (crmmodels.Result, error) {
	var respStruct crmmodels.Result
	reqUrl := fmt.Sprintf("https://api.hubapi.com/crm/v3/objects/deals/%d", id)
	reqBody, err := json.Marshal(patchBody)
	if err != nil {
		return respStruct, fmt.Errorf("error marshalling patch body: %s", err)
	}
	req, err := retryablehttp.NewRequest("PATCH", reqUrl, reqBody)
	if err != nil {
		return respStruct, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken()))
	resp, err := c.Client().Do(req)
	if err != nil {
		return respStruct, fmt.Errorf("error making request: %s", err)
	}
	defer resp.Body.Close()
	contactRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return respStruct, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		return respStruct, fmt.Errorf("error returned by endpoint, status code:%d \nBody: %s", resp.StatusCode, contactRawBody)
	}
	err = json.Unmarshal(contactRawBody, &respStruct)
	if err != nil {
		return respStruct, fmt.Errorf("error parsing body: %s", err)
	}
	return respStruct, nil
}

func (c *DealService) SearchDeals(body crmmodels.SearchBody) (crmmodels.SearchResponse, error) {
	var respStruct crmmodels.SearchResponse
	reqUrl := "https://api.hubapi.com/crm/v3/objects/deals/search"
	reqBody, err := json.Marshal(body)
	if err != nil {
		return respStruct, fmt.Errorf("error marshalling search body: %s", err)
	}
	req, err := retryablehttp.NewRequest("POST", reqUrl, reqBody)
	if err != nil {
		return respStruct, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken()))
	resp, err := c.Client().Do(req)
	if err != nil {
		return respStruct, fmt.Errorf("error making request: %s", err)
	}
	defer resp.Body.Close()
	contactRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return respStruct, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 200 {
		return respStruct, fmt.Errorf("error returned by endpoint: %s", contactRawBody)
	}
	err = json.Unmarshal(contactRawBody, &respStruct)
	if err != nil {
		return respStruct, fmt.Errorf("error parsing body: %s", err)
	}
	return respStruct, nil
}

func (c *DealService) GetDeal(id int, opts ...sharedmodels.GetOptions) (crmmodels.Result, error) {
	var respStruct crmmodels.Result
	reqUrl := fmt.Sprintf("https://api.hubapi.com/crm/v3/objects/deals/%d", id)
	req, err := retryablehttp.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return respStruct, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken()))
	queryParams := url.Values{}
	if len(opts) != 0 {
		if len(opts[0].Properties) != 0 {
			for _, property := range opts[0].Properties {
				queryParams.Add("properties", property)
			}
		}
		if len(opts[0].PropertiesWithHistory) != 0 {
			for _, property := range opts[0].PropertiesWithHistory {
				queryParams.Add("propertiesWithHistory", property)
			}
		}
		if len(opts[0].Associations) != 0 {
			for _, association := range opts[0].Associations {
				queryParams.Add("associations", association)
			}
		}
		if opts[0].Archived {
			queryParams.Add("archived", "true")
		}
	}
	req.URL.RawQuery = queryParams.Encode()
	resp, err := c.Client().Do(req)
	if err != nil {
		return respStruct, fmt.Errorf("error making request: %s", err)
	}
	defer resp.Body.Close()
	contactRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return respStruct, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 200 {
		return respStruct, fmt.Errorf("error returned by endpoint: %s", contactRawBody)
	}
	err = json.Unmarshal(contactRawBody, &respStruct)
	if err != nil {
		return respStruct, fmt.Errorf("error parsing body: %s", err)
	}
	return respStruct, nil
}

func (c *DealService) DeleteDeal(id int) (err error) {
	resp, err := c.SendRequest(http.MethodDelete, fmt.Sprintf("/crm/v3/objects/deals/%d", id), nil)
	if err != nil {
		return fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleDeleteResponse(resp)
}

func (c *DealService) GetDealByUniqueProperty(value string, opts ...sharedmodels.GetOptions) (crmmodels.Result, error) {
	if opts[0].IdProperty == "" {
		return crmmodels.Result{}, fmt.Errorf("idProperty must be set for unique property search")
	}
	resp, err := c.SendRequest(http.MethodGet, fmt.Sprintf("/crm/v3/objects/deals/%s", value), nil, opts...)
	if err != nil {
		return shared.HandleError(resp, err)
	}
	return shared.HandleResponse(resp)
}

func (c *DealService) GetDeals(opts ...sharedmodels.GetOptions) (crmmodels.ListResponse, error) {
	resp, err := c.SendRequest(http.MethodGet, "/crm/v3/objects/deals", nil, opts...)
	if err != nil {
		return crmmodels.ListResponse{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleListResponse(resp)
}
