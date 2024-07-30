package crm

import (
	"github.com/karman-digital/hubspot/hubspot/api/crm/engagements"
	"github.com/karman-digital/hubspot/hubspot/interfaces"
)

type CRM struct {
	Contacts      interfaces.Contact
	Owners        interfaces.Owners
	Properties    interfaces.Properties
	Associations  interfaces.Associations
	Deals         interfaces.Deal
	CustomObjects interfaces.CustomObject
	Engagements   engagements.Engagements
	Companies     interfaces.Company
	Products      interfaces.Products
	LineItems     interfaces.LineItems
}
