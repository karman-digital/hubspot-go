package calls

import (
	"fmt"
	"net/http"

	crmmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm"
	sharedmodels "github.com/karman-digital/hubspot/hubspot/api/models/shared"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (c *CallsService) GetCall(id string, opts ...sharedmodels.GetOptions) (crmmodels.Result, error) {
	resp, err := c.SendRequest(http.MethodGet, fmt.Sprintf("/crm/v3/objects/calls/%s", id), nil, opts...)
	if err != nil {
		return crmmodels.Result{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleResponse(resp)
}
