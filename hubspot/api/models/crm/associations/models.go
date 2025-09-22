package associationsmodels

import (
	sharedmodels "github.com/karman-digital/hubspot/hubspot/api/models/shared"
)

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

type Association struct {
	Results []AssociationResultPair `json:"results,omitempty"`
}

type AssociationResultPair struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type BatchAssociationGetResponse struct {
	CompletedAt string                     `json:"completedAt"`
	RequestedAt string                     `json:"requestedAt"`
	StartedAt   string                     `json:"startedAt"`
	NumErrors   int                        `json:"numErrors"`
	Links       map[string]string          `json:"links"`
	Results     []BatchAssociationResult   `json:"results"`
	Errors      []sharedmodels.ErrorDetail `json:"errors,omitempty"`
	Status      string                     `json:"status"`
}

type BatchAssociationResult struct {
	From   From                `json:"from"`
	Paging sharedmodels.Paging `json:"paging"`
	To     []ToItem            `json:"to"`
}

type From struct {
	ID string `json:"id"`
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

type AssociationGetResponse struct {
	Paging  sharedmodels.Paging `json:"paging"`
	Results []ToItem            `json:"results"`
}

type BatchCreateAssociationsBody struct {
	Inputs []AssociationInput `json:"inputs"`
}

type AssociationInput struct {
	Types []AssociationType `json:"types"`
	From  AssociationObject `json:"from"`
	To    AssociationObject `json:"to"`
}

type AssociationObject struct {
	Id string `json:"id"`
}
