package batchcustomobjects

import "github.com/karman-digital/hubspot/hubspot/api/credentials"

func NewBatchCustomObjectService(creds *credentials.Credentials) *BatchCustomObjectService {
	return &BatchCustomObjectService{
		creds,
	}
}
