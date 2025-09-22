package webhookmodels

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

type AccountDataPayload struct {
	ActionType string `json:"actionType"`
	PortalId   string `json:"portalId"`
	UserId     string `json:"userId"`
	UserEmail  string `json:"userEmail"`
	AppId      string `json:"appId"`
}

type GenericPayload struct {
	ObjectId                int    `json:"objectId"`
	PropertyName            string `json:"propertyName,omitempty"`
	PropertyValue           string `json:"propertyValue,omitempty"`
	ChangeSource            string `json:"changeSource,omitempty"`
	SubscriptionId          int    `json:"subscriptionId"`
	SubscriptionType        string `json:"subscriptionType"`
	EventId                 int64  `json:"eventId"`
	PortalId                int    `json:"portalId"`
	AppId                   int    `json:"appId"`
	OccurredAt              int64  `json:"occurredAt"`
	AttemptNumber           int    `json:"attemptNumber"`
	ObjectTypeId            string `json:"objectTypeId"`
	IsSensitive             bool   `json:"isSensitive,omitempty"`
	ChangeFlag              string `json:"changeFlag,omitempty"`
	NewObjectId             int    `json:"newObjectId,omitempty"`
	PrimaryObjectId         int    `json:"primaryObjectId,omitempty"`
	MergedObjectIds         []int  `json:"mergedObjectIds,omitempty"`
	NumberOfPropertiesMoved int    `json:"numberOfPropertiesMoved,omitempty"`
	AssociationType         string `json:"associationType,omitempty"`
	FromObjectId            int    `json:"fromObjectId,omitempty"`
	ToObjectId              int    `json:"toObjectId,omitempty"`
	AssociationRemoved      bool   `json:"associationRemoved,omitempty"`
	IsPrimaryAssociation    bool   `json:"isPrimaryAssociation,omitempty"`
	AssociationCategory     string `json:"associationCategory,omitempty"`
	AssociationTypeId       int    `json:"associationTypeId,omitempty"`
	FromObjectTypeId        string `json:"fromObjectTypeId,omitempty"`
	ToObjectTypeId          string `json:"toObjectTypeId,omitempty"`
	SourceId                string `json:"sourceId,omitempty"`
}
