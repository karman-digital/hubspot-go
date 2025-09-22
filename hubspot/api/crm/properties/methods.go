package properties

import (
	"encoding/json"
	"fmt"
	propertiesmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm/properties"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
	"io"
	"net/http"
)

func (c *PropertiesService) CreatePropertyGroup(propertyGroup propertiesmodels.PropertyGroupBody, objectType string) error {
	reqBody, err := json.Marshal(propertyGroup)
	if err != nil {
		return fmt.Errorf("error marshalling body: %s", err)
	}
	resp, err := c.SendRequest(http.MethodPost, fmt.Sprintf("/crm/v3/properties/%s/groups", objectType), reqBody)
	if err != nil {
		return err
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

func (c *PropertiesService) CreateProperty(objectType string, propertyData propertiesmodels.PropertyBody) error {
	reqBody, err := json.Marshal(propertyData)
	if err != nil {
		return fmt.Errorf("error marshalling body: %s", err)
	}
	resp, err := c.SendRequest(http.MethodPost, fmt.Sprintf("/crm/v3/properties/%s", objectType), reqBody)
	if err != nil {
		return err
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

func (c *PropertiesService) GetProperty(objectType string, propertyName string) (propertiesmodels.PropertyResponse, error) {
	resp, err := c.SendRequest(http.MethodGet, fmt.Sprintf("/crm/v3/properties/%s/%s", objectType, propertyName), nil)
	if err != nil {
		return propertiesmodels.PropertyResponse{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return propertiesmodels.PropertyResponse{}, fmt.Errorf("error reading body: %v", err)
		}
		return propertiesmodels.PropertyResponse{}, fmt.Errorf("error returned by endpoint. status code: %s, error: %v", resp.Status, string(body))
	}
	var propertyResponse propertiesmodels.PropertyResponse
	err = json.NewDecoder(resp.Body).Decode(&propertyResponse)
	if err != nil {
		return propertiesmodels.PropertyResponse{}, fmt.Errorf("error decoding response: %s", err)
	}
	return propertyResponse, nil
}

func (c *PropertiesService) UpdateProperty(objectType string, propertyName string, propertyData propertiesmodels.PropertyBody) (propertiesmodels.PropertyResponse, error) {
	body, err := json.Marshal(propertyData)
	if err != nil {
		return propertiesmodels.PropertyResponse{}, fmt.Errorf("error marshalling body: %s", err)
	}
	resp, err := c.SendRequest(http.MethodPatch, fmt.Sprintf("/crm/v3/properties/%s/%s", objectType, propertyName), body)
	if err != nil {
		return propertiesmodels.PropertyResponse{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return propertiesmodels.PropertyResponse{}, fmt.Errorf("error reading body: %v", err)
		}
		return propertiesmodels.PropertyResponse{}, fmt.Errorf("error returned by endpoint. status code: %s, error: %v", resp.Status, string(body))
	}
	var propertyResponse propertiesmodels.PropertyResponse
	if err := json.NewDecoder(resp.Body).Decode(&propertyResponse); err != nil {
		return propertiesmodels.PropertyResponse{}, fmt.Errorf("error decoding response: %s", err)
	}
	return propertyResponse, nil
}
