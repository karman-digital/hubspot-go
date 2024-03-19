package communicationpreferences

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
)

func NewCommunicationPreferencesService(creds *credentials.Credentials) *CommunicationPreferencesService {
	return &CommunicationPreferencesService{
		creds,
	}
}
