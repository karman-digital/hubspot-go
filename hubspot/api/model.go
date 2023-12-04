package hubspot

import (
	"github.com/hashicorp/go-retryablehttp"
	apptypes "github.com/karman-digital/hatch-shared/types"
	"github.com/karman-digital/hubspot/hubspot/api/models"
)

type credentials struct {
	Client       *retryablehttp.Client
	AccessToken  apptypes.AccessToken
	RefreshToken apptypes.RefreshToken
}

type HubspotAPI interface {
	RetrieveAccessToken() apptypes.AccessToken
	RetrieveRefreshToken() apptypes.RefreshToken
	SetAccessToken(accessToken apptypes.AccessToken)
	SetRefreshToken(refreshToken apptypes.RefreshToken)
	RefreshTokenPair(clientSecret string, clientId string, redirectUri string) error
	ValidateBearerToken() (bool, error)
	UpdateContact(id int, patchBody models.PatchBody) (models.ContactResponse, error)
}
