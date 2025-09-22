package meetings

import (
	"fmt"
	"net/http"

	crmmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm"
	sharedmodels "github.com/karman-digital/hubspot/hubspot/api/models/shared"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (e *MeetingsService) GetMeeting(id string, opts ...sharedmodels.GetOptions) (crmmodels.Result, error) {
	resp, err := e.SendRequest(http.MethodGet, fmt.Sprintf("/crm/v3/objects/meetings/%s", id), nil, opts...)
	if err != nil {
		return crmmodels.Result{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleResponse(resp)
}
