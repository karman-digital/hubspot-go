package hshelpers

import hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"

func CreateContactGetOptions(properties []string, propertiesWithHistory []string, associations []string, archived bool) hubspotmodels.ContactGetOptions {
	return hubspotmodels.ContactGetOptions{
		Properties:            properties,
		PropertiesWithHistory: propertiesWithHistory,
		Associations:          associations,
		Archived:              archived,
	}
}
