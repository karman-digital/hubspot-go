package batchdeal

import "github.com/karman-digital/hubspot/hubspot/api/credentials"

func NewBatchDealService(creds *credentials.Credentials) *BatchDealService {
	return &BatchDealService{
		creds,
	}
}
