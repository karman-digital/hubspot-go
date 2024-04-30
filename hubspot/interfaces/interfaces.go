package interfaces

import hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"

type Auth interface {
	RefreshTokenPair() error
	ValidateBearerToken() (bool, error)
}

type Associations interface {
	CreateDefaultAssociation(fromObject, toObject string, fromId, toId int) (hubspotmodels.BatchResponse, error)
	BatchCreateDefaultAssociations(fromObject, toObject string, associations hubspotmodels.BatchCreateDefaultAssociationsBody) (hubspotmodels.BatchResponse, error)
	BatchGetAssociations(fromObject, toObject string, body hubspotmodels.BatchGetAssociationsBody) (hubspotmodels.BatchAssociationGetResponse, error)
}

type Contact interface {
	Batch
	UpdateContact(id int, patchBody hubspotmodels.PatchBody) (hubspotmodels.ContactResponse, error)
	CreateContact(body hubspotmodels.PostBody) (hubspotmodels.ContactResponse, error)
	SearchContacts(body hubspotmodels.SearchBody) (hubspotmodels.SearchResponse, error)
	GetContact(id int, opts ...hubspotmodels.GetOptions) (hubspotmodels.Result, error)
}

type Deal interface {
	Batch
	UpdateDeal(id int, patchBody hubspotmodels.PatchBody) (hubspotmodels.Result, error)
	CreateDeal(body hubspotmodels.PostBody) (hubspotmodels.Result, error)
	GetDeal(id int, opts ...hubspotmodels.GetOptions) (hubspotmodels.Result, error)
	SearchDeals(body hubspotmodels.SearchBody) (hubspotmodels.SearchResponse, error)
}

type CustomObject interface {
	CustomBatch
	UpdateCustomObject(id int, patchBody hubspotmodels.PatchBody, objectType string) (hubspotmodels.Result, error)
	CreateCustomObject(body hubspotmodels.PostBody, objectType string) (hubspotmodels.Result, error)
	GetCustomObject(id int, objectType string, opts ...hubspotmodels.GetOptions) (hubspotmodels.Result, error)
	SearchCustomObjects(body hubspotmodels.SearchBody, objectType string) (hubspotmodels.SearchResponse, error)
}

type Owners interface {
	GetOwner(id int) (hubspotmodels.Owner, error)
	GetOwners(after ...string) (hubspotmodels.OwnerResponse, error)
	GetAllOwners() ([]hubspotmodels.Owner, error)
}

type Properties interface {
	CreatePropertyGroup(propertyGroup hubspotmodels.PropertyGroupBody, objectType string) error
	CreateProperty(objectType string, propertyData hubspotmodels.PropertyBody) error
}

type CommunicationPreferences interface {
	GetCommunicationPreferences() (hubspotmodels.CommunicationPreferencesResponse, error)
	UnsubscribeFromCommunicationPreference(contactEmail string, subscriptionId int, legalOptions ...hubspotmodels.CommunicationLegalBasis) error
	SubscribeToCommunicationPreference(contactEmail string, subscriptionId int, legalOptions ...hubspotmodels.CommunicationLegalBasis) error
	GetCommunicationPreferenceStatus(contactEmail string) (hubspotmodels.CommunicationPreferenceStatusResponse, error)
}

type Batch interface {
	BatchCreate(body hubspotmodels.BatchCreateBody) (hubspotmodels.BatchResponse, error)
	BatchGet(body hubspotmodels.BatchGetBody) (hubspotmodels.BatchResponse, error)
	BatchUpdate(body hubspotmodels.BatchUpdateBody) (hubspotmodels.BatchResponse, error)
}

type CustomBatch interface {
	BatchCreate(body hubspotmodels.BatchCreateBody, objectType string) (hubspotmodels.BatchResponse, error)
	BatchGet(body hubspotmodels.BatchGetBody, objectType string) (hubspotmodels.BatchResponse, error)
	BatchUpdate(body hubspotmodels.BatchUpdateBody, objectType string) (hubspotmodels.BatchResponse, error)
}
