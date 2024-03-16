package hubspotmodels

import "errors"

type AccessToken string
type RefreshToken string

func (r *RefreshToken) Set(s string) error {
	if s == "" {
		return errors.New("refresh token cannot be empty")
	}
	*r = RefreshToken(s)
	return nil
}

func (r RefreshToken) String() string {
	return string(r)
}

func (a *AccessToken) Set(s string) error {
	if s == "" {
		return errors.New("access token cannot be empty")
	}
	*a = AccessToken(s)
	return nil
}

func (a AccessToken) String() string {
	return string(a)
}
