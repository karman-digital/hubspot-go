package models

import "time"

type ContactResponse struct {
	Id         string         `json:"id"`
	Properties map[string]any `json:"properties"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	Archived   bool           `json:"archived"`
}
