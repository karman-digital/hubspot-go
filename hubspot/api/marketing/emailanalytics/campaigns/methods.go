package emailanalyticscampaigns

import (
	"fmt"
	"net/http"
	"net/url"

	emailanalyticscampaignmodels "github.com/karman-digital/hubspot/hubspot/api/models/marketing/emailanalytics/campaigns"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (e *EmailAnalyticsCampaignService) GetEmailCampaigns(offset ...string) (emailanalyticscampaignmodels.EmailCampaignsResponse, error) {
	reqUrl := "/email/public/v1/campaigns"
	if len(offset) > 0 && offset[0] != "" {
		reqUrl += "?offset=" + url.QueryEscape(offset[0])
	}
	resp, err := e.SendRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return emailanalyticscampaignmodels.EmailCampaignsResponse{}, fmt.Errorf("error creating request: %s", err)
	}
	return shared.HandleEmailCampaignsResponse(resp)
}

func (e *EmailAnalyticsCampaignService) GetEmailCampaign(campaignId int) (emailanalyticscampaignmodels.EmailCampaignDetail, error) {
	reqUrl := fmt.Sprintf("/email/public/v1/campaigns/%d", campaignId)
	resp, err := e.SendRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return emailanalyticscampaignmodels.EmailCampaignDetail{}, fmt.Errorf("error creating request: %s", err)
	}
	return shared.HandleEmailCampaignDetailResponse(resp)
}

