package notes

import (
	"encoding/json"
	"fmt"
	"net/http"

	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (n *NotesService) CreateNoteWithAssociations(noteBody hubspotmodels.NotePostBody, associations ...hubspotmodels.ObjectCreationAssociation) (hubspotmodels.Result, error) {
	var notesResp hubspotmodels.Result
	for _, association := range associations {
		if noteBody.Associations == nil {
			noteBody.Associations = []hubspotmodels.ObjectCreationAssociation{}
		}
		noteBody.Associations = append(noteBody.Associations, association)
	}
	reqBody, err := json.Marshal(noteBody)
	if err != nil {
		return notesResp, err
	}
	resp, err := n.SendRequest(http.MethodPost, "/crm/v3/objects/notes", reqBody)
	if err != nil {
		return hubspotmodels.Result{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleResponse(resp)

}
