package emailmodels

import "time"

type MarketingEmail struct {
	Id        string                 `json:"id"`
	Name      string                 `json:"name"`
	Campaign  string                 `json:"campaign,omitempty"`
	Content   map[string]interface{} `json:"content,omitempty"`
	CreatedAt time.Time              `json:"createdAt,omitempty"`
	UpdatedAt time.Time              `json:"updatedAt,omitempty"`
	// Additional fields may be present depending on email configuration
	// Reference the API documentation for complete structure
}

