package properties

import (
	"net/http"

	sharedmodels "github.com/karman-digital/hubspot/hubspot/api/models/shared"
)

type requestSender interface {
	SendRequest(method, path string, body []byte, opts ...sharedmodels.GetOptions) (*http.Response, error)
}

type PropertiesService struct {
	sender requestSender
}

func newPropertiesService(sender requestSender) *PropertiesService {
	return &PropertiesService{sender: sender}
}
