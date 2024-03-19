package properties

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (c *PropertiesService) CreatePropertyGroup(propertyGroup hubspotmodels.PropertyGroupBody, objectType string) error {
	posturl := fmt.Sprintf("https://api.hubapi.com/crm/v3/properties/%s/groups", objectType)
	body, err := json.Marshal(propertyGroup)
	if err != nil {
		return fmt.Errorf("error marshalling body: %s", err)
	}
	req, err := retryablehttp.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken()))
	resp, err := c.Client().Do(req)
	if err != nil {
		return fmt.Errorf("error making post request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		if resp.StatusCode == http.StatusConflict {
			return shared.ErrResourceAlreadyExists
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("error reading body: %v", err)
		}
		return fmt.Errorf("error returned by endpoint. status code: %s, error: %v", resp.Status, string(body))
	}
	return nil
}

func (c *PropertiesService) CreateProperty(objectType string, propertyData hubspotmodels.PropertyBody) error {
	posturl := fmt.Sprintf("https://api.hubapi.com/crm/v3/properties/%s", objectType)
	body, err := json.Marshal(propertyData)
	if err != nil {
		return fmt.Errorf("error marshalling body: %s", err)
	}
	req, err := retryablehttp.NewRequest("POST", posturl, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken()))
	resp, err := c.Client().Do(req)
	if err != nil {
		return fmt.Errorf("error making post request: %s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		if resp.StatusCode == http.StatusConflict {
			return shared.ErrResourceAlreadyExists
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("error reading body: %v", err)
		}
		return fmt.Errorf("error returned by endpoint. status code: %s, error: %v", resp.Status, string(body))
	}
	return nil
}

func (c *PropertiesService) GetProperty(ObjectType string, PropertyName string) (hubspotmodels.PropertyResponse, error) {
	geturl := fmt.Sprintf("https://api.hubapi.com/crm/v3/properties/%s/%s", ObjectType, PropertyName)
	req, err := retryablehttp.NewRequest(http.MethodGet, geturl, nil)
	if err != nil {
		return hubspotmodels.PropertyResponse{}, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken()))
	resp, err := c.Client().Do(req)
	if err != nil {
		return hubspotmodels.PropertyResponse{}, fmt.Errorf("error making post request: %s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return hubspotmodels.PropertyResponse{}, fmt.Errorf("error reading body: %v", err)
		}
		return hubspotmodels.PropertyResponse{}, fmt.Errorf("error returned by endpoint. status code: %s, error: %v", resp.Status, string(body))
	}
	var propertyResponse hubspotmodels.PropertyResponse
	err = json.NewDecoder(resp.Body).Decode(&propertyResponse)
	if err != nil {
		return hubspotmodels.PropertyResponse{}, fmt.Errorf("error decoding response: %s", err)
	}
	return propertyResponse, nil
}

func (c *PropertiesService) UpdateProperty(ObjectType string, PropertyName string, propertyData hubspotmodels.PropertyBody) error {
	puturl := fmt.Sprintf("https://api.hubapi.com/crm/v3/properties/%s/%s", ObjectType, PropertyName)
	body, err := json.Marshal(propertyData)
	if err != nil {
		return fmt.Errorf("error marshalling body: %s", err)
	}
	req, err := retryablehttp.NewRequest(http.MethodPut, puturl, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken()))
	resp, err := c.Client().Do(req)
	if err != nil {
		return fmt.Errorf("error making post request: %s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("error reading body: %v", err)
		}
		return fmt.Errorf("error returned by endpoint. status code: %s, error: %v", resp.Status, string(body))
	}
	return nil
}
