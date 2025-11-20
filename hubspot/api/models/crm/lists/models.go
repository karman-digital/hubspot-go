package lists

import (
	"time"

	sharedmodels "github.com/karman-digital/hubspot/hubspot/api/models/shared"
)

type SearchListsBody struct {
	Query              string   `json:"query,omitempty"`
	ProcessingTypes    []string `json:"processingTypes,omitempty"`
	AdditionalProperties []string `json:"additionalProperties,omitempty"`
	Limit              int      `json:"limit,omitempty"`
	After              string   `json:"after,omitempty"`
}

type List struct {
	ListId              string                 `json:"listId"`
	Name                string                 `json:"name"`
	ObjectTypeId        string                 `json:"objectTypeId"`
	ProcessingType      string                 `json:"processingType"`
	CreatedAt           time.Time              `json:"createdAt"`
	UpdatedAt           time.Time              `json:"updatedAt"`
	Size                int                    `json:"size,omitempty"`
	AdditionalProperties map[string]interface{} `json:"additionalProperties,omitempty"`
	FilterBranch        *FilterBranch          `json:"filterBranch,omitempty"`
}

type FilterBranch struct {
	FilterBranchType string      `json:"filterBranchType"`
	Filters          []Filter    `json:"filters,omitempty"`
	Branches         []FilterBranch `json:"branches,omitempty"`
}

type Filter struct {
	PropertyName string      `json:"propertyName"`
	Operator     string      `json:"operator"`
	Value        interface{} `json:"value,omitempty"`
	Values       []interface{} `json:"values,omitempty"`
}

type SearchListsResponse struct {
	Results []List              `json:"results"`
	Paging  sharedmodels.Paging `json:"paging,omitempty"`
}

type ListMembership struct {
	RecordId string `json:"recordId"`
}

type ListMembershipsResponse struct {
	Results []ListMembership    `json:"results"`
	Paging  sharedmodels.Paging `json:"paging,omitempty"`
}

