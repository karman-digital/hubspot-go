package hubspotmodels

type FileImportResponse struct {
	Links FileImportStatus `json:"links"`
	ID    string           `json:"id"`
}

type FileImportStatus struct {
	Status string `json:"status"`
}

type FileImportBody struct {
	Name                        string `json:"name"`
	DuplicateValidationStrategy string `json:"duplicateValidationStrategy,omitempty"`
	TTL                         string `json:"ttl,omitempty"`
	Overwrite                   bool   `json:"overwrite,omitempty"`
	FolderID                    string `json:"folderId,omitempty"`
	FolderPath                  string `json:"folderPath,omitempty"`
	Access                      string `json:"access"`
	DuplicateValidationScope    string `json:"duplicateValidationScope,omitempty"`
	URL                         string `json:"url"`
}
