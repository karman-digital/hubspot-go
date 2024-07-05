package company

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	batchcompany "github.com/karman-digital/hubspot/hubspot/api/crm/company/batch"
)

func NewCompanyService(creds *credentials.Credentials) *CompanyService {
	return &CompanyService{
		creds,
		batchcompany.NewBatchCompanyService(creds),
	}
}
