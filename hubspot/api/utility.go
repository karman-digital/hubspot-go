package hubspot

import (
	"fmt"

	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
)

func (c *credentials) SetAccessToken(accessToken string) error {
	var token AccessToken
	if err := token.Set(accessToken); err != nil {
		return fmt.Errorf("error setting access token: %w", err)
	}
	c.AccessToken = token
	return nil
}

func (c *credentials) SetRefreshToken(refreshToken string) error {
	var token RefreshToken
	if err := token.Set(refreshToken); err != nil {
		return fmt.Errorf("error setting refresh token: %w", err)
	}
	c.RefreshToken = token
	return nil
}

func (c *credentials) RetrieveAccessToken() AccessToken {
	return c.AccessToken
}

func (c *credentials) RetrieveRefreshToken() RefreshToken {
	return c.RefreshToken
}

func (c *credentials) RetrievePortalId() PortalId {
	return c.PortalId
}

func (c *credentials) SetPortalId(portalId int) error {
	var id PortalId
	if err := id.Set(portalId); err != nil {
		return fmt.Errorf("error setting portal id: %w", err)
	}
	c.PortalId = id
	return nil
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
		return ErrApiCall
	}
}
