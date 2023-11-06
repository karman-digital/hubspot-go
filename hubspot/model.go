package hubspot

import (
	"github.com/hashicorp/go-retryablehttp"
	apptypes "github.com/karman-digital/hatch-shared/types"
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
	SetAccessToken(accessToken string)
	SetRefreshToken(refreshToken string)
	RefreshTokenPair(clientSecret string, clientId string, redirectUri string) error
	ValidateBearerToken() (bool, error)
}
