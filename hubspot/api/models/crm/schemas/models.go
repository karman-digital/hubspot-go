package schemasmodels

type Schema struct {
	ObjectTypeID string                  `json:"objectTypeId"`
	Associations []AssociationDefinition `json:"associations"`
}

type AssociationDefinition struct {
	ID               string `json:"id"`
	FromObjectTypeID string `json:"fromObjectTypeId"`
	ToObjectTypeID   string `json:"toObjectTypeId"`
}
