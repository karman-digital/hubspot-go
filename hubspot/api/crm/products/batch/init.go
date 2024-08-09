package batchproducts

import "github.com/karman-digital/hubspot/hubspot/api/credentials"

func NewBatchProductService(creds *credentials.Credentials) *BatchProductService {
	return &BatchProductService{
		creds,
	}
}
