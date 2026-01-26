package emailanalyticseventsmodels

type EmailEvent struct {
	Id                  string               `json:"id"`
	AppId               int64                `json:"appId"`
	AppName             string               `json:"appName"`
	Created             int64                `json:"created"`
	EmailCampaignId     int64                `json:"emailCampaignId"`
	PortalId            int64                `json:"portalId"`
	Recipient           string               `json:"recipient"`
	Type                string               `json:"type"`
	Browser             *EmailEventBrowser   `json:"browser,omitempty"`
	Hmid                string               `json:"hmid,omitempty"`
	Location            *EmailEventLocation  `json:"location,omitempty"`
	Response            string               `json:"response,omitempty"`
	SentBy              *EmailEventSentBy    `json:"sentBy,omitempty"`
	SendId              string               `json:"sendId,omitempty"`
	SmtpId              *string              `json:"smtpId,omitempty"`
	Subject             string               `json:"subject,omitempty"`
	UserAgent           string               `json:"userAgent,omitempty"`
	LinkId              int64                `json:"linkId,omitempty"`
	DeviceType          string               `json:"deviceType,omitempty"`
	LinkIdV2            string               `json:"linkIdV2,omitempty"`
	Url                 string               `json:"url,omitempty"`
	Referer             string               `json:"referer,omitempty"`
	FilteredEvent       bool                 `json:"filteredEvent,omitempty"`
	Status                  string                      `json:"status,omitempty"`
	Category                string                      `json:"category,omitempty"`
	Attempt                 int                         `json:"attempt,omitempty"`
	FranklinResponseType     string                      `json:"franklinResponseType,omitempty"`
	Subscriptions            []EmailEventSubscription    `json:"subscriptions,omitempty"`
	PortalSubscriptionStatus string                      `json:"portalSubscriptionStatus,omitempty"`
	Source                  string                      `json:"source,omitempty"`
	SuppressedReason        string                      `json:"suppressedReason,omitempty"`
	SuppressedMessage       string                      `json:"suppressedMessage,omitempty"`
	From                    string                      `json:"from,omitempty"`
	ReplyTo                 []string                    `json:"replyTo,omitempty"`
	Cc                      []string                    `json:"cc,omitempty"`
	Bcc                     []string                    `json:"bcc,omitempty"`
}

type EmailEventBrowser struct {
	Family      string   `json:"family"`
	Name        string   `json:"name"`
	Producer    string   `json:"producer"`
	ProducerUrl string   `json:"producerUrl"`
	Type        string   `json:"type"`
	Url         string   `json:"url"`
	Version     []string `json:"version"`
}

type EmailEventLocation struct {
	City      string  `json:"city"`
	Country   string  `json:"country"`
	State     string  `json:"state"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type EmailEventSentBy struct {
	Created int64  `json:"created"`
	Id      string `json:"id"`
}

type EmailEventSubscription struct {
	Id              int64                      `json:"id"`
	Status          string                     `json:"status"`
	LegalBasisChange *EmailEventLegalBasisChange `json:"legalBasisChange,omitempty"`
}

type EmailEventLegalBasisChange struct {
	LegalBasisType        string `json:"legalBasisType"`
	LegalBasisExplanation string `json:"legalBasisExplanation"`
	OptState              string `json:"optState"`
}

type EmailEventsResponse struct {
	Events  []EmailEvent `json:"events"`
	HasMore bool         `json:"hasMore"`
	Offset  string       `json:"offset"`
}
