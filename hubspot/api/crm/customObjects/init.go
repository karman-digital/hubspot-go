package customObjects

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	batchcustomobjects "github.com/karman-digital/hubspot/hubspot/api/crm/customObjects/batch"
)

func NewCustomObjectService(creds *credentials.Credentials) *CustomObjectService {
	return &CustomObjectService{
		creds,
		batchcustomobjects.NewBatchCustomObjectService(creds),
	}
}
