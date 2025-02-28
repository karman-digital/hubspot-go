package hubdb

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
)

func NewHubDBService(creds *credentials.Credentials) *HubDBService {
	return &HubDBService{
		creds,
	}
}
