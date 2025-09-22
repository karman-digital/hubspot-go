package credentials

import (
	"fmt"

	authmodels "github.com/karman-digital/hubspot/hubspot/api/models/auth"
)

func (c *Credentials) SetRedirectUri(redirectUri string) error {
	var uri authmodels.RedirectUri
	uri.Set(redirectUri)
	c.redirectUri = uri
	return nil
}

func (c *Credentials) SetClientId(clientId string) error {
	var id authmodels.ClientId
	id.Set(clientId)
	c.clientId = id
	return nil
}

func (c *Credentials) SetClientSecret(clientSecret string) error {
	var secret authmodels.ClientSecret
	secret.Set(clientSecret)
	c.clientSecret = secret
	return nil
}

func (c *Credentials) SetAccessToken(accessToken string) error {
	var token authmodels.AccessToken
	if err := token.Set(accessToken); err != nil {
		return fmt.Errorf("error setting access token: %w", err)
	}
	c.accessToken = token
	return nil
}

func (c *Credentials) SetRefreshToken(refreshToken string) error {
	var token authmodels.RefreshToken
	if err := token.Set(refreshToken); err != nil {
		return fmt.Errorf("error setting refresh token: %w", err)
	}
	c.refreshToken = token
	return nil
}
