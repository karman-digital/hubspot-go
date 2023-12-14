package hshelpers

import hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"

func CreateBatchCreateDefaultAssociationsBody(associations []hubspotmodels.AssociationPair) hubspotmodels.BatchCreateDefaultAssociationsBody {
	return hubspotmodels.BatchCreateDefaultAssociationsBody{
		Inputs: associations,
	}
}
