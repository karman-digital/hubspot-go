package credentials

import (
	"fmt"

	"github.com/hashicorp/go-retryablehttp"
	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
)

func (c *Credentials) SetAccessToken(accessToken string) error {
	var token hubspotmodels.AccessToken
	if err := token.Set(accessToken); err != nil {
		return fmt.Errorf("error setting access token: %w", err)
	}
	c.accessToken = token
	return nil
}

func (c *Credentials) SetRefreshToken(refreshToken string) error {
	var token hubspotmodels.RefreshToken
	if err := token.Set(refreshToken); err != nil {
		return fmt.Errorf("error setting refresh token: %w", err)
	}
	c.refreshToken = token
	return nil
}

func (c Credentials) AccessToken() *hubspotmodels.AccessToken {
	return &c.accessToken
}

func (c Credentials) RefreshToken() *hubspotmodels.RefreshToken {
	return &c.refreshToken
}

func (c *Credentials) SetClientId(clientId string) error {
	var id hubspotmodels.ClientId
	id.Set(clientId)
	c.clientId = id
	return nil
}

func (c *Credentials) SetClientSecret(clientSecret string) error {
	var secret hubspotmodels.ClientSecret
	secret.Set(clientSecret)
	c.clientSecret = secret
	return nil
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

func (c *Credentials) SetRedirectUri(redirectUri string) error {
	var uri hubspotmodels.RedirectUri
	uri.Set(redirectUri)
	c.redirectUri = uri
	return nil
}

func (c Credentials) RedirectUri() *hubspotmodels.RedirectUri {
	return &c.redirectUri
}
