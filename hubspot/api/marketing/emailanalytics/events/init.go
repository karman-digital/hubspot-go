package emailanalyticsevents

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
)

func NewEmailAnalyticsEventsService(creds *credentials.Credentials) *EmailAnalyticsEventsService {
	return &EmailAnalyticsEventsService{
		creds,
	}
}

