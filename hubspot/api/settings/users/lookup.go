package users

import (
	"fmt"
	usermodels "github.com/karman-digital/hubspot/hubspot/api/models/users"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
	"net/http"
	"net/url"
)

func (u *UsersService) GetByEmail(email string) (usermodels.UserBody, bool, error) {
	response, err := u.SendRequest(http.MethodGet, "/settings/users/2026-09/"+url.PathEscape(email)+"?idProperty=EMAIL", nil)
	if err != nil {
		return usermodels.UserBody{}, false, err
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusNotFound {
		return usermodels.UserBody{}, false, nil
	}
	user, err := shared.HandleUserResponse(response)
	if err != nil {
		return user, false, err
	}
	if user.Id == "" || user.Email != email {
		return user, false, fmt.Errorf("HubSpot user identity mismatch")
	}
	return user, true, nil
}
