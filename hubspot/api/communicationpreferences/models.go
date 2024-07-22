package communicationpreferences

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	"github.com/karman-digital/hubspot/hubspot/interfaces"
)

type CommunicationPreferencesService struct {
	*credentials.Credentials
	interfaces.BatchPreferences
}
