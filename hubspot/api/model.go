package hubspot

import (
	"github.com/hashicorp/go-retryablehttp"
	husbpotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
	apptypes "github.com/karman-digital/integrations/types"
)

type credentials struct {
	Client       *retryablehttp.Client
	AccessToken  apptypes.AccessToken
	RefreshToken apptypes.RefreshToken
	PortalId     apptypes.PortalId
}

type HubspotAPI interface {
	RetrieveAccessToken() apptypes.AccessToken
	RetrieveRefreshToken() apptypes.RefreshToken
	RetrievePortalId() apptypes.PortalId
	SetAccessToken(accessToken apptypes.AccessToken)
	SetRefreshToken(refreshToken apptypes.RefreshToken)
	SetPortalId(portalId apptypes.PortalId)
	RefreshTokenPair(clientSecret string, clientId string, redirectUri string) error
	ValidateBearerToken() (bool, error)
	UpdateContact(id int, patchBody husbpotmodels.PatchBody) (husbpotmodels.ContactResponse, error)
	BatchCreateContact(body []husbpotmodels.PostBody) (husbpotmodels.BatchContactResponse, error)
	CreateContact(body husbpotmodels.PostBody) (husbpotmodels.ContactResponse, error)
}
