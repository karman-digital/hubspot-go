package pipelines

import "github.com/karman-digital/hubspot/hubspot/api/credentials"

type PipelineService struct {
	*credentials.Credentials
}

func NewPipelineService(creds *credentials.Credentials) *PipelineService {
	return &PipelineService{Credentials: creds}
}
