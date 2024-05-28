package notes

import hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"

func (n *Notes) CreateNoteWithAssociations(body hubspotmodels.NotePostBody, associations ...hubspotmodels.ObjectCreationAssociation) (hubspotmodels.ObjectResponse, error) {
	return hubspotmodels.ObjectResponse{}, nil
}
