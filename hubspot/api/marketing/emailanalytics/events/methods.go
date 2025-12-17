package emailanalyticsevents

import (
	"fmt"
	"net/http"
	"net/url"

	emailanalyticseventsmodels "github.com/karman-digital/hubspot/hubspot/api/models/marketing/emailanalytics/events"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (e *EmailAnalyticsEventsService) GetEmailEvents(opts ...EmailEventsOptions) (emailanalyticseventsmodels.EmailEventsResponse, error) {
	reqUrl := "/email/public/v1/events"
	var options EmailEventsOptions
	if len(opts) > 0 {
		options = opts[0]
	}
	
	queryParams := url.Values{}
	if options.AppId != 0 {
		queryParams.Add("appId", fmt.Sprintf("%d", options.AppId))
	}
	if options.CampaignId != 0 {
		queryParams.Add("campaignId", fmt.Sprintf("%d", options.CampaignId))
	}
	if options.Recipient != "" {
		queryParams.Add("recipient", options.Recipient)
	}
	if options.EventType != "" {
		queryParams.Add("eventType", options.EventType)
	}
	if options.StartTimestamp != 0 {
		queryParams.Add("startTimestamp", fmt.Sprintf("%d", options.StartTimestamp))
	}
	if options.EndTimestamp != 0 {
		queryParams.Add("endTimestamp", fmt.Sprintf("%d", options.EndTimestamp))
	}
	if options.Offset != "" {
		queryParams.Add("offset", options.Offset)
	}
	if options.Limit != 0 {
		queryParams.Add("limit", fmt.Sprintf("%d", options.Limit))
	}
	if options.ExcludeFilteredEvents {
		queryParams.Add("excludeFilteredEvents", "true")
	}
	
	if len(queryParams) > 0 {
		reqUrl += "?" + queryParams.Encode()
	}
	
	resp, err := e.SendRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return emailanalyticseventsmodels.EmailEventsResponse{}, fmt.Errorf("error creating request: %s", err)
	}
	return shared.HandleEmailEventsResponse(resp)
}

