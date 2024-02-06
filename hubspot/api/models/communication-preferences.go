package hubspotmodels

import "time"

type CommunicationPreferencesResponse struct {
	SubscriptionDefinitions []SubscriptionDefinition `json:"subscriptionDefinitions"`
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
