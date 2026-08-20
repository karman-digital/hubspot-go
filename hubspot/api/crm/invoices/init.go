package invoices

import "github.com/karman-digital/hubspot/hubspot/api/credentials"

type Service struct {
	*credentials.Credentials
}

func NewService(creds *credentials.Credentials) *Service {
	return &Service{Credentials: creds}
}
