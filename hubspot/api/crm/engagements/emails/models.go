package emails

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	batchemails "github.com/karman-digital/hubspot/hubspot/api/crm/engagements/emails/batch"
)

type EmailsService struct {
	*batchemails.BatchEmailsService
	*credentials.Credentials
}
