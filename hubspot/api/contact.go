package hubspot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"

	"github.com/hashicorp/go-retryablehttp"
	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
)

func (c *credentials) CreateContact(body hubspotmodels.PostBody) (hubspotmodels.ContactResponse, error) {
	var contactResp hubspotmodels.ContactResponse
	reqUrl := "https://api.hubapi.com/crm/v3/objects/contacts"
	reqBody, err := json.Marshal(body)
	if err != nil {
		return contactResp, fmt.Errorf("error marshalling post body: %s", err)
	}
	req, err := retryablehttp.NewRequest("POST", reqUrl, reqBody)
	if err != nil {
		return contactResp, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
	resp, err := c.Client.Do(req)
	if err != nil {
		return contactResp, fmt.Errorf("error making request: %s", err)
	}
	defer resp.Body.Close()
	contactRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return contactResp, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 200 {
		return contactResp, fmt.Errorf("error returned by endpoint: %s", contactRawBody)
	}
	err = json.Unmarshal(contactRawBody, &contactResp)
	if err != nil {
		return contactResp, fmt.Errorf("error parsing body: %s", err)
	}
	return contactResp, nil
}

func (c *credentials) BatchCreateContact(body hubspotmodels.BatchCreateBody) (hubspotmodels.BatchResponse, error) {
	var contactResp hubspotmodels.BatchResponse
	reqUrl := "https://api.hubapi.com/crm/v3/objects/contacts/batch/create"
	reqBody, err := json.Marshal(body)
	if err != nil {
		return contactResp, fmt.Errorf("error marshalling post body: %s", err)
	}
	req, err := retryablehttp.NewRequest("POST", reqUrl, reqBody)
	if err != nil {
		return contactResp, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
	resp, err := c.Client.Do(req)
	if err != nil {
		return contactResp, fmt.Errorf("error making request: %s", err)
	}
	defer resp.Body.Close()
	contactRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return contactResp, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 201 && resp.StatusCode != 207 {
		var errorResp hubspotmodels.ErrorResponseBody
		err := json.Unmarshal(contactRawBody, &errorResp)
		if err != nil {
			return contactResp, fmt.Errorf("error parsing error body: %s", err)
		}
		return contactResp, handleBatchResponseCodes(errorResp, resp.StatusCode)
	}
	err = json.Unmarshal(contactRawBody, &contactResp)
	if err != nil {
		return contactResp, fmt.Errorf("error parsing body: %s", err)
	}
	if resp.StatusCode == 207 {
		return contactResp, ErrBatchCreate
	}
	return contactResp, nil
}

func (c *credentials) BatchGetContacts(body hubspotmodels.BatchGetBody) (hubspotmodels.BatchResponse, error) {
	var contactResp hubspotmodels.BatchResponse
	reqUrl := "https://api.hubapi.com/crm/v3/objects/contacts/batch/read"
	reqBody, err := json.Marshal(body)
	if err != nil {
		return contactResp, fmt.Errorf("error marshalling post body: %s", err)
	}
	req, err := retryablehttp.NewRequest("POST", reqUrl, reqBody)
	if err != nil {
		return contactResp, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
	resp, err := c.Client.Do(req)
	if err != nil {
		return contactResp, fmt.Errorf("error making request: %s", err)
	}
	defer resp.Body.Close()
	contactRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return contactResp, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 200 && resp.StatusCode != 207 {
		var errorResp hubspotmodels.ErrorResponseBody
		err := json.Unmarshal(contactRawBody, &errorResp)
		if err != nil {
			return contactResp, fmt.Errorf("error parsing error body: %s", err)
		}
		return contactResp, handleBatchResponseCodes(errorResp, resp.StatusCode)
	}
	err = json.Unmarshal(contactRawBody, &contactResp)
	if err != nil {
		return contactResp, fmt.Errorf("error parsing body: %s", err)
	}
	if resp.StatusCode == 207 {
		return contactResp, ErrBatchCreate
	}
	return contactResp, nil
}

func (c *credentials) BatchUpdateContacts(body hubspotmodels.BatchUpdateBody) (hubspotmodels.BatchResponse, error) {
	var contactResp hubspotmodels.BatchResponse
	reqUrl := "https://api.hubapi.com/crm/v3/objects/contacts/batch/update"
	reqBody, err := json.Marshal(body)
	if err != nil {
		return contactResp, fmt.Errorf("error marshalling post body: %s", err)
	}
	req, err := retryablehttp.NewRequest("POST", reqUrl, reqBody)
	if err != nil {
		return contactResp, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
	resp, err := c.Client.Do(req)
	if err != nil {
		return contactResp, fmt.Errorf("error making request: %s", err)
	}
	defer resp.Body.Close()
	contactRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return contactResp, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 200 && resp.StatusCode != 207 {
		var errorResp hubspotmodels.ErrorResponseBody
		err := json.Unmarshal(contactRawBody, &errorResp)
		if err != nil {
			return contactResp, fmt.Errorf("error parsing error body: %s", err)
		}
		return contactResp, handleBatchResponseCodes(errorResp, resp.StatusCode)
	}
	err = json.Unmarshal(contactRawBody, &contactResp)
	if err != nil {
		return contactResp, fmt.Errorf("error parsing body: %s", err)
	}
	if resp.StatusCode == 207 {
		return contactResp, ErrBatchCreate
	}
	return contactResp, nil
}

func (c *credentials) UpdateContact(id int, patchBody hubspotmodels.PatchBody) (hubspotmodels.ContactResponse, error) {
	var contactResp hubspotmodels.ContactResponse
	reqUrl := fmt.Sprintf("https://api.hubapi.com/crm/v3/objects/contacts/%d", id)
	reqBody, err := json.Marshal(patchBody)
	if err != nil {
		return contactResp, fmt.Errorf("error marshalling patch body: %s", err)
	}
	req, err := retryablehttp.NewRequest("PATCH", reqUrl, reqBody)
	if err != nil {
		return contactResp, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
	resp, err := c.Client.Do(req)
	if err != nil {
		return contactResp, fmt.Errorf("error making request: %s", err)
	}
	defer resp.Body.Close()
	contactRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return contactResp, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 200 {
		return contactResp, fmt.Errorf("error returned by endpoint, status code:%d \nBody: %s", resp.StatusCode, contactRawBody)
	}
	err = json.Unmarshal(contactRawBody, &contactResp)
	if err != nil {
		return contactResp, fmt.Errorf("error parsing body: %s", err)
	}
	return contactResp, nil
}

func (c *credentials) SearchContacts(body hubspotmodels.SearchBody) (hubspotmodels.SearchResponse, error) {
	var contactResp hubspotmodels.SearchResponse
	reqUrl := "https://api.hubapi.com/crm/v3/objects/contacts/search"
	reqBody, err := json.Marshal(body)
	if err != nil {
		return contactResp, fmt.Errorf("error marshalling search body: %s", err)
	}
	req, err := retryablehttp.NewRequest("POST", reqUrl, reqBody)
	if err != nil {
		return contactResp, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
	resp, err := c.Client.Do(req)
	if err != nil {
		return contactResp, fmt.Errorf("error making request: %s", err)
	}
	defer resp.Body.Close()
	contactRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return contactResp, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 200 {
		return contactResp, fmt.Errorf("error returned by endpoint: %s", contactRawBody)
	}
	err = json.Unmarshal(contactRawBody, &contactResp)
	if err != nil {
		return contactResp, fmt.Errorf("error parsing body: %s", err)
	}
	return contactResp, nil
}

func (c *credentials) GetContact(id int, opts ...hubspotmodels.ContactGetOptions) (hubspotmodels.ContactResponse, error) {
	var contactResp hubspotmodels.ContactResponse
	reqUrl := fmt.Sprintf("https://api.hubapi.com/crm/v3/objects/contacts/%d", id)
	req, err := retryablehttp.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return contactResp, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
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
	resp, err := c.Client.Do(req)
	if err != nil {
		return contactResp, fmt.Errorf("error making request: %s", err)
	}
	defer resp.Body.Close()
	contactRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return contactResp, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 200 {
		return contactResp, fmt.Errorf("error returned by endpoint: %s", contactRawBody)
	}
	err = json.Unmarshal(contactRawBody, &contactResp)
	if err != nil {
		return contactResp, fmt.Errorf("error parsing body: %s", err)
	}
	return contactResp, nil
}
