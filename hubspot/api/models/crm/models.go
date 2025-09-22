package crmmodels

import (
	"time"

	associationsmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm/associations"
	sharedmodels "github.com/karman-digital/hubspot/hubspot/api/models/shared"
)

type ContactResponse struct {
	Id         string         `json:"id"`
	Properties map[string]any `json:"properties"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	Archived   bool           `json:"archived"`
}

type ObjectResponse struct {
	Id         string         `json:"id"`
	Properties map[string]any `json:"properties"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	Archived   bool           `json:"archived"`
}

type Properties map[string]any

type PatchBody struct {
	Properties Properties `json:"properties"`
}

type PostBody struct {
	Properties   Properties                                     `json:"properties"`
	Associations []associationsmodels.ObjectCreationAssociation `json:"associations,omitempty"`
}

type Result struct {
	ID           string                                    `json:"id"`
	Properties   Properties                                `json:"properties"`
	CreatedAt    time.Time                                 `json:"createdAt"`
	UpdatedAt    time.Time                                 `json:"updatedAt"`
	Archived     bool                                      `json:"archived"`
	Associations map[string]associationsmodels.Association `json:"associations"`
}

type ListResponse struct {
	Results []Result            `json:"results"`
	Paging  sharedmodels.Paging `json:"paging"`
}

type BatchResponse struct {
	sharedmodels.BatchResponseBase
	Results []Result `json:"results"`
}

type BatchUpdateBody struct {
	Inputs []BatchUpdateInput `json:"inputs"`
}

type BatchCreateBody struct {
	Inputs []PostBody `json:"inputs"`
}

type BatchDeleteBody struct {
	Inputs []BatchInput `json:"inputs"`
}

type BatchUpdateInput struct {
	Id         string         `json:"id"`
	Properties map[string]any `json:"properties"`
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

type BatchInput struct {
	Id int `json:"id"`
}

type SearchBody struct {
	Query        string        `json:"query,omitempty"`
	Limit        int           `json:"limit,omitempty"`
	After        string        `json:"after,omitempty"`
	Sorts        []string      `json:"sorts,omitempty"`
	Properties   []string      `json:"properties,omitempty"`
	FilterGroups []FilterGroup `json:"filterGroups,omitempty"`
}

type FilterGroup struct {
	Filters []Filter `json:"filters,omitempty"`
}

type Filter struct {
	HighValue    string   `json:"highValue,omitempty"`
	PropertyName string   `json:"propertyName,omitempty"`
	Values       []string `json:"values,omitempty"`
	Value        string   `json:"value,omitempty"`
	Operator     string   `json:"operator,omitempty"`
}

type SearchResponse struct {
	Total   int                 `json:"total,omitempty"`
	Paging  sharedmodels.Paging `json:"paging,omitempty"`
	Results []Result            `json:"results,omitempty"`
}
