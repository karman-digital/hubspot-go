package blogs

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	"github.com/karman-digital/hubspot/hubspot/interfaces"
)

type BlogService struct {
	*credentials.Credentials
	interfaces.BlogTags
}
