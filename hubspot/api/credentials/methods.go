package credentials

import (
	"fmt"

	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
)

func (c *Credentials) SetAccessToken(accessToken string) error {
	var token hubspotmodels.AccessToken
	if err := token.Set(accessToken); err != nil {
		return fmt.Errorf("error setting access token: %w", err)
	}
	c.AccessToken = token
	return nil
}

func (c *Credentials) SetRefreshToken(refreshToken string) error {
	var token hubspotmodels.RefreshToken
	if err := token.Set(refreshToken); err != nil {
		return fmt.Errorf("error setting refresh token: %w", err)
	}
	c.RefreshToken = token
	return nil
}

func (c *Credentials) RetrieveAccessToken() *hubspotmodels.AccessToken {
	return &c.AccessToken
}

func (c *Credentials) RetrieveRefreshToken() *hubspotmodels.RefreshToken {
	return &c.RefreshToken
}
