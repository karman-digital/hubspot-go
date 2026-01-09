package owners

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/hashicorp/go-retryablehttp"
	ownersmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm/owners"
)

func (c *OwnerService) GetAllOwners() ([]ownersmodels.Owner, error) {
	var allOwners []ownersmodels.Owner
	after := ""
	for {
		opts := ownersmodels.GetOwnersOptions{}
		if after != "" {
			opts.After = after
		}
		ownerResponse, err := c.GetOwners(opts)
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

func (c *OwnerService) GetOwners(opts ...ownersmodels.GetOwnersOptions) (ownersmodels.OwnerResponse, error) {
	ownerResponse := ownersmodels.OwnerResponse{}
	reqUrl := "https://api.hubapi.com/crm/v3/owners"
	req, err := retryablehttp.NewRequest("GET", reqUrl, strings.NewReader(""))
	if err != nil {
		return ownerResponse, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken()))
	req.Header.Set("Content-Type", "application/json")
	queryParams := url.Values{}
	if len(opts) != 0 {
		if opts[0].After != "" {
			queryParams.Add("after", opts[0].After)
		}
		if opts[0].Archived {
			queryParams.Add("archived", "true")
		}
		if opts[0].Email != "" {
			queryParams.Add("email", opts[0].Email)
		}
		if opts[0].Limit != 0 {
			queryParams.Add("limit", fmt.Sprintf("%d", opts[0].Limit))
		}
	}
	req.URL.RawQuery = queryParams.Encode()
	resp, err := c.Client().Do(req)
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

func (c *OwnerService) GetOwner(id int) (ownersmodels.Owner, error) {
	owner := ownersmodels.Owner{}
	reqUrl := fmt.Sprintf("https://api.hubapi.com/crm/v3/owners/%d", id)
	req, err := retryablehttp.NewRequest("GET", reqUrl, strings.NewReader(""))
	if err != nil {
		return owner, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken()))
	resp, err := c.Client().Do(req)
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
