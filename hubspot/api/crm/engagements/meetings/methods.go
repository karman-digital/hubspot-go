package meetings

import (
	"fmt"
	"net/http"

	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (e *MeetingsService) GetMeeting(id string, opts ...hubspotmodels.GetOptions) (hubspotmodels.Result, error) {
	resp, err := e.SendRequest(http.MethodGet, fmt.Sprintf("/crm/v3/objects/meetings/%s", id), nil, opts...)
	if err != nil {
		return hubspotmodels.Result{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleResponse(resp)
}
