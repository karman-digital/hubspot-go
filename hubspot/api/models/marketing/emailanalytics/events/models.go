package emailanalyticseventsmodels

type EmailEvent struct {
	Id              int64               `json:"id"`
	AppId           int64               `json:"appId"`
	AppName         string              `json:"appName"`
	Created         int64               `json:"created"`
	EmailCampaignId int64               `json:"emailCampaignId"`
	PortalId        int64               `json:"portalId"`
	Recipient       string              `json:"recipient"`
	Type            string              `json:"type"`
	Browser         *EmailEventBrowser  `json:"browser,omitempty"`
	Hmid            string              `json:"hmid,omitempty"`
	Location        *EmailEventLocation `json:"location,omitempty"`
	Response        string              `json:"response,omitempty"`
	SentBy          *EmailEventSentBy   `json:"sentBy,omitempty"`
	SendId          int64               `json:"sendId,omitempty"`
	SmtpId          string              `json:"smtpId,omitempty"`
	Subject         string              `json:"subject,omitempty"`
	UserAgent       string              `json:"userAgent,omitempty"`
}

type EmailEventBrowser struct {
	Family      string `json:"family"`
	Name        string `json:"name"`
	Producer    string `json:"producer"`
	ProducerUrl string `json:"producerUrl"`
	Type        string `json:"type"`
	Url         string `json:"url"`
	Version     string `json:"version"`
}

type EmailEventLocation struct {
	City      string  `json:"city"`
	Country   string  `json:"country"`
	State     string  `json:"state"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type EmailEventSentBy struct {
	Created int64 `json:"created"`
	Id      int64 `json:"id"`
}

type EmailEventsResponse struct {
	Events  []EmailEvent `json:"events"`
	HasMore bool         `json:"hasMore"`
	Offset  string       `json:"offset"`
}
