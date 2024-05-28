package notes

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/hashicorp/go-retryablehttp"
	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
)

func (n *NotesService) CreateNoteWithAssociations(noteBody hubspotmodels.NotePostBody, associations ...hubspotmodels.ObjectCreationAssociation) (hubspotmodels.ObjectResponse, error) {
	var notesResp hubspotmodels.ObjectResponse
	reqUrl := "https://api.hubapi.com/crm/v3/objects/notes"
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
	req, err := retryablehttp.NewRequest("POST", reqUrl, reqBody)
	if err != nil {
		return notesResp, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", n.AccessToken()))
	resp, err := n.Client().Do(req)
	if err != nil {
		return notesResp, fmt.Errorf("error making request: %s", err)
	}
	defer resp.Body.Close()
	contactRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return notesResp, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 201 {
		return notesResp, fmt.Errorf("error returned by endpoint: %s", contactRawBody)
	}
	err = json.Unmarshal(contactRawBody, &resp)
	if err != nil {
		return notesResp, fmt.Errorf("error parsing body: %s", err)
	}
	return notesResp, nil
}
