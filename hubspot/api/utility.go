package hubspot

import (
	apptypes "github.com/karman-digital/integrations/types"
)

func (c *credentials) SetAccessToken(accessToken apptypes.AccessToken) {
	c.AccessToken = accessToken
}

func (c *credentials) SetRefreshToken(refreshToken apptypes.RefreshToken) {
	c.RefreshToken = refreshToken
}

func (c *credentials) RetrieveAccessToken() apptypes.AccessToken {
	return c.AccessToken
}

func (c *credentials) RetrieveRefreshToken() apptypes.RefreshToken {
	return c.RefreshToken
}

func (c *credentials) RetrievePortalId() apptypes.AppId {
	return c.PortalId
}

func (c *credentials) SetPortalId(portalId apptypes.AppId) {
	c.PortalId = portalId
}
