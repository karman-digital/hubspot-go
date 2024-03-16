package hubspotapp

import (
	"github.com/karman-digital/hubspot/hubspot/api/auth"
	"github.com/karman-digital/hubspot/hubspot/api/communicationpreferences"
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	"github.com/karman-digital/hubspot/hubspot/api/crm"
	"github.com/karman-digital/hubspot/hubspot/interfaces"
)

type Hubspot struct {
	credentials.Credentials
	ApiClient
}

func NewApiClient(credentials *credentials.Credentials) ApiClient {
	return ApiClient{
		Auth:                     auth.NewAuthService(credentials),
		CRM:                      crm.NewCrmService(credentials),
		CommunicationPreferences: communicationpreferences.NewCommunicationPreferencesService(credentials),
	}
}

func NewHubspotInstance(credentials *credentials.Credentials) *Hubspot {
	return &Hubspot{
		Credentials: *credentials,
		ApiClient:   NewApiClient(credentials),
	}
}

type ApiClient struct {
	interfaces.Auth
	crm.CRM
	interfaces.CommunicationPreferences
}
