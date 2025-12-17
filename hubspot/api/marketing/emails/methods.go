package emails

import (
	"fmt"
	"net/http"

	emailmodels "github.com/karman-digital/hubspot/hubspot/api/models/marketing/emails"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (e *MarketingEmailService) GetMarketingEmail(emailId string) (emailmodels.MarketingEmail, error) {
	reqUrl := fmt.Sprintf("/marketing/v3/emails/%s", emailId)
	resp, err := e.SendRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return emailmodels.MarketingEmail{}, fmt.Errorf("error creating request: %s", err)
	}
	return shared.HandleMarketingEmailResponse(resp)
}

