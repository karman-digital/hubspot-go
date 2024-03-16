package hubspotapp

import (
	"fmt"

	"github.com/karman-digital/hubspot/hubspot/api/auth"
	"github.com/karman-digital/hubspot/hubspot/api/communicationpreferences"
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	"github.com/karman-digital/hubspot/hubspot/api/crm"
	"github.com/karman-digital/hubspot/hubspot/interfaces"
)

type Hubspot struct {
	credentials.Credentials
	ApiClient
	PortalId
}

func InitHubspot() *Hubspot {
	return &Hubspot{}
}

func (h *Hubspot) InitClient(credentials *credentials.Credentials) {
	h.Credentials = *credentials
	h.ApiClient = NewApiClient(credentials)
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

func (c Hubspot) RetrievePortalId() PortalId {
	return c.PortalId
}

func (c *Hubspot) SetPortalId(portalId int) error {
	var id PortalId
	if err := id.Set(portalId); err != nil {
		return fmt.Errorf("error setting portal id: %w", err)
	}
	c.PortalId = id
	return nil
}
