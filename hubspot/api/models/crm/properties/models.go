package propertiesmodels

type PropertyGroupBody struct {
	Name  string `json:"name"`
	Label string `json:"label"`
}

type PropertyBody struct {
	Name                 string               `json:"name,omitempty"`
	Label                string               `json:"label,omitempty"`
	GroupName            string               `json:"groupName,omitempty"`
	Type                 string               `json:"type,omitempty"`
	FieldType            string               `json:"fieldType,omitempty"`
	Options              []EnumerationOptions `json:"options,omitempty"`
	FormField            bool                 `json:"formField,omitempty"`
	HasUniqueValue       bool                 `json:"hasUniqueValue,omitempty"`
	ExternalOptions      bool                 `json:"externalOptions,omitempty"`
	ReferencedObjectType string               `json:"referencedObjectType,omitempty"`
}

type PropertyUpdateBody struct {
	Label              string               `json:"label"`
	Description        string               `json:"description"`
	GroupName          string               `json:"groupName"`
	Type               string               `json:"type"`
	FieldType          string               `json:"fieldType"`
	Options            []EnumerationOptions `json:"options"`
	FormField          bool                 `json:"formField"`
	Hidden             bool                 `json:"hidden"`
	DisplayOrder       int                  `json:"displayOrder"`
	CalculationFormula string               `json:"calculationFormula,omitempty"`
}

type EnumerationOptions struct {
	Label        string `json:"label"`
	Value        string `json:"value"`
	Hidden       bool   `json:"hidden"`
	Description  string `json:"description,omitempty"`
	DisplayOrder int    `json:"displayOrder"`
}

type ModificationMetadata struct {
	ReadOnlyOptions    bool `json:"readOnlyOptions"`
	ReadOnlyValue      bool `json:"readOnlyValue"`
	ReadOnlyDefinition bool `json:"readOnlyDefinition"`
	Archivable         bool `json:"archivable"`
}

type PropertyResponse struct {
	PropertyBody
	Description          string               `json:"description"`
	Hidden               bool                 `json:"hidden"`
	DisplayOrder         int                  `json:"displayOrder"`
	HasUniqueValue       bool                 `json:"hasUniqueValue"`
	FormField            bool                 `json:"formField"`
	Archived             bool                 `json:"archived"`
	CreatedAt            string               `json:"createdAt"`
	UpdatedAt            string               `json:"updatedAt"`
	CalculationFormula   string               `json:"calculationFormula"`
	ModificationMetadata ModificationMetadata `json:"modificationMetadata"`
}
