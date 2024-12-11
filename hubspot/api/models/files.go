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

type FileImportStatusResponse struct {
	Status      string     `json:"status"`
	Result      FileResult `json:"result"`
	StartedAt   string     `json:"startedAt"`
	CompletedAt string     `json:"completedAt"`
	TaskID      string     `json:"taskId"`
}

type FileResult struct {
	ID                string `json:"id"`
	CreatedAt         string `json:"createdAt"`
	UpdatedAt         string `json:"updatedAt"`
	Archived          bool   `json:"archived"`
	ParentFolderID    string `json:"parentFolderId"`
	Name              string `json:"name"`
	Path              string `json:"path"`
	Size              int64  `json:"size"`
	Type              string `json:"type"`
	Extension         string `json:"extension"`
	DefaultHostingURL string `json:"defaultHostingUrl"`
	URL               string `json:"url"`
	IsUsableInContent bool   `json:"isUsableInContent"`
	Access            string `json:"access"`
	FileMD5           string `json:"fileMd5"`
	SourceGroup       string `json:"sourceGroup"`
}
