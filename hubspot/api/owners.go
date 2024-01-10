package hubspot

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

func (c *credentials) GetOwners(after ...string) (hubspotmodels.OwnerResponse, error) {
	ownerResponse := hubspotmodels.OwnerResponse{}
	reqUrl := "https://api.hubapi.com/crm/v3/owners"
	req, err := retryablehttp.NewRequest("GET", reqUrl, strings.NewReader(""))
	if err != nil {
		return ownerResponse, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
	req.Header.Set("Content-Type", "application/json")
	queryParams := url.Values{}
	if len(after) != 0 {
		queryParams.Add("after", after[0])
	}
	req.URL.RawQuery = queryParams.Encode()
	client := retryablehttp.NewClient()
	resp, err := client.Do(req)
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
