package customObjects

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	"github.com/karman-digital/hubspot/hubspot/interfaces"
)

type CustomObjectService struct {
	*credentials.Credentials
	interfaces.CustomBatch
}
