package calls

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	batchcalls "github.com/karman-digital/hubspot/hubspot/api/crm/engagements/calls/batch"
)

type CallsService struct {
	*credentials.Credentials
	*batchcalls.BatchCallsService
}
