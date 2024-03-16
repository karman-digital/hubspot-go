package properties

import "github.com/karman-digital/hubspot/hubspot/api/credentials"

func NewPropertiesService(creds *credentials.Credentials) *PropertiesService {
	return &PropertiesService{
		creds: creds,
	}
}
