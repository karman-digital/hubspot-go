package hubspotmodels

import "time"

type ContactResponse struct {
	Id         string         `json:"id"`
	Properties map[string]any `json:"properties"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	Archived   bool           `json:"archived"`
}

type BatchResponse struct {
	CompletedAt time.Time         `json:"completedAt"`
	NumErrors   int               `json:"numErrors"`
	RequestedAt time.Time         `json:"requestedAt"`
	StartedAt   time.Time         `json:"startedAt"`
	Links       map[string]string `json:"links"`
	Results     []Result          `json:"results"`
	Errors      []ErrorDetail     `json:"errors"`
	Status      string            `json:"status"`
}

type BatchUpdateBody struct {
	Inputs []BatchUpdateInput `json:"inputs"`
}

type BatchCreateBody struct {
	Inputs []PostBody `json:"inputs"`
}

type BatchUpdateInput struct {
	Id         string         `json:"id"`
	Properties map[string]any `json:"properties"`
}

type Association struct {
	Results []AssociationResultPair `json:"results,omitempty"`
}

type AssociationResultPair struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type BatchGetBody struct {
	PropertiesWithHistory []string        `json:"propertiesWithHistory,omitempty"`
	IDProperty            string          `json:"idProperty,omitempty"`
	Inputs                []BatchGetInput `json:"inputs,omitempty"`
	Properties            []string        `json:"properties,omitempty"`
}

type BatchGetInput struct {
	ID string `json:"id,omitempty"`
}
