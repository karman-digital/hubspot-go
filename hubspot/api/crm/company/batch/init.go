package batchcompany

import "github.com/karman-digital/hubspot/hubspot/api/credentials"

func NewBatchCompanyService(creds *credentials.Credentials) *BatchCompanyService {
	return &BatchCompanyService{
		creds,
	}
}
