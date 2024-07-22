package batchcommunicationpreferences

import "github.com/karman-digital/hubspot/hubspot/api/credentials"

func NewBatchCommunicationPreferencesService(creds *credentials.Credentials) *BatchCommunicationPreferencesService {
	return &BatchCommunicationPreferencesService{
		creds,
	}
}
