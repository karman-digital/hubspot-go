package hubspotapp

import (
	"fmt"

	"github.com/karman-digital/hubspot/hubspot/api/cms"
	"github.com/karman-digital/hubspot/hubspot/api/communicationpreferences"
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	"github.com/karman-digital/hubspot/hubspot/api/crm"
	"github.com/karman-digital/hubspot/hubspot/api/filesystem"
	"github.com/karman-digital/hubspot/hubspot/api/marketing"
	"github.com/karman-digital/hubspot/hubspot/api/settings"
)

func InitHubspot() *Hubspot {
	return &Hubspot{}
}

func (h *Hubspot) InitClient(credentials *credentials.Credentials) {
	h.Credentials = credentials
	h.ApiClient = NewApiClient(credentials)
}

func NewApiClient(credentials *credentials.Credentials) ApiClient {
	return ApiClient{
		CRM:                      crm.NewCrmService(credentials),
		CMS:                      cms.NewCmsService(credentials),
		CommunicationPreferences: communicationpreferences.NewCommunicationPreferencesService(credentials),
		FileSystem:               filesystem.NewFilesystemService(credentials),
		Settings:                 settings.NewSettingsService(credentials),
		Marketing:                marketing.NewMarketingService(credentials),
	}
}

func NewHubspotInstance(credentials *credentials.Credentials) *Hubspot {
	return &Hubspot{
		Credentials: credentials,
		ApiClient:   NewApiClient(credentials),
	}
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
