package batchemails

import "github.com/karman-digital/hubspot/hubspot/api/credentials"

func NewBatchEmailsService(creds *credentials.Credentials) *BatchEmailsService {
	return &BatchEmailsService{
		creds,
	}
}
