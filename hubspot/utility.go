package hubspot

import (
	"fmt"

	apptypes "github.com/karman-digital/hatch-shared/types"
)

func (c *credentials) SetAccessToken(accessToken string) error {
	if accessToken == "" {
		return fmt.Errorf("access token cannot be empty")
	}
	c.AccessToken = apptypes.AccessToken(accessToken)
	return nil
}

func (c *credentials) SetRefreshToken(refreshToken string) error {
	if refreshToken == "" {
		return fmt.Errorf("refresh token cannot be empty")
	}
	c.RefreshToken = apptypes.RefreshToken(refreshToken)
	return nil
}

func (c *credentials) RetrieveAccessToken() apptypes.AccessToken {
	return c.AccessToken
}

func (c *credentials) RetrieveRefreshToken() apptypes.RefreshToken {
	return c.RefreshToken
}
