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
	GetAssociations(fromObject, toObject string, id int) (hubspotmodels.AssociationGetResponse, error)
}

type Contact interface {
	Batch
	UpdateContact(id int, patchBody hubspotmodels.PatchBody) (hubspotmodels.ContactResponse, error)
	CreateContact(body hubspotmodels.PostBody) (hubspotmodels.ContactResponse, error)
	SearchContacts(body hubspotmodels.SearchBody) (hubspotmodels.SearchResponse, error)
	GetContact(id int, opts ...hubspotmodels.GetOptions) (hubspotmodels.Result, error)
	GetContactByUniqueProperty(value string, opts ...hubspotmodels.GetOptions) (hubspotmodels.Result, error)
}

type Company interface {
	Batch
	UpdateCompany(id int, patchBody hubspotmodels.PatchBody) (hubspotmodels.Result, error)
	CreateCompany(body hubspotmodels.PostBody) (hubspotmodels.Result, error)
	SearchCompanies(body hubspotmodels.SearchBody) (hubspotmodels.SearchResponse, error)
	GetCompany(id int, opts ...hubspotmodels.GetOptions) (hubspotmodels.Result, error)
	GetCompanyByUniqueProperty(value string, opts ...hubspotmodels.GetOptions) (hubspotmodels.Result, error)
}

type Deal interface {
	Batch
	UpdateDeal(id int, patchBody hubspotmodels.PatchBody) (hubspotmodels.Result, error)
	CreateDeal(body hubspotmodels.PostBody) (hubspotmodels.Result, error)
	GetDeal(id int, opts ...hubspotmodels.GetOptions) (hubspotmodels.Result, error)
	SearchDeals(body hubspotmodels.SearchBody) (hubspotmodels.SearchResponse, error)
	DeleteDeal(id int) error
	GetDealByUniqueProperty(value string, opts ...hubspotmodels.GetOptions) (hubspotmodels.Result, error)
}

type CustomObject interface {
	CustomBatch
	UpdateCustomObject(id int, patchBody hubspotmodels.PatchBody, objectType string) (hubspotmodels.Result, error)
	CreateCustomObject(body hubspotmodels.PostBody, objectType string) (hubspotmodels.Result, error)
	GetCustomObject(id int, objectType string, opts ...hubspotmodels.GetOptions) (hubspotmodels.Result, error)
	SearchCustomObjects(body hubspotmodels.SearchBody, objectType string) (hubspotmodels.SearchResponse, error)
	GetCustomObjects(objectType string, opts ...hubspotmodels.GetOptions) (hubspotmodels.ListResponse, error)
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
	BatchPreferences
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

type Notes interface {
	CreateNoteWithAssociations(noteBody hubspotmodels.NotePostBody, associations ...hubspotmodels.ObjectCreationAssociation) (hubspotmodels.Result, error)
}

type Tasks interface {
	CreateTaskWithAssociations(taskBody hubspotmodels.TaskPostBody, associations ...hubspotmodels.ObjectCreationAssociation) (hubspotmodels.Result, error)
}

type BatchPreferences interface {
	BatchUpdateCommunicationPreferences(body hubspotmodels.BatchCommunicationPreferencesPostBody) (hubspotmodels.BatchCommunicationPreferencesResponse, error)
}

type Products interface {
	Batch
	GetProductByUniqueId(uniqueId string, opts ...hubspotmodels.GetOptions) (hubspotmodels.Result, error)
}

type LineItems interface {
	Batch
	CreateLineItem(body hubspotmodels.PostBody) (hubspotmodels.Result, error)
	GetLineItem(id int, opts ...hubspotmodels.GetOptions) (hubspotmodels.Result, error)
	UpdateLineItem(id int, patchBody hubspotmodels.PatchBody) (hubspotmodels.Result, error)
}

type Files interface {
	ImportFileViaUrl(body hubspotmodels.FileImportBody) (hubspotmodels.FileImportResponse, error)
}
