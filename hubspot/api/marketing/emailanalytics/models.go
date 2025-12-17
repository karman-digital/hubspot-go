package emailanalytics

import (
	emailanalyticscampaigns "github.com/karman-digital/hubspot/hubspot/api/marketing/emailanalytics/campaigns"
	emailanalyticsevents "github.com/karman-digital/hubspot/hubspot/api/marketing/emailanalytics/events"
)

type EmailAnalytics struct {
	Campaigns *emailanalyticscampaigns.EmailAnalyticsCampaignService
	Events    *emailanalyticsevents.EmailAnalyticsEventsService
}
