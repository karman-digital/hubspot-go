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
