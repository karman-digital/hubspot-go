package campaignassetsmodels

import (
	sharedmodels "github.com/karman-digital/hubspot/hubspot/api/models/shared"
)

type CampaignAsset struct {
	Id      string         `json:"id"`
	Name    string         `json:"name"`
	Metrics map[string]int `json:"metrics,omitempty"`
}

type CampaignAssetsResponse struct {
	Results []CampaignAsset      `json:"results"`
	Paging  sharedmodels.Paging  `json:"paging,omitempty"`
}

