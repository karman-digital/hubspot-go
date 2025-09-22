package notemodels

import associationsmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm/associations"

type NotePostBody struct {
	Properties   NoteProperties                                 `json:"properties"`
	Associations []associationsmodels.ObjectCreationAssociation `json:"associations,omitempty"`
}

type NoteProperties struct {
	HsTimestamp     string `json:"hs_timestamp"`
	HsNoteBody      string `json:"hs_note_body"`
	HubspotOwnerId  int    `json:"hubspot_owner_id,omitempty"`
	HsAttachmentIds string `json:"hs_attachment_ids,omitempty"`
}
