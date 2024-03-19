package credentials

import (
	"github.com/hashicorp/go-retryablehttp"
	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
)

func (c Credentials) AccessToken() *hubspotmodels.AccessToken {
	return &c.accessToken
}

func (c Credentials) RefreshToken() *hubspotmodels.RefreshToken {
	return &c.refreshToken
}

func (c Credentials) ClientId() *hubspotmodels.ClientId {
	return &c.clientId
}

func (c Credentials) ClientSecret() *hubspotmodels.ClientSecret {
	return &c.clientSecret
}

func (c *Credentials) SetClient(client *retryablehttp.Client) {
	c.client = client
}

func (c Credentials) Client() *retryablehttp.Client {
	return c.client
}

func (c Credentials) RedirectUri() *hubspotmodels.RedirectUri {
	return &c.redirectUri
}
