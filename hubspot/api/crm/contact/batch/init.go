package batchcontact

import "github.com/karman-digital/hubspot/hubspot/api/credentials"

func NewBatchContactService(creds *credentials.Credentials) *BatchContactService {
	return &BatchContactService{
		creds: creds,
	}
}
