package emailanalyticsevents

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
)

type EmailAnalyticsEventsService struct {
	*credentials.Credentials
}

type EmailEventsOptions struct {
	AppId                int64  `url:"appId,omitempty"`
	CampaignId           int64  `url:"campaignId,omitempty"`
	Recipient            string `url:"recipient,omitempty"`
	EventType            string `url:"eventType,omitempty"`
	StartTimestamp       int64  `url:"startTimestamp,omitempty"`
	EndTimestamp         int64  `url:"endTimestamp,omitempty"`
	Offset               string `url:"offset,omitempty"`
	Limit                int    `url:"limit,omitempty"`
	ExcludeFilteredEvents bool  `url:"excludeFilteredEvents,omitempty"`
}

