package owners

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/hashicorp/go-retryablehttp"
	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
)

func (c *OwnerService) GetAllOwners() ([]hubspotmodels.Owner, error) {
	var allOwners []hubspotmodels.Owner
	after := ""
	for {
		ownerResponse, err := c.GetOwners(after)
		if err != nil {
			return nil, err
		}
		allOwners = append(allOwners, ownerResponse.Results...)
		if ownerResponse.Paging.Next.After != "" {
			after = ownerResponse.Paging.Next.After
		} else {
			break
		}
	}
	return allOwners, nil
}

func (c *OwnerService) GetOwners(after ...string) (hubspotmodels.OwnerResponse, error) {
	ownerResponse := hubspotmodels.OwnerResponse{}
	reqUrl := "https://api.hubapi.com/crm/v3/owners"
	req, err := retryablehttp.NewRequest("GET", reqUrl, strings.NewReader(""))
	if err != nil {
		return ownerResponse, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.creds.AccessToken))
	req.Header.Set("Content-Type", "application/json")
	queryParams := url.Values{}
	if len(after) != 0 {
		if after[0] != "" {
			queryParams.Add("after", after[0])
		}
	}
	req.URL.RawQuery = queryParams.Encode()
	resp, err := c.creds.Client.Do(req)
	if err != nil {
		return ownerResponse, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ownerResponse, err
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("error making http get request returned code: %d \n with body %v", resp.StatusCode, string(body))
		return ownerResponse, err
	}
	err = json.Unmarshal(body, &ownerResponse)
	if err != nil {
		return ownerResponse, err
	}
	return ownerResponse, nil
}

func (c *OwnerService) GetOwner(id int) (hubspotmodels.Owner, error) {
	owner := hubspotmodels.Owner{}
	reqUrl := fmt.Sprintf("https://api.hubapi.com/crm/v3/owners/%d", id)
	req, err := retryablehttp.NewRequest("GET", reqUrl, strings.NewReader(""))
	if err != nil {
		return owner, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.creds.AccessToken))
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.creds.Client.Do(req)
	if err != nil {
		return owner, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return owner, err
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("error making http get request returned code: %d \n with body %v", resp.StatusCode, string(body))
		return owner, err
	}
	err = json.Unmarshal(body, &owner)
	if err != nil {
		return owner, err
	}
	return owner, nil
}
