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

type Result struct {
	ID           string       `json:"id"`
	Properties   Properties   `json:"properties"`
	CreatedAt    time.Time    `json:"createdAt"`
	UpdatedAt    time.Time    `json:"updatedAt"`
	Archived     bool         `json:"archived"`
	Associations Associations `json:"associations"`
}

type Associations struct {
	Companies AssociationResult `json:"companies"`
	Deals     AssociationResult `json:"deals"`
	Contacts  AssociationResult `json:"contacts"`
}

type AssociationResult struct {
	Results []AssociationResultPair `json:"results,omitempty"`
}

type AssociationResultPair struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type ErrorDetail struct {
	SubCategory map[string]interface{} `json:"subCategory"`
	Context     map[string][]string    `json:"context"`
	Links       map[string]string      `json:"links"`
	ID          string                 `json:"id"`
	Category    string                 `json:"category"`
	Message     string                 `json:"message"`
	Errors      []NestedError          `json:"errors"`
	Status      string                 `json:"status"`
}

type NestedError struct {
	SubCategory string              `json:"subCategory"`
	Code        string              `json:"code"`
	In          string              `json:"in"`
	Context     map[string][]string `json:"context"`
	Message     string              `json:"message"`
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
