package associations

import "github.com/karman-digital/hubspot/hubspot/api/credentials"

func NewAssociationService(creds *credentials.Credentials) *AssociationService {
	return &AssociationService{
		creds,
	}
}
