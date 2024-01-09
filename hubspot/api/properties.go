package hubspot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
)

func (c *credentials) CreatePropertyGroup(propertyGroup hubspotmodels.PropertyGroupBody, objectType string) error {
	posturl := fmt.Sprintf("https://api.hubapi.com/crm/v3/properties/%s/groups", objectType)
	body, err := json.Marshal(propertyGroup)
	if err != nil {
		return fmt.Errorf("error marshalling body: %s", err)
	}
	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("error creating request: %s", err)
	}
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
	r.Header.Set("User-Agent", "Hatch Integration")
	r.Header.Set("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		return fmt.Errorf("error making post request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		if resp.StatusCode == http.StatusConflict {
			return ErrAlreadyExists
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("error reading body: %v", err)
		}
		return fmt.Errorf("error returned by endpoint. status code: %s, error: %v", resp.Status, string(body))
	}
	return nil
}

func (c *credentials) CreateProperty(objectType string, propertyData hubspotmodels.PropertyBody) error {
	posturl := fmt.Sprintf("https://api.hubapi.com/crm/v3/properties/%s", objectType)
	body, err := json.Marshal(propertyData)
	if err != nil {
		return fmt.Errorf("error marshalling body: %s", err)
	}
	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return fmt.Errorf("error creating request: %s", err)
	}
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		return fmt.Errorf("error making post request: %s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		if resp.StatusCode == http.StatusConflict {
			return ErrAlreadyExists
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("error reading body: %v", err)
		}
		return fmt.Errorf("error returned by endpoint. status code: %s, error: %v", resp.Status, string(body))
	}
	return nil
}
