package deals

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	batchdeal "github.com/karman-digital/hubspot/hubspot/api/crm/deals/batch"
)

func NewDealService(creds *credentials.Credentials) *DealService {
	return &DealService{
		creds,
		batchdeal.NewBatchDealService(creds),
	}
}
