package contact

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	batchcontact "github.com/karman-digital/hubspot/hubspot/api/crm/contact/batch"
)

func NewContactService(creds *credentials.Credentials) *ContactService {
	return &ContactService{
		creds: creds,
		Batch: batchcontact.NewBatchContactService(creds),
	}
}
