package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (u *UsersService) Create(body hubspotmodels.UserBody) (hubspotmodels.UserBody, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return hubspotmodels.UserBody{}, fmt.Errorf("error marshalling post body: %s", err)
	}
	resp, err := u.SendRequest(http.MethodPost, "/settings/v3/users", reqBody)
	if err != nil {
		return hubspotmodels.UserBody{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleUserResponse(resp)
}
