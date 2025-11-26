package memberships

import (
	"fmt"
	"net/http"
	"net/url"

	listsmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm/lists"
	sharedmodels "github.com/karman-digital/hubspot/hubspot/api/models/shared"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (m *MembershipsService) GetListMemberships(listId string, opts ...sharedmodels.GetOptions) (listsmodels.ListMembershipsResponse, error) {
	reqUrl := fmt.Sprintf("/crm/v3/lists/%s/memberships", listId)
	
	if len(opts) != 0 {
		queryParams := url.Values{}
		if opts[0].After != "" {
			queryParams.Add("after", opts[0].After)
		}
		if opts[0].Limit != 0 {
			queryParams.Add("limit", fmt.Sprintf("%d", opts[0].Limit))
		}
		if encoded := queryParams.Encode(); encoded != "" {
			reqUrl = fmt.Sprintf("%s?%s", reqUrl, encoded)
		}
	}
	
	resp, err := m.SendRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return listsmodels.ListMembershipsResponse{}, fmt.Errorf("error creating request: %s", err)
	}
	return shared.HandleListMembershipsResponse(resp)
}
