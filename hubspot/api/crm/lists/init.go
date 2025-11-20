package lists

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	listdata "github.com/karman-digital/hubspot/hubspot/api/crm/lists/listdata"
	"github.com/karman-digital/hubspot/hubspot/api/crm/lists/memberships"
)

func NewListsService(creds *credentials.Credentials) Lists {
	return Lists{
		ListData:    listdata.NewListDataService(creds),
		Memberships: memberships.NewMembershipsService(creds),
	}
}

