package hubspot

import "fmt"

func (c *credentials) SetAccessToken(accessToken string) error {
	if accessToken == "" {
		return fmt.Errorf("access token cannot be empty")
	}
	c.AccessToken = AccessToken(accessToken)
	return nil
}

func (c *credentials) SetRefreshToken(refreshToken string) error {
	if refreshToken == "" {
		return fmt.Errorf("refresh token cannot be empty")
	}
	c.RefreshToken = RefreshToken(refreshToken)
	return nil
}

func (c *credentials) RetrieveAccessToken() AccessToken {
	return c.AccessToken
}

func (c *credentials) RetrieveRefreshToken() RefreshToken {
	return c.RefreshToken
}
