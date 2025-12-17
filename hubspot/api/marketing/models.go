package marketing

import (
	"github.com/karman-digital/hubspot/hubspot/api/marketing/campaigns"
	"github.com/karman-digital/hubspot/hubspot/api/marketing/emailanalytics"
	"github.com/karman-digital/hubspot/hubspot/api/marketing/emails"
)

type Marketing struct {
	Campaigns     *campaigns.CampaignService
	Emails        *emails.MarketingEmailService
	EmailAnalytics emailanalytics.EmailAnalytics
}

