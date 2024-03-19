package credentials

import (
	"github.com/hashicorp/go-retryablehttp"
	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
)

type Credentials struct {
	client       *retryablehttp.Client
	accessToken  hubspotmodels.AccessToken
	refreshToken hubspotmodels.RefreshToken
	clientId     hubspotmodels.ClientId
	clientSecret hubspotmodels.ClientSecret
	redirectUri  hubspotmodels.RedirectUri
}
