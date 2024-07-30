package products

import "github.com/karman-digital/hubspot/hubspot/api/credentials"

func NewProductService(creds *credentials.Credentials) *ProductService {
	return &ProductService{
		creds,
	}
}
