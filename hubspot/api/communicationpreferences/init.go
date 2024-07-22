package communicationpreferences

import (
	batchcommunicationpreferences "github.com/karman-digital/hubspot/hubspot/api/communicationpreferences/batch"
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
)

func NewCommunicationPreferencesService(creds *credentials.Credentials) *CommunicationPreferencesService {
	return &CommunicationPreferencesService{
		creds,
		batchcommunicationpreferences.NewBatchCommunicationPreferencesService(creds),
	}
}
