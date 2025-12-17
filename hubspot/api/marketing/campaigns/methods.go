package campaigns

import (
	"fmt"
	"net/http"

	campaignassetsmodels "github.com/karman-digital/hubspot/hubspot/api/models/marketing/campaigns/assets"
	campaignmodels "github.com/karman-digital/hubspot/hubspot/api/models/marketing/campaigns"
	sharedmodels "github.com/karman-digital/hubspot/hubspot/api/models/shared"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (c *CampaignService) GetCampaigns(opts ...sharedmodels.GetOptions) (campaignmodels.CampaignsResponse, error) {
	reqUrl := "/marketing/v3/campaigns"
	resp, err := c.SendRequest(http.MethodGet, reqUrl, nil, opts...)
	if err != nil {
		return campaignmodels.CampaignsResponse{}, fmt.Errorf("error creating request: %s", err)
	}
	return shared.HandleCampaignsResponse(resp)
}

func (c *CampaignService) GetCampaignAssets(campaignGuid string, assetType string, opts ...sharedmodels.GetOptions) (campaignassetsmodels.CampaignAssetsResponse, error) {
	reqUrl := fmt.Sprintf("/marketing/v3/campaigns/%s/assets/%s", campaignGuid, assetType)
	resp, err := c.SendRequest(http.MethodGet, reqUrl, nil, opts...)
	if err != nil {
		return campaignassetsmodels.CampaignAssetsResponse{}, fmt.Errorf("error creating request: %s", err)
	}
	return shared.HandleCampaignAssetsResponse(resp)
}

