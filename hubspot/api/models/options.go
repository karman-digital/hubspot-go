package hubspotmodels

type GetOptions struct {
	Properties            []string `url:"properties,omitempty"`
	PropertiesWithHistory []string `url:"propertiesWithHistory,omitempty"`
	Associations          []string `url:"associations,omitempty"`
	Archived              bool     `url:"archived,omitempty"`
}
