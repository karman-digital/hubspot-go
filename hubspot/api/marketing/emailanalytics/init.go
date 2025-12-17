package emailanalytics

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	emailanalyticscampaigns "github.com/karman-digital/hubspot/hubspot/api/marketing/emailanalytics/campaigns"
	emailanalyticsevents "github.com/karman-digital/hubspot/hubspot/api/marketing/emailanalytics/events"
)

func NewEmailAnalyticsService(creds *credentials.Credentials) EmailAnalytics {
	return EmailAnalytics{
		Campaigns: emailanalyticscampaigns.NewEmailAnalyticsCampaignService(creds),
		Events:    emailanalyticsevents.NewEmailAnalyticsEventsService(creds),
	}
}

