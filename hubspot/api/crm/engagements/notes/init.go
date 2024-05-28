package notes

import "github.com/karman-digital/hubspot/hubspot/api/credentials"

func NewNotesService(creds *credentials.Credentials) *Notes {
	return &Notes{
		creds,
	}
}
