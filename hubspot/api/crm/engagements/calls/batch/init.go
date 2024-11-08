package batchcalls

import "github.com/karman-digital/hubspot/hubspot/api/credentials"

func NewBatchCallsService(creds *credentials.Credentials) *BatchCallsService {
	return &BatchCallsService{
		creds,
	}
}
