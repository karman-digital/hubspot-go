package schemasmodels

type Schema struct {
	Associations []AssociationDefinition `json:"associations"`
}

type AssociationDefinition struct {
	ID               string `json:"id"`
	FromObjectTypeID string `json:"fromObjectTypeId"`
	ToObjectTypeID   string `json:"toObjectTypeId"`
}
