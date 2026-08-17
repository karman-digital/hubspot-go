package schemasmodels

type AssociationDefinitionBody struct {
	FromObjectTypeID string `json:"fromObjectTypeId"`
	ToObjectTypeID   string `json:"toObjectTypeId"`
}
