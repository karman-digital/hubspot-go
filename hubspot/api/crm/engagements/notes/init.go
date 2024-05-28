package notes

import "github.com/karman-digital/hubspot/hubspot/api/credentials"

func NewNotesService(creds *credentials.Credentials) *NotesService {
	return &NotesService{
		creds,
	}
}
