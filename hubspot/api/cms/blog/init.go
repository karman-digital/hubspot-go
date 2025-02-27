package blogs

import (
	blogtags "github.com/karman-digital/hubspot/hubspot/api/cms/blog/tags"
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
)

func NewBlogService(creds *credentials.Credentials) *BlogService {
	return &BlogService{
		creds,
		blogtags.NewBlogTagsService(creds),
	}
}
