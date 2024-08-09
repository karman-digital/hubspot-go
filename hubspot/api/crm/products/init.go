package products

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	batchproducts "github.com/karman-digital/hubspot/hubspot/api/crm/products/batch"
)

func NewProductService(creds *credentials.Credentials) *ProductService {
	return &ProductService{
		creds,
		batchproducts.NewBatchProductService(creds),
	}
}
