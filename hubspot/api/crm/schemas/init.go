package schemas

import "github.com/karman-digital/hubspot/hubspot/api/credentials"

func NewSchemaService(creds *credentials.Credentials) *SchemaService {
	return &SchemaService{Credentials: creds}
}
