package products

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	"github.com/karman-digital/hubspot/hubspot/interfaces"
)

type ProductService struct {
	*credentials.Credentials
	interfaces.Batch
}
