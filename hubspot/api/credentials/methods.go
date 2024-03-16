package credentials

import "fmt"

func (c *Credentials) SetAccessToken(accessToken string) error {
	var token AccessToken
	if err := token.Set(accessToken); err != nil {
		return fmt.Errorf("error setting access token: %w", err)
	}
	c.AccessToken = token
	return nil
}

func (c *Credentials) SetRefreshToken(refreshToken string) error {
	var token RefreshToken
	if err := token.Set(refreshToken); err != nil {
		return fmt.Errorf("error setting refresh token: %w", err)
	}
	c.RefreshToken = token
	return nil
}

func (c *Credentials) RetrieveAccessToken() *AccessToken {
	return &c.AccessToken
}

func (c *Credentials) RetrieveRefreshToken() *RefreshToken {
	return &c.RefreshToken
}

func (c *Credentials) RetrievePortalId() *PortalId {
	return &c.PortalId
}

func (c *Credentials) SetPortalId(portalId int) error {
	var id PortalId
	if err := id.Set(portalId); err != nil {
		return fmt.Errorf("error setting portal id: %w", err)
	}
	c.PortalId = id
	return nil
}
