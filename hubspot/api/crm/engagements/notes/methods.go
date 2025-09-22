package notes

import (
	"encoding/json"
	"fmt"
	"net/http"

	crmmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm"
	associationsmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm/associations"
	notemodels "github.com/karman-digital/hubspot/hubspot/api/models/crm/notes"
	sharedmodels "github.com/karman-digital/hubspot/hubspot/api/models/shared"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (n *NotesService) CreateNoteWithAssociations(noteBody notemodels.NotePostBody, associations ...associationsmodels.ObjectCreationAssociation) (crmmodels.Result, error) {
	var notesResp crmmodels.Result
	for _, association := range associations {
		if noteBody.Associations == nil {
			noteBody.Associations = []associationsmodels.ObjectCreationAssociation{}
		}
		noteBody.Associations = append(noteBody.Associations, association)
	}
	reqBody, err := json.Marshal(noteBody)
	if err != nil {
		return notesResp, err
	}
	resp, err := n.SendRequest(http.MethodPost, "/crm/v3/objects/notes", reqBody)
	if err != nil {
		return crmmodels.Result{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleCreateResponse(resp)
}

func (n *NotesService) GetNote(noteId string, opts ...sharedmodels.GetOptions) (crmmodels.Result, error) {
	resp, err := n.SendRequest(http.MethodGet, fmt.Sprintf("/crm/v3/objects/notes/%s", noteId), nil, opts...)
	if err != nil {
		return crmmodels.Result{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleResponse(resp)
}
