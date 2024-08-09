package lineItems

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	batchlineitems "github.com/karman-digital/hubspot/hubspot/api/crm/lineItems/batch"
)

func NewLineItemsService(creds *credentials.Credentials) *LineItemsService {
	return &LineItemsService{
		creds,
		batchlineitems.NewBatchLineItemService(creds),
	}
}
