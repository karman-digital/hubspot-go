package hsconvert

import (
	domainmodels "github.com/karman-digital/hatch-domain/domain/models"
	"github.com/karman-digital/hubspot/hubspot/models"
)

func ConvertContactToClientDomain(contact models.ObjectBody, properties []domainmodels.SyncValues) domainmodels.Client {
	client := domainmodels.Client{}
	return client
}
