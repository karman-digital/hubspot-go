package auth

import "github.com/karman-digital/hubspot/hubspot/api/credentials"

func NewAuthService(creds *credentials.Credentials) *AuthService {
	return &AuthService{
		creds,
	}
}
