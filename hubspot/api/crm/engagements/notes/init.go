package notes

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	batchnotes "github.com/karman-digital/hubspot/hubspot/api/crm/engagements/notes/batch"
)

func NewNotesService(creds *credentials.Credentials) *NotesService {
	return &NotesService{
		batchnotes.NewBatchNotesService(creds),
		creds,
	}
}
