package crm

import "github.com/karman-digital/hubspot/hubspot/interfaces"

type CRM struct {
	Contacts     interfaces.Contact
	Owners       interfaces.Owners
	Properties   interfaces.Properties
	Associations interfaces.Associations
}
