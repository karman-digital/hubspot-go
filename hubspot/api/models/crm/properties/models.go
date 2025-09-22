package propertiesmodels

type PropertyGroupBody struct {
	Name  string `json:"name"`
	Label string `json:"label"`
}

type PropertyBody struct {
	Name      string               `json:"name,omitempty"`
	Label     string               `json:"label,omitempty"`
	GroupName string               `json:"groupName,omitempty"`
	Type      string               `json:"type,omitempty"`
	FieldType string               `json:"fieldType,omitempty"`
	Options   []EnumerationOptions `json:"options,omitempty"`
}

type EnumerationOptions struct {
	Label        string `json:"label"`
	Value        string `json:"value"`
	Hidden       bool   `json:"hidden"`
	Description  string `json:"description,omitempty"`
	DisplayOrder int    `json:"displayOrder,omitempty"`
}

type ModificationMetadata struct {
	ReadOnlyOptions    bool `json:"readOnlyOptions"`
	ReadOnlyValue      bool `json:"readOnlyValue"`
	ReadOnlyDefinition bool `json:"readOnlyDefinition"`
	Archivable         bool `json:"archivable"`
}

type PropertyResponse struct {
	PropertyBody
	Hidden               bool                 `json:"hidden"`
	DisplayOrder         int                  `json:"displayOrder"`
	HasUniqueValue       bool                 `json:"hasUniqueValue"`
	FormField            bool                 `json:"formField"`
	ModificationMetadata ModificationMetadata `json:"modificationMetadata"`
}
