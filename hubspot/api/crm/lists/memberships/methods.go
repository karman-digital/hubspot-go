package memberships

import (
	"fmt"
	"net/http"

	listsmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm/lists"
	sharedmodels "github.com/karman-digital/hubspot/hubspot/api/models/shared"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (m *MembershipsService) GetListMemberships(listId string, opts ...sharedmodels.GetOptions) (listsmodels.ListMembershipsResponse, error) {
	reqUrl := fmt.Sprintf("/crm/v3/lists/%s/memberships", listId)
	resp, err := m.SendRequest(http.MethodGet, reqUrl, nil, opts...)
	if err != nil {
		return listsmodels.ListMembershipsResponse{}, fmt.Errorf("error creating request: %s", err)
	}
	return shared.HandleListMembershipsResponse(resp)
}
