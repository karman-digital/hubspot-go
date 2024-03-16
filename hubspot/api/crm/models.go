package crm

import "github.com/karman-digital/hubspot/hubspot/interfaces"

type CRM struct {
	ContactService     interfaces.Contact
	OwnerService       interfaces.Owners
	PropertiesService  interfaces.Properties
	AssociationService interfaces.Associations
}
