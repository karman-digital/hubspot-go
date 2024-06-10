package hubspotmodels

import "time"

type ObjectResponse struct {
	Id         string         `json:"id"`
	Properties map[string]any `json:"properties"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	Archived   bool           `json:"archived"`
}

type Result struct {
	ID           string                 `json:"id"`
	Properties   Properties             `json:"properties"`
	CreatedAt    time.Time              `json:"createdAt"`
	UpdatedAt    time.Time              `json:"updatedAt"`
	Archived     bool                   `json:"archived"`
	Associations map[string]Association `json:"associations"`
}

type ListResponse struct {
	Results []Result `json:"results"`
	Paging  Paging   `json:"paging"`
}
