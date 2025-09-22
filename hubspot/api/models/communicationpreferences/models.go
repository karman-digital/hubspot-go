package communicationmodels

import (
	"time"

	sharedmodels "github.com/karman-digital/hubspot/hubspot/api/models/shared"
)

type CommunicationPreferencesResponse struct {
	SubscriptionDefinitions []SubscriptionDefinition `json:"subscriptionDefinitions"`
}

type CommunicationPreferenceStatusResponse struct {
	Recipient            string               `json:"recipient"`
	SubscriptionStatuses []SubscriptionStatus `json:"subscriptionStatuses"`
}

type SubscriptionDefinition struct {
	ID                  string    `json:"id"`
	Name                string    `json:"name"`
	Description         string    `json:"description"`
	Purpose             string    `json:"purpose"`
	CommunicationMethod string    `json:"communicationMethod"`
	IsActive            bool      `json:"isActive"`
	IsDefault           bool      `json:"isDefault"`
	IsInternal          bool      `json:"isInternal"`
	CreatedAt           time.Time `json:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt"`
}

type CommunicationPreferencesPostBody struct {
	EmailAddress   string `json:"emailAddress"`
	SubscriptionId string `json:"subscriptionId"`
	CommunicationLegalBasis
}

type CommunicationLegalBasis struct {
	LegalBasis            string `json:"legalBasis,omitempty"`
	LegalBasisExplanation string `json:"legalBasisExplanation,omitempty"`
}

type SubscriptionStatus struct {
	ID                    string `json:"id"`
	Name                  string `json:"name"`
	Description           string `json:"description"`
	Status                string `json:"status"`
	SourceOfStatus        string `json:"sourceOfStatus"`
	PreferenceGroupName   string `json:"preferenceGroupName"`
	LegalBasis            string `json:"legalBasis"`
	LegalBasisExplanation string `json:"legalBasisExplanation"`
}

type BatchCommunicationPreferencesPostBody struct {
	Inputs []CommunicationPreferencesBatchInput `json:"inputs"`
}

type CommunicationPreferencesBatchInput struct {
	StatusState           string `json:"statusState"`
	Channel               string `json:"channel"`
	SubscriberIdString    string `json:"subscriberIdString"`
	LegalBasis            string `json:"legalBasis,omitempty"`
	SubscriptionId        int    `json:"subscriptionId"`
	LegalBasisExplanation string `json:"legalBasisExplanation,omitempty"`
}

type BatchCommunicationPreferencesResponse struct {
	sharedmodels.BatchResponseBase
	Results []V4CommunicationPreferenceResult `json:"results"`
}

type V4CommunicationPreferenceResult struct {
	Channel                string    `json:"channel"`
	SubscriberIdString     string    `json:"subscriberIdString"`
	LegalBasis             string    `json:"legalBasis"`
	SetStatisSuccessReason string    `json:"setStatisSuccessReason"`
	Source                 string    `json:"source"`
	SubscriptionId         int       `json:"subscriptionId"`
	LegalBasisExplanation  string    `json:"legalBasisExplanation"`
	BusinessUnitId         int       `json:"businessUnitId"`
	Status                 string    `json:"status"`
	TimeStamp              time.Time `json:"timeStamp"`
}
