package listdata

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
)

func NewListDataService(creds *credentials.Credentials) *ListDataService {
	return &ListDataService{
		creds,
	}
}

