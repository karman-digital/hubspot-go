package hubspot

import (
	"github.com/hashicorp/go-retryablehttp"
)

func InitHubspotAPI() HubspotAPI {
	client := retryablehttp.NewClient()
	client.Logger = nil
	return &credentials{
		Client: client,
	}
}

type HubspotInitFunc func() HubspotAPI
