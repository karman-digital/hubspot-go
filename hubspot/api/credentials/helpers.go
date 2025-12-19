package credentials

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/hashicorp/go-retryablehttp"
	sharedmodels "github.com/karman-digital/hubspot/hubspot/api/models/shared"
)

func (c *Credentials) SendRequest(method, path string, body []byte, opts ...sharedmodels.GetOptions) (*http.Response, error) {
	fullURL := "https://api.hubapi.com" + path
	req, err := retryablehttp.NewRequest(method, fullURL, body)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken()))
	
	if len(opts) != 0 {
		queryParams := generateQueryParams(opts[0])
		req.URL.RawQuery = queryParams.Encode()
	}
	
	// Log the exact request being made
	fmt.Printf("SendRequest - Method: %s\n", method)
	fmt.Printf("SendRequest - Full URL: %s\n", req.URL.String())
	fmt.Printf("SendRequest - Headers:\n")
	for name, values := range req.Header {
		if name == "Authorization" {
			token := c.AccessToken().String()
			if len(token) > 20 {
				fmt.Printf("  %s: Bearer %s...\n", name, token[:20])
			} else {
				fmt.Printf("  %s: Bearer %s\n", name, token)
			}
		} else {
			fmt.Printf("  %s: %v\n", name, values)
		}
	}
	if req.URL.RawQuery != "" {
		fmt.Printf("SendRequest - Query Params: %s\n", req.URL.RawQuery)
	} else {
		fmt.Printf("SendRequest - Query Params: (none)\n")
	}
	
	resp, err := c.Client().Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %s", err)
	}
	return resp, nil
}

func generateQueryParams(opts sharedmodels.GetOptions) url.Values {
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
	if opts.After != "" {
		queryParams.Add("after", opts.After)
	}
	if opts.Limit != 0 {
		queryParams.Add("limit", fmt.Sprintf("%d", opts.Limit))
	}
	if opts.Archived {
		queryParams.Add("archived", "true")
	}
	if opts.IdProperty != "" {
		queryParams.Add("idProperty", opts.IdProperty)
	}
	return queryParams
}
