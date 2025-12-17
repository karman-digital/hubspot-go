package campaigns

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
)

func NewCampaignService(creds *credentials.Credentials) *CampaignService {
	return &CampaignService{
		creds,
	}
}

