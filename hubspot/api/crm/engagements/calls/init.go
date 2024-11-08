package calls

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	batchcalls "github.com/karman-digital/hubspot/hubspot/api/crm/engagements/calls/batch"
)

func NewCallsService(creds *credentials.Credentials) *CallsService {
	return &CallsService{
		BatchCallsService: batchcalls.NewBatchCallsService(creds),
		Credentials:       creds,
	}
}
