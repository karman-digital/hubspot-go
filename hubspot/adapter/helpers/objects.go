package hshelpers

import sharedmodels "github.com/karman-digital/hubspot/hubspot/api/models/shared"

func CreateContactGetOptions(properties []string, propertiesWithHistory []string, associations []string, archived bool) sharedmodels.GetOptions {
	return sharedmodels.GetOptions{
		Properties:            properties,
		PropertiesWithHistory: propertiesWithHistory,
		Associations:          associations,
		Archived:              archived,
	}
}

func CreateGetOptions(properties []string, propertiesWithHistory []string, associations []string, archived bool) sharedmodels.GetOptions {
	return sharedmodels.GetOptions{
		Properties:            properties,
		PropertiesWithHistory: propertiesWithHistory,
		Associations:          associations,
		Archived:              archived,
	}
}
