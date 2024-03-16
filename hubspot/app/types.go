package hubspotapp

import (
	"errors"
	"strconv"
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
