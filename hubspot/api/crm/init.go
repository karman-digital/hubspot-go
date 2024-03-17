package crm

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	"github.com/karman-digital/hubspot/hubspot/api/crm/associations"
	"github.com/karman-digital/hubspot/hubspot/api/crm/contact"
	"github.com/karman-digital/hubspot/hubspot/api/crm/owners"
	"github.com/karman-digital/hubspot/hubspot/api/crm/properties"
)

func NewCrmService(creds *credentials.Credentials) CRM {
	return CRM{
		Contacts:     contact.NewContactService(creds),
		Owners:       owners.NewOwnerService(creds),
		Associations: associations.NewAssociationService(creds),
		Properties:   properties.NewPropertiesService(creds),
	}
}
