package owners

import "github.com/karman-digital/hubspot/hubspot/api/credentials"

func NewOwnerService(creds *credentials.Credentials) *OwnerService {
	return &OwnerService{
		creds: creds,
	}
}
