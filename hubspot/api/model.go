package hubspot

import (
	"github.com/hashicorp/go-retryablehttp"
	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
)

type credentials struct {
	Client       *retryablehttp.Client
	AccessToken  AccessToken
	RefreshToken RefreshToken
	PortalId     PortalId
}

type HubspotAPI interface {
	RetrieveAccessToken() *AccessToken
	RetrieveRefreshToken() *RefreshToken
	RetrievePortalId() *PortalId
	SetAccessToken(accessToken string) error
	SetRefreshToken(refreshToken string) error
	SetPortalId(portalId int) error
	RefreshTokenPair(clientSecret string, clientId string, redirectUri string) error
	ValidateBearerToken() (bool, error)
	UpdateContact(id int, patchBody hubspotmodels.PatchBody) (hubspotmodels.ContactResponse, error)
	BatchCreateContact(body hubspotmodels.BatchCreateBody) (hubspotmodels.BatchResponse, error)
	CreateContact(body hubspotmodels.PostBody) (hubspotmodels.ContactResponse, error)
	SearchContacts(body hubspotmodels.SearchBody) (hubspotmodels.SearchResponse, error)
	BatchGetContacts(body hubspotmodels.BatchGetBody) (hubspotmodels.BatchResponse, error)
	BatchUpdateContacts(body hubspotmodels.BatchUpdateBody) (hubspotmodels.BatchResponse, error)
	CreateDefaultAssociation(fromObject, toObject string, fromId, toId int) (hubspotmodels.BatchResponse, error)
	BatchCreateDefaultAssociations(fromObject, toObject string, associations hubspotmodels.BatchCreateDefaultAssociationsBody) (hubspotmodels.BatchResponse, error)
	CreatePropertyGroup(propertyGroup hubspotmodels.PropertyGroupBody, objectType string) error
	CreateProperty(objectType string, propertyData hubspotmodels.PropertyBody) error
	GetOwner(id int) (hubspotmodels.Owner, error)
	GetOwners(after ...string) (hubspotmodels.OwnerResponse, error)
	GetProperty(ObjectType string, PropertyName string) (hubspotmodels.PropertyResponse, error)
	UpdateProperty(ObjectType string, PropertyName string, propertyData hubspotmodels.PropertyBody) error
	GetContact(id int, opts ...hubspotmodels.ContactGetOptions) (hubspotmodels.Result, error)
	GetAllOwners() ([]hubspotmodels.Owner, error)
	GetCommunicationPreferences() (hubspotmodels.CommunicationPreferencesResponse, error)
	UnsubscribeFromCommunicationPreference(contactEmail string, subscriptionId int, legalOptions ...hubspotmodels.CommunicationLegalBasis) error
	SubscribeToCommunicationPreference(contactEmail string, subscriptionId int, legalOptions ...hubspotmodels.CommunicationLegalBasis) error
}

type HubspotObject string

type hubspotDefaultObjects struct {
	Contact string
	Company string
	Deal    string
	Ticket  string
}

var DefaultObjectNames = hubspotDefaultObjects{
	Contact: "contacts",
	Company: "companies",
	Deal:    "deals",
	Ticket:  "tickets",
}
