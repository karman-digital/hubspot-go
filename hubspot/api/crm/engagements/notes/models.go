package notes

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	"github.com/karman-digital/hubspot/hubspot/interfaces"
)

type NotesService struct {
	interfaces.Batch
	*credentials.Credentials
}
