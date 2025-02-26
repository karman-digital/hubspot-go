package blogs

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
)

func NewBlogService(creds *credentials.Credentials) *BlogService {
	return &BlogService{
		creds,
	}
}
