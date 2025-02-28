package cms

import (
	blogs "github.com/karman-digital/hubspot/hubspot/api/cms/blog"
	"github.com/karman-digital/hubspot/hubspot/api/cms/hubdb"
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
)

func NewCmsService(creds *credentials.Credentials) CMS {
	return CMS{
		Blogs: blogs.NewBlogService(creds),
		HubDB: hubdb.NewHubDBService(creds),
	}
}
