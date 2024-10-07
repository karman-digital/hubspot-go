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

type AccountDataPayload struct {
	ActionType string `json:"actionType"`
	PortalId   string `json:"portalId"`
	UserId     string `json:"userId"`
	UserEmail  string `json:"userEmail"`
	AppId      string `json:"appId"`
}

type GenericPayload struct {
	ObjectID                int    `json:"objectId"`
	PropertyName            string `json:"propertyName,omitempty"`
	PropertyValue           string `json:"propertyValue,omitempty"`
	ChangeSource            string `json:"changeSource,omitempty"`
	SubscriptionID          int    `json:"subscriptionId"`
	SubscriptionType        string `json:"subscriptionType"`
	EventID                 int64  `json:"eventId"`
	PortalID                int    `json:"portalId"`
	AppID                   int    `json:"appId"`
	OccurredAt              int64  `json:"occurredAt"`
	AttemptNumber           int    `json:"attemptNumber"`
	ObjectTypeID            string `json:"objectTypeId"`
	IsSensitive             bool   `json:"isSensitive,omitempty"`
	ChangeFlag              string `json:"changeFlag,omitempty"`
	NewObjectID             int    `json:"newObjectId,omitempty"`
	PrimaryObjectID         int    `json:"primaryObjectId,omitempty"`
	MergedObjectIDs         []int  `json:"mergedObjectIds,omitempty"`
	NumberOfPropertiesMoved int    `json:"numberOfPropertiesMoved,omitempty"`
	AssociationType         string `json:"associationType,omitempty"`
	FromObjectID            int    `json:"fromObjectId,omitempty"`
	ToObjectID              int    `json:"toObjectId,omitempty"`
	AssociationRemoved      bool   `json:"associationRemoved,omitempty"`
	IsPrimaryAssociation    bool   `json:"isPrimaryAssociation,omitempty"`
	AssociationCategory     string `json:"associationCategory,omitempty"`
	AssociationTypeID       int    `json:"associationTypeId,omitempty"`
	FromObjectTypeID        string `json:"fromObjectTypeId,omitempty"`
	ToObjectTypeID          string `json:"toObjectTypeId,omitempty"`
}
