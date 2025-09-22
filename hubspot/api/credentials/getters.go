package credentials

import (
	"github.com/hashicorp/go-retryablehttp"
	authmodels "github.com/karman-digital/hubspot/hubspot/api/models/auth"
)

func (c Credentials) AccessToken() *authmodels.AccessToken {
	return &c.accessToken
}

func (c Credentials) RefreshToken() *authmodels.RefreshToken {
	return &c.refreshToken
}

func (c Credentials) ClientId() *authmodels.ClientId {
	return &c.clientId
}

func (c Credentials) ClientSecret() *authmodels.ClientSecret {
	return &c.clientSecret
}

func (c *Credentials) SetClient(client *retryablehttp.Client) {
	c.client = client
}

func (c Credentials) Client() *retryablehttp.Client {
	return c.client
}

func (c Credentials) RedirectUri() *authmodels.RedirectUri {
	return &c.redirectUri
}
