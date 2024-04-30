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

type BatchAssociationGetResponse struct {
	CompletedAt string                   `json:"completedAt"`
	RequestedAt string                   `json:"requestedAt"`
	StartedAt   string                   `json:"startedAt"`
	Links       map[string]string        `json:"links"`
	Results     []BatchAssociationResult `json:"results"`
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
	ToObjectId       string                 `json:"toObjectId"`
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
