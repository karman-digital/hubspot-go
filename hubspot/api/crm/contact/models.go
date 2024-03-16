package contact

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	"github.com/karman-digital/hubspot/hubspot/interfaces"
)

type ContactService struct {
	interfaces.Batch
	creds *credentials.Credentials
}
