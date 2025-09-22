package credentials

import (
	"github.com/hashicorp/go-retryablehttp"
	authmodels "github.com/karman-digital/hubspot/hubspot/api/models/auth"
)

type Credentials struct {
	client       *retryablehttp.Client
	accessToken  authmodels.AccessToken
	refreshToken authmodels.RefreshToken
	clientId     authmodels.ClientId
	clientSecret authmodels.ClientSecret
	redirectUri  authmodels.RedirectUri
}
