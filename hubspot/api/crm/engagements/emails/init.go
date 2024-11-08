package emails

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	batchemails "github.com/karman-digital/hubspot/hubspot/api/crm/engagements/emails/batch"
)

func NewEmailsService(creds *credentials.Credentials) *EmailsService {
	return &EmailsService{
		BatchEmailsService: batchemails.NewBatchEmailsService(creds),
		Credentials:        creds,
	}
}
