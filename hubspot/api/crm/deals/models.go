package deals

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	"github.com/karman-digital/hubspot/hubspot/interfaces"
)

type DealService struct {
	*credentials.Credentials
	interfaces.Batch
}
