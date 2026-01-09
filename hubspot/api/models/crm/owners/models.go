package ownersmodels

import (
	"time"

	sharedmodels "github.com/karman-digital/hubspot/hubspot/api/models/shared"
)

type OwnerTeam struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Primary bool   `json:"primary"`
}

type Owner struct {
	ID        string      `json:"id"`
	Email     string      `json:"email"`
	FirstName string      `json:"firstName"`
	LastName  string      `json:"lastName"`
	UserID    int         `json:"userId"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
	Archived  bool        `json:"archived"`
	Teams     []OwnerTeam `json:"teams"`
}

type OwnerResponse struct {
	Results []Owner             `json:"results"`
	Paging  sharedmodels.Paging `json:"paging"`
}

type GetOwnersOptions struct {
	After    string `url:"after,omitempty"`
	Archived bool   `url:"archived,omitempty"`
	Email    string `url:"email,omitempty"`
	Limit    int    `url:"limit,omitempty"`
}
