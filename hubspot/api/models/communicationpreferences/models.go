package communicationmodels

import (
	"time"

	sharedmodels "github.com/karman-digital/hubspot/hubspot/api/models/shared"
)

type CommunicationPreferencesResponse struct {
	sharedmodels.BatchResponseBase
	Results []SubscriptionDefinition `json:"results"`
}

type CommunicationPreferenceStatusResponse struct {
	sharedmodels.BatchResponseBase
	Results []PublicStatus `json:"results"`
}

type SubscriptionDefinition struct {
	BusinessUnitID      int64     `json:"businessUnitId"`
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
	Channel        string `json:"channel"`
	StatusState    string `json:"statusState"`
	SubscriptionID int64  `json:"subscriptionId"`
	CommunicationLegalBasis
}

type CommunicationLegalBasis struct {
	LegalBasis            string `json:"legalBasis,omitempty"`
	LegalBasisExplanation string `json:"legalBasisExplanation,omitempty"`
}

type PublicStatus struct {
	BusinessUnitID         int64     `json:"businessUnitId"`
	Channel                string    `json:"channel"`
	SubscriberIDString     string    `json:"subscriberIdString"`
	SubscriptionID         int64     `json:"subscriptionId"`
	SubscriptionName       string    `json:"subscriptionName"`
	Status                 string    `json:"status"`
	Source                 string    `json:"source"`
	LegalBasis             *string   `json:"legalBasis"`
	LegalBasisExplanation  *string   `json:"legalBasisExplanation"`
	SetStatusSuccessReason *string   `json:"setStatusSuccessReason"`
	Timestamp              time.Time `json:"timestamp"`
}

// SubscriptionStatus is retained as an alias for callers migrating from the
// legacy response model.
type SubscriptionStatus = PublicStatus

type BatchCommunicationPreferencesPostBody struct {
	Inputs []CommunicationPreferencesBatchInput `json:"inputs"`
}

type CommunicationPreferencesBatchInput struct {
	StatusState           string `json:"statusState"`
	Channel               string `json:"channel"`
	SubscriberIdString    string `json:"subscriberIdString"`
	LegalBasis            string `json:"legalBasis,omitempty"`
	SubscriptionId        int64  `json:"subscriptionId"`
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
	SetStatusSuccessReason string    `json:"setStatusSuccessReason"`
	Source                 string    `json:"source"`
	SubscriptionId         int64     `json:"subscriptionId"`
	LegalBasisExplanation  string    `json:"legalBasisExplanation"`
	BusinessUnitId         int64     `json:"businessUnitId"`
	Status                 string    `json:"status"`
	Timestamp              time.Time `json:"timestamp"`
}
