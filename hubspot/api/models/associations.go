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
