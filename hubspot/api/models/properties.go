package hubspotmodels

type PropertyGroupBody struct {
	Name  string `json:"name"`
	Label string `json:"label"`
}

type PropertyBody struct {
	Name      string               `json:"name"`
	Label     string               `json:"label"`
	GroupName string               `json:"groupName"`
	Type      string               `json:"type"`
	FieldType string               `json:"fieldType"`
	Options   []EnumerationOptions `json:"options,omitempty"`
}

type EnumerationOptions struct {
	Label  string `json:"label"`
	Value  string `json:"value"`
	Hidden bool   `json:"hidden"`
}
