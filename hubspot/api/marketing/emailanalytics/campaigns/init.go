package emailanalyticscampaigns

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
)

func NewEmailAnalyticsCampaignService(creds *credentials.Credentials) *EmailAnalyticsCampaignService {
	return &EmailAnalyticsCampaignService{
		creds,
	}
}

