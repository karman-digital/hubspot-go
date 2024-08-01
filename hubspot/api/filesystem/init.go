package filesystem

import (
	"github.com/karman-digital/hubspot/hubspot/api/credentials"
	"github.com/karman-digital/hubspot/hubspot/api/filesystem/files"
)

func NewFilesystemService(creds *credentials.Credentials) Filesystem {
	return Filesystem{
		Files: files.NewFilesService(creds),
	}
}
