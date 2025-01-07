package users

import "github.com/karman-digital/hubspot/hubspot/api/credentials"

func NewUserService(creds *credentials.Credentials) *UsersService {
	return &UsersService{
		Credentials: creds,
	}
}
