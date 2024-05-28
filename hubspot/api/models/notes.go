package hubspotmodels

type NotePostBody struct {
	Properties   NoteProperties              `json:"properties"`
	Associations []ObjectCreationAssociation `json:"associations"`
}

type NoteProperties struct {
	HsTimestamp     int64  `json:"hs_timestamp"`
	HsNoteBody      string `json:"hs_note_body"`
	HubspotOwnerId  int    `json:"hubspot_owner_id"`
	HsAttachmentIds string `json:"hs_attachment_ids"`
}
