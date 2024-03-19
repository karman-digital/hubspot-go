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

type ClientSecret string

func (c ClientSecret) String() string {
	return string(c)
}

func (c *ClientSecret) Set(s string) {
	*c = ClientSecret(s)
}

type ClientId string

func (c ClientId) String() string {
	return string(c)
}

func (c *ClientId) Set(s string) {
	*c = ClientId(s)
}

type RedirectUri string

func (r RedirectUri) String() string {
	return string(r)
}

func (r *RedirectUri) Set(s string) {
	*r = RedirectUri(s)
}
