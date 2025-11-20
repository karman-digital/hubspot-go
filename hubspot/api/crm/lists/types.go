package lists

import (
	listsmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm/lists"
	sharedmodels "github.com/karman-digital/hubspot/hubspot/api/models/shared"
	"github.com/karman-digital/hubspot/hubspot/api/crm/lists/listdata"
	"github.com/karman-digital/hubspot/hubspot/api/crm/lists/memberships"
)

type Lists struct {
	ListData    *listdata.ListDataService
	Memberships *memberships.MembershipsService
}

func (l Lists) SearchLists(body listsmodels.SearchListsBody) (listsmodels.SearchListsResponse, error) {
	return l.ListData.SearchLists(body)
}

func (l Lists) GetListMemberships(listId string, opts ...sharedmodels.GetOptions) (listsmodels.ListMembershipsResponse, error) {
	return l.Memberships.GetListMemberships(listId, opts...)
}

