package hubspot

import (
	"github.com/hashicorp/go-retryablehttp"
	apptypes "github.com/karman-digital/hatch-shared/types"
	hubspotmodels "github.com/karman-digital/hubspot/hubspot/models"
)

type AccessToken string
type RefreshToken string

type credentials struct {
	Client       *retryablehttp.Client
	AccessToken  apptypes.AccessToken
	RefreshToken apptypes.RefreshToken
}

type HubspotAPI interface {
	RetrieveAccessToken() apptypes.AccessToken
	RetrieveRefreshToken() apptypes.RefreshToken
	SetAccessToken(accessToken string) error
	SetRefreshToken(refreshToken string) error
	RefreshTokenPair(clientSecret string, clientId string, redirectUri string) error
	ValidateBearerToken() (bool, error)
	UpdateContact(id int, patchBody hubspotmodels.PatchBody) (hubspotmodels.ContactResponse, error)
}
