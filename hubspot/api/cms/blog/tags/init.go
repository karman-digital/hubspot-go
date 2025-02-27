package blogtags

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
)

func NewBlogTagsService(creds *credentials.Credentials) *BlogTagsService {
	return &BlogTagsService{
		creds,
	}
}
