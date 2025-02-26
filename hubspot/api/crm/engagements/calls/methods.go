package calls

import (
	"fmt"
	"net/http"

	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (c *CallsService) GetCall(id string, opts ...hubspotmodels.GetOptions) (hubspotmodels.Result, error) {
	resp, err := c.SendRequest(http.MethodGet, fmt.Sprintf("/crm/v3/objects/calls/%s", id), nil, opts...)
	if err != nil {
		return hubspotmodels.Result{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleResponse(resp)
}
