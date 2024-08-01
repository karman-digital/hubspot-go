package hubspotapp

import (
	"errors"
	"strconv"

	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	"github.com/karman-digital/hubspot/hubspot/api/crm"
	"github.com/karman-digital/hubspot/hubspot/api/filesystem"
	"github.com/karman-digital/hubspot/hubspot/interfaces"
)

type PortalId int

func (p *PortalId) Set(i int) error {
	if i == 0 {
		return errors.New("portal id cannot be zero")
	}
	*p = PortalId(i)
	return nil
}

func (p PortalId) Int() int {
	return int(p)
}

func (p PortalId) String() string {
	return strconv.Itoa(int(p))
}

type ApiClient struct {
	CRM crm.CRM
	interfaces.CommunicationPreferences
	FileSystem filesystem.Filesystem
}

type Hubspot struct {
	*credentials.Credentials
	ApiClient
	PortalId
}
