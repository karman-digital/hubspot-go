package lineItems

import "github.com/karman-digital/hubspot/hubspot/api/credentials"

func NewLineItemsService(creds *credentials.Credentials) *LineItemsService {
	return &LineItemsService{
		creds,
	}
}
