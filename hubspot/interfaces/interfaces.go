package interfaces

import (
	blogmodels "github.com/karman-digital/hubspot/hubspot/api/models/cms/blogs"
	blogtagmodels "github.com/karman-digital/hubspot/hubspot/api/models/cms/blogtags"
	hubdbmodels "github.com/karman-digital/hubspot/hubspot/api/models/cms/hubdb"
	communicationmodels "github.com/karman-digital/hubspot/hubspot/api/models/communicationpreferences"
	crmmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm"
	associationsmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm/associations"
	listsmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm/lists"
	notemodels "github.com/karman-digital/hubspot/hubspot/api/models/crm/notes"
	ownersmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm/owners"
	propertiesmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm/properties"
	taskmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm/tasks"
	filesmodels "github.com/karman-digital/hubspot/hubspot/api/models/files"
	graphqlmodels "github.com/karman-digital/hubspot/hubspot/api/models/graphql"
	sharedmodels "github.com/karman-digital/hubspot/hubspot/api/models/shared"
	usermodels "github.com/karman-digital/hubspot/hubspot/api/models/users"
)

type Auth interface {
	RefreshTokenPair() error
	ValidateBearerToken() (bool, error)
}

type Associations interface {
	CreateDefaultAssociation(fromObject, toObject string, fromId, toId int) (crmmodels.BatchResponse, error)
	BatchCreateDefaultAssociations(fromObject, toObject string, associations associationsmodels.BatchCreateDefaultAssociationsBody) (crmmodels.BatchResponse, error)
	BatchGetAssociations(fromObject, toObject string, body associationsmodels.BatchGetAssociationsBody) (associationsmodels.BatchAssociationGetResponse, error)
	GetAssociations(fromObject, toObject string, id int) (associationsmodels.AssociationGetResponse, error)
	BatchCreateAssociations(fromObject, toObject string, body associationsmodels.BatchCreateAssociationsBody) (crmmodels.BatchResponse, error)
	CreateAssociation(fromObject, toObject, fromObjectType, toObjectType string, body []associationsmodels.AssociationType) (crmmodels.Result, error)
	BatchArchiveAssociationLabels(fromObject, toObject string, body associationsmodels.BatchCreateAssociationsBody) error
}

type Contact interface {
	Batch
	UpdateContact(id int, patchBody crmmodels.PatchBody) (crmmodels.Result, error)
	CreateContact(body crmmodels.PostBody) (crmmodels.Result, error)
	SearchContacts(body crmmodels.SearchBody) (crmmodels.SearchResponse, error)
	GetContact(id int, opts ...sharedmodels.GetOptions) (crmmodels.Result, error)
	GetContactByUniqueProperty(value string, opts ...sharedmodels.GetOptions) (crmmodels.Result, error)
	DeleteContact(id int) error
}

type Company interface {
	Batch
	UpdateCompany(id int, patchBody crmmodels.PatchBody) (crmmodels.Result, error)
	CreateCompany(body crmmodels.PostBody) (crmmodels.Result, error)
	SearchCompanies(body crmmodels.SearchBody) (crmmodels.SearchResponse, error)
	GetCompany(id int, opts ...sharedmodels.GetOptions) (crmmodels.Result, error)
	GetCompanyByUniqueProperty(value string, opts ...sharedmodels.GetOptions) (crmmodels.Result, error)
	GetCompanies(opts ...sharedmodels.GetOptions) (crmmodels.ListResponse, error)
	DeleteCompany(id int) error
}

type Deal interface {
	Batch
	UpdateDeal(id int, patchBody crmmodels.PatchBody) (crmmodels.Result, error)
	CreateDeal(body crmmodels.PostBody) (crmmodels.Result, error)
	GetDeal(id int, opts ...sharedmodels.GetOptions) (crmmodels.Result, error)
	SearchDeals(body crmmodels.SearchBody) (crmmodels.SearchResponse, error)
	DeleteDeal(id int) error
	GetDealByUniqueProperty(value string, opts ...sharedmodels.GetOptions) (crmmodels.Result, error)
	GetDeals(opts ...sharedmodels.GetOptions) (crmmodels.ListResponse, error)
}

type CustomObject interface {
	CustomBatch
	UpdateCustomObject(id int, patchBody crmmodels.PatchBody, objectType string) (crmmodels.Result, error)
	UpdateCustomObjectByUniqueId(id, idProperty string, patchBody crmmodels.PatchBody, objectType string) (crmmodels.Result, error)
	CreateCustomObject(body crmmodels.PostBody, objectType string) (crmmodels.Result, error)
	GetCustomObject(id int, objectType string, opts ...sharedmodels.GetOptions) (crmmodels.Result, error)
	SearchCustomObjects(body crmmodels.SearchBody, objectType string) (crmmodels.SearchResponse, error)
	GetCustomObjects(objectType string, opts ...sharedmodels.GetOptions) (crmmodels.ListResponse, error)
	GetCustomObjectByUniqueProperty(id string, objectType string, opts ...sharedmodels.GetOptions) (crmmodels.Result, error)
	DeleteCustomObject(id int, objectType string) error
}

type Owners interface {
	GetOwner(id int) (ownersmodels.Owner, error)
	GetOwners(opts ...ownersmodels.GetOwnersOptions) (ownersmodels.OwnerResponse, error)
	GetAllOwners() ([]ownersmodels.Owner, error)
}

type Properties interface {
	CreatePropertyGroup(propertyGroup propertiesmodels.PropertyGroupBody, objectType string) error
	CreateProperty(objectType string, propertyData propertiesmodels.PropertyBody) error
	GetProperty(objectType, propertyName string) (propertiesmodels.PropertyResponse, error)
	UpdateProperty(objectType, propertyName string, propertyData propertiesmodels.PropertyBody) (propertiesmodels.PropertyResponse, error)
}

type CommunicationPreferences interface {
	BatchPreferences
	GetCommunicationPreferences() (communicationmodels.CommunicationPreferencesResponse, error)
	UnsubscribeFromCommunicationPreference(contactEmail string, subscriptionId int, legalOptions ...communicationmodels.CommunicationLegalBasis) error
	SubscribeToCommunicationPreference(contactEmail string, subscriptionId int, legalOptions ...communicationmodels.CommunicationLegalBasis) error
	GetCommunicationPreferenceStatus(contactEmail string) (communicationmodels.CommunicationPreferenceStatusResponse, error)
}

type Batch interface {
	BatchCreate(body crmmodels.BatchCreateBody) (crmmodels.BatchResponse, error)
	BatchGet(body crmmodels.BatchGetBody) (crmmodels.BatchResponse, error)
	BatchDelete(body crmmodels.BatchDeleteBody) error
	BatchUpdate(body crmmodels.BatchUpdateBody) (crmmodels.BatchResponse, error)
}

type CustomBatch interface {
	BatchCreate(body crmmodels.BatchCreateBody, objectType string) (crmmodels.BatchResponse, error)
	BatchGet(body crmmodels.BatchGetBody, objectType string) (crmmodels.BatchResponse, error)
	BatchUpdate(body crmmodels.BatchUpdateBody, objectType string) (crmmodels.BatchResponse, error)
	BatchDelete(body crmmodels.BatchDeleteBody, objectType string) error
}

type Notes interface {
	Batch
	CreateNoteWithAssociations(noteBody notemodels.NotePostBody, associations ...associationsmodels.ObjectCreationAssociation) (crmmodels.Result, error)
	GetNote(noteId string, opts ...sharedmodels.GetOptions) (crmmodels.Result, error)
}

type Tasks interface {
	CreateTaskWithAssociations(taskBody taskmodels.TaskPostBody, associations ...associationsmodels.ObjectCreationAssociation) (crmmodels.Result, error)
}

type Calls interface {
	Batch
	GetCall(id string, opts ...sharedmodels.GetOptions) (crmmodels.Result, error)
}

type Emails interface {
	Batch
	GetEmail(id string, opts ...sharedmodels.GetOptions) (crmmodels.Result, error)
}

type BatchPreferences interface {
	BatchUpdateCommunicationPreferences(body communicationmodels.BatchCommunicationPreferencesPostBody) (communicationmodels.BatchCommunicationPreferencesResponse, error)
}

type Products interface {
	Batch
	GetProductByUniqueId(uniqueId string, opts ...sharedmodels.GetOptions) (crmmodels.Result, error)
}

type LineItems interface {
	Batch
	CreateLineItem(body crmmodels.PostBody) (crmmodels.Result, error)
	GetLineItem(id int, opts ...sharedmodels.GetOptions) (crmmodels.Result, error)
	UpdateLineItem(id int, patchBody crmmodels.PatchBody) (crmmodels.Result, error)
}

type Files interface {
	ImportFileViaUrl(body filesmodels.FileImportBody) (filesmodels.FileUploadResult, error)
	GetSignedUrl(fileId string, signedUrlOptions ...filesmodels.SignedUrlOptions) (filesmodels.SignedUrlResponse, error)
	UploadFile(fileName string, fileContent []byte, opts ...filesmodels.UploadFileOptions) (filesmodels.FileUploadResult, error)
	GetFileByPath(filePath string, properties ...[]string) (filesmodels.FileStatResponse, error)
	UpdateFile(fileId string, fileName string, fileContent []byte, opts ...filesmodels.UpdateFileOptions) (filesmodels.FileUploadResult, error)
}

type Users interface {
	Create(body usermodels.UserBody) (usermodels.UserBody, error)
}

type Blog interface {
	BlogTags
	GetAllBlogPosts(opts blogmodels.BlogFilterOptions) (blogmodels.BlogPostsResponse, error)
}

type BlogTags interface {
	GetBatchBlogTags(opts blogtagmodels.BlogTagsBatchInput) (blogtagmodels.BatchBlogTagResponse, error)
}

type HubDB interface {
	GetTableRow(tableId, rowId string) (hubdbmodels.HubDBRowResponse, error)
}

type Meetings interface {
	GetMeeting(id string, opts ...sharedmodels.GetOptions) (crmmodels.Result, error)
}

type GraphQL interface {
	MakeRequest(query string, variables map[string]interface{}) (map[string]interface{}, error)
	MakeRequestWithFullResponse(query string, variables map[string]interface{}) (graphqlmodels.GraphQLResponse, error)
}

type Lists interface {
	SearchLists(body listsmodels.SearchListsBody) (listsmodels.SearchListsResponse, error)
	GetLists(listIds []string, includeFilters bool) (listsmodels.ListsByIdResponse, error)
	GetListMemberships(listId string, opts ...sharedmodels.GetOptions) (listsmodels.ListMembershipsResponse, error)
}
