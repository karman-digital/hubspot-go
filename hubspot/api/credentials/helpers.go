package credentials

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/hashicorp/go-retryablehttp"
	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
)

func (c *Credentials) SendRequest(method, path string, body []byte, opts ...hubspotmodels.GetOptions) (*http.Response, error) {
	req, err := retryablehttp.NewRequest(method, "https://api.hubapi.com"+path, body)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken()))
	if len(opts) != 0 {
		queryParams := generateQueryParams(opts[0])
		req.URL.RawQuery = queryParams.Encode()
	}
	resp, err := c.Client().Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %s", err)
	}
	return resp, nil
}

func generateQueryParams(opts hubspotmodels.GetOptions) url.Values {
	queryParams := url.Values{}
	if len(opts.Properties) != 0 {
		for _, property := range opts.Properties {
			queryParams.Add("properties", property)
		}
	}
	if len(opts.PropertiesWithHistory) != 0 {
		for _, property := range opts.PropertiesWithHistory {
			queryParams.Add("propertiesWithHistory", property)
		}
	}
	if len(opts.Associations) != 0 {
		for _, association := range opts.Associations {
			queryParams.Add("associations", association)
		}
	}
	if opts.Archived {
		queryParams.Add("archived", "true")
	}
	return queryParams
}
