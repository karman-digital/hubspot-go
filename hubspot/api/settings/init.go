package settings

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	"github.com/karman-digital/hubspot/hubspot/api/settings/users"
)

func NewSettingsService(creds *credentials.Credentials) Settings {
	return Settings{
		Users: users.NewUserService(creds),
	}
}
