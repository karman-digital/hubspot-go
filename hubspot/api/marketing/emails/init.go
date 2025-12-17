package emails

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
)

func NewMarketingEmailService(creds *credentials.Credentials) *MarketingEmailService {
	return &MarketingEmailService{
		creds,
	}
}

