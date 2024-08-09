package batchlineitems

import "github.com/karman-digital/hubspot/hubspot/api/credentials"

func NewBatchLineItemService(creds *credentials.Credentials) *BatchLineItemService {
	return &BatchLineItemService{
		creds,
	}
}
