package batchnotes

import "github.com/karman-digital/hubspot/hubspot/api/credentials"

func NewBatchNotesService(creds *credentials.Credentials) *BatchNotesService {
	return &BatchNotesService{
		creds,
	}
}
