package company

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	"github.com/karman-digital/hubspot/hubspot/interfaces"
)

type CompanyService struct {
	*credentials.Credentials
	interfaces.Batch
}
