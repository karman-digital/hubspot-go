package memberships

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
)

func NewMembershipsService(creds *credentials.Credentials) *MembershipsService {
	return &MembershipsService{
		creds,
	}
}

