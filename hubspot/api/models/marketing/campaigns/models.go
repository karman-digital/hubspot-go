package campaignmodels

import (
	"time"

	sharedmodels "github.com/karman-digital/hubspot/hubspot/api/models/shared"
)

type Campaign struct {
	Id            string                 `json:"id"`
	BusinessUnits []BusinessUnit         `json:"businessUnits"`
	CreatedAt     time.Time              `json:"createdAt"`
	UpdatedAt     time.Time              `json:"updatedAt"`
	Properties    map[string]interface{} `json:"properties"`
}

type BusinessUnit struct {
	Id int `json:"id"`
}

type CampaignsResponse struct {
	Results []Campaign          `json:"results"`
	Total   int                 `json:"total"`
	Paging  sharedmodels.Paging `json:"paging"`
}
