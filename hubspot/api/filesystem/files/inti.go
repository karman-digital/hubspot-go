package files

import "github.com/karman-digital/hubspot/hubspot/api/credentials"

func NewFilesService(creds *credentials.Credentials) *FilesService {
	return &FilesService{
		Credentials: creds,
	}
}
