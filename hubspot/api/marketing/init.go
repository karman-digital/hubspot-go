package marketing

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	"github.com/karman-digital/hubspot/hubspot/api/marketing/campaigns"
	"github.com/karman-digital/hubspot/hubspot/api/marketing/emailanalytics"
	"github.com/karman-digital/hubspot/hubspot/api/marketing/emails"
)

func NewMarketingService(creds *credentials.Credentials) Marketing {
	return Marketing{
		Campaigns:      campaigns.NewCampaignService(creds),
		Emails:         emails.NewMarketingEmailService(creds),
		EmailAnalytics: emailanalytics.NewEmailAnalyticsService(creds),
	}
}

