package hubspot

import (
	"github.com/hashicorp/go-retryablehttp"
	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
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
	UpdateContact(id int, patchBody hubspotmodels.PatchBody) (hubspotmodels.ContactResponse, error)
	BatchCreateContact(body []hubspotmodels.PostBody) (hubspotmodels.BatchContactResponse, error)
	CreateContact(body hubspotmodels.PostBody) (hubspotmodels.ContactResponse, error)
	SearchContacts(body hubspotmodels.SearchBody) (hubspotmodels.SearchResponse, error)
}
