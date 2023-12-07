package hubspot

import (
	"github.com/hashicorp/go-retryablehttp"
	"github.com/karman-digital/hubspot/hubspot/api/models"
	apptypes "github.com/karman-digital/integrations/types"
)

type credentials struct {
	Client       *retryablehttp.Client
	AccessToken  apptypes.AccessToken
	RefreshToken apptypes.RefreshToken
	PortalId     apptypes.AppId
}

type HubspotAPI interface {
	RetrieveAccessToken() apptypes.AccessToken
	RetrieveRefreshToken() apptypes.RefreshToken
	RetrievePortalId() apptypes.AppId
	SetAccessToken(accessToken apptypes.AccessToken)
	SetRefreshToken(refreshToken apptypes.RefreshToken)
	SetPortalId(portalId apptypes.AppId)
	RefreshTokenPair(clientSecret string, clientId string, redirectUri string) error
	ValidateBearerToken() (bool, error)
	UpdateContact(id int, patchBody models.PatchBody) (models.ContactResponse, error)
}
