package hubspot

import (
	"fmt"

	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
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

func (c *credentials) RetrievePortalId() apptypes.PortalId {
	return c.PortalId
}

func (c *credentials) SetPortalId(portalId apptypes.PortalId) {
	c.PortalId = portalId
}

func handleBatchResponseCodes(errorResp hubspotmodels.ErrorResponseBody, statusCode int) error {
	switch statusCode {
	case 200:
		return nil
	case 400:
		if errorResp.Category == "VALIDATION" {
			fmt.Printf("validation error returned by endpoint: %s", errorResp.Message)
			return ErrPropertyValidation
		}
		return fmt.Errorf("error returned by endpoint: %s", errorResp.Message)
	case 409:
		fmt.Printf("object already exists: %s", errorResp.Message)
		return ErrObjectAlreadyExists
	default:
		return ErrBatchCreate
	}
}
