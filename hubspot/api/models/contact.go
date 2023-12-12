package models

import "time"

type ContactResponse struct {
	Id         string         `json:"id"`
	Properties map[string]any `json:"properties"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	Archived   bool           `json:"archived"`
}

type BatchContactResponse struct {
	CompletedAt time.Time         `json:"completedAt"`
	RequestedAt time.Time         `json:"requestedAt"`
	StartedAt   time.Time         `json:"startedAt"`
	Links       map[string]string `json:"links"`
	Results     []Result          `json:"results"`
	Status      string            `json:"status"`
}

type Result struct {
	ID         string     `json:"id"`
	Properties Properties `json:"properties"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
	Archived   bool       `json:"archived"`
}
