package hubspotmodels

import "time"

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
