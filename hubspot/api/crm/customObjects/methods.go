package customObjects

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/hashicorp/go-retryablehttp"
	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (c *CustomObjectService) CreateCustomObject(body hubspotmodels.PostBody, objectType string) (hubspotmodels.Result, error) {
	var respStruct hubspotmodels.Result
	reqUrl := fmt.Sprintf("https://api.hubapi.com/crm/v3/objects/%s", objectType)
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

func (c *CustomObjectService) UpdateCustomObject(id int, patchBody hubspotmodels.PatchBody, objectType string) (hubspotmodels.Result, error) {
	var respStruct hubspotmodels.Result
	reqUrl := fmt.Sprintf("https://api.hubapi.com/crm/v3/objects/%s/%d", objectType, id)
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

func (c *CustomObjectService) UpdateCustomObjectByUniqueId(id, idProperty string, patchBody hubspotmodels.PatchBody, objectType string) (hubspotmodels.Result, error) {
	reqBody, err := json.Marshal(patchBody)
	if err != nil {
		return hubspotmodels.Result{}, fmt.Errorf("error marshalling patch body: %s", err)
	}
	resp, err := c.SendRequest(http.MethodPatch, fmt.Sprintf("/crm/v3/objects/%s/%s?idProperty=%s", objectType, id, idProperty), reqBody)
	if err != nil {
		return hubspotmodels.Result{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleResponse(resp)
}

func (c *CustomObjectService) SearchCustomObjects(body hubspotmodels.SearchBody, objectType string) (hubspotmodels.SearchResponse, error) {
	var respStruct hubspotmodels.SearchResponse
	reqUrl := fmt.Sprintf("https://api.hubapi.com/crm/v3/objects/%s/search", objectType)
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

func (c *CustomObjectService) GetCustomObject(id int, objectType string, opts ...hubspotmodels.GetOptions) (hubspotmodels.Result, error) {
	var respStruct hubspotmodels.Result
	reqUrl := fmt.Sprintf("https://api.hubapi.com/crm/v3/objects/%s/%d", objectType, id)
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

func (c *CustomObjectService) GetCustomObjectByUniqueProperty(id string, objectType string, opts ...hubspotmodels.GetOptions) (hubspotmodels.Result, error) {
	var respStruct hubspotmodels.Result
	reqUrl := fmt.Sprintf("https://api.hubapi.com/crm/v3/objects/%s/%s", objectType, id)
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
		if len(opts[0].IdProperty) != 0 {
			queryParams.Add("idProperty", opts[0].IdProperty)
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
		if resp.StatusCode == 404 {
			return respStruct, shared.ErrResourceNotFound
		}
		return respStruct, fmt.Errorf("error returned by endpoint: %s", contactRawBody)
	}
	err = json.Unmarshal(contactRawBody, &respStruct)
	if err != nil {
		return respStruct, fmt.Errorf("error parsing body: %s", err)
	}
	return respStruct, nil
}

func (c *CustomObjectService) GetCustomObjects(objectType string, opts ...hubspotmodels.GetOptions) (hubspotmodels.ListResponse, error) {
	var respStruct hubspotmodels.ListResponse
	reqUrl := fmt.Sprintf("https://api.hubapi.com/crm/v3/objects/%s", objectType)
	req, err := retryablehttp.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return respStruct, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken()))
	queryParams := url.Values{}
	if len(opts) != 0 {
		if opts[0].After != "" {
			queryParams.Add("after", opts[0].After)
		}
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
