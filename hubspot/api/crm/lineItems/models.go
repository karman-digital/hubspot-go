package lineItems

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	"github.com/karman-digital/hubspot/hubspot/interfaces"
)

type LineItemsService struct {
	*credentials.Credentials
	interfaces.Batch
}
