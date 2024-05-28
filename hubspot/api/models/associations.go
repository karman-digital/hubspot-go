package hubspotmodels

type AssociationPair struct {
	From AssociationId `json:"from"`
	To   AssociationId `json:"to"`
}

type AssociationId struct {
	Id string `json:"id"`
}

type BatchCreateDefaultAssociationsBody struct {
	Inputs []AssociationPair `json:"inputs"`
}

type BatchCreateAssociationBody struct {
	Inputs []BatchCreateAssociationInput `json:"inputs"`
}

type BatchCreateAssociationInput struct {
	Types []AssociationType `json:"types"`
	AssociationPair
}

type AssociationType struct {
	AssociationCategory string `json:"associationCategory"`
	AssociationTypeId   int    `json:"associationTypeId"`
}

type ObjectCreationAssociation struct {
	Types []AssociationType `json:"types"`
	To    AssociationId     `json:"to"`
}

type BatchAssociationGetResponse struct {
	CompletedAt string                   `json:"completedAt"`
	RequestedAt string                   `json:"requestedAt"`
	StartedAt   string                   `json:"startedAt"`
	NumErrors   int                      `json:"numErrors"`
	Links       map[string]string        `json:"links"`
	Results     []BatchAssociationResult `json:"results"`
	Errors      []ErrorDetail            `json:"errors,omitempty"`
	Status      string                   `json:"status"`
}

type BatchAssociationResult struct {
	From   From     `json:"from"`
	Paging Paging   `json:"paging"`
	To     []ToItem `json:"to"`
}

type From struct {
	ID string `json:"id"`
}

type Navigation struct {
	Link   string `json:"link"`
	After  string `json:"after,omitempty"`
	Before string `json:"before,omitempty"`
}

type ToItem struct {
	AssociationTypes []BatchAssociationType `json:"associationTypes"`
	ToObjectId       int                    `json:"toObjectId"`
}

type BatchAssociationType struct {
	TypeId   int    `json:"typeId"`
	Label    string `json:"label"`
	Category string `json:"category"`
}

type BatchGetAssociationsBody struct {
	Inputs []BatchGetAssociationsInput `json:"inputs"`
}

type BatchGetAssociationsInput struct {
	Id    string `json:"id"`
	After string `json:"after,omitempty"`
}

type ErrorDetail struct {
	SubCategory any                 `json:"subCategory"`
	Context     map[string][]string `json:"context"`
	Links       map[string]string   `json:"links"`
	ID          string              `json:"id"`
	Category    string              `json:"category"`
	Message     string              `json:"message"`
	Errors      []NestedError       `json:"errors"`
	Status      string              `json:"status"`
}

type NestedError struct {
	SubCategory string              `json:"subCategory"`
	Code        string              `json:"code"`
	In          string              `json:"in"`
	Context     map[string][]string `json:"context"`
	Message     string              `json:"message"`
}

type AssociationGetResponse struct {
	Paging  Paging   `json:"paging"`
	Results []ToItem `json:"results"`
}
