package credentials

import (
	"github.com/hashicorp/go-retryablehttp"
	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
)

type Credentials struct {
	Client       *retryablehttp.Client
	AccessToken  hubspotmodels.AccessToken
	RefreshToken hubspotmodels.RefreshToken
}
