package hubspotmodels

type Payload struct {
	ObjectID         int    `json:"objectId"`
	PropertyName     string `json:"propertyName"`
	PropertyValue    string `json:"propertyValue"`
	ChangeSource     string `json:"changeSource,omitempty"`
	SubscriptionID   int    `json:"subscriptionId"`
	SubscriptionType string `json:"subscriptionType"`
	EventID          int64  `json:"eventId"`
	PortalID         int    `json:"portalId"`
	AppID            int    `json:"appId"`
	OccurredAt       int64  `json:"occurredAt"`
	AttemptNumber    int    `json:"attemptNumber"`
	SourceId         string `json:"sourceId,omitempty"`
}
