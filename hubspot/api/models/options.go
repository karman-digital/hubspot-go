package hubspotmodels

type GetOptions struct {
	Properties            []string `url:"properties,omitempty"`
	PropertiesWithHistory []string `url:"propertiesWithHistory,omitempty"`
	Associations          []string `url:"associations,omitempty"`
	Archived              bool     `url:"archived,omitempty"`
	After                 string   `url:"after,omitempty"`
	Limit                 int      `url:"limit,omitempty"`
	IdProperty            string   `url:"idProperty,omitempty"`
}
