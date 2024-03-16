package credentials

import (
	"errors"
	"strconv"

	"github.com/hashicorp/go-retryablehttp"
)

type Credentials struct {
	Client       *retryablehttp.Client
	AccessToken  AccessToken
	RefreshToken RefreshToken
	PortalId     PortalId
}

type PortalId int
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
