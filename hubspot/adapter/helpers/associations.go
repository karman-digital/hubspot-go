package hshelpers

import associationsmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm/associations"

func CreateBatchCreateDefaultAssociationsBody(associations []associationsmodels.AssociationPair) associationsmodels.BatchCreateDefaultAssociationsBody {
	return associationsmodels.BatchCreateDefaultAssociationsBody{
		Inputs: associations,
	}
}
