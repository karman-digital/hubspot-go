package hubspot

import "github.com/hashicorp/go-retryablehttp"

type AccessToken string
type RefreshToken string

type credentials struct {
	Client       *retryablehttp.Client
	AccessToken  AccessToken
	RefreshToken RefreshToken
}

type HubspotAPI interface {
	RetrieveAccessToken() AccessToken
	RetrieveRefreshToken() RefreshToken
	SetAccessToken(accessToken string)
	SetRefreshToken(refreshToken string)
	RefreshTokenPair(clientSecret string, clientId string, redirectUri string) error
	ValidateBearerToken() (bool, error)
}

func (r RefreshToken) String() string {
	return string(r)
}

func (a AccessToken) String() string {
	return string(a)
}
