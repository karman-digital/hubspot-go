package filesmodels

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
	Status      string           `json:"status"`
	Result      FileUploadResult `json:"result"`
	StartedAt   string           `json:"startedAt"`
	CompletedAt string           `json:"completedAt"`
	TaskID      string           `json:"taskId"`
}

type FileUploadResult struct {
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

type SignedUrlResponse struct {
	URL       string `json:"url"`
	ExpiresAt string `json:"expiresAt"`
	Name      string `json:"name"`
	Extension string `json:"extension"`
	Type      string `json:"type"`
	Size      int64  `json:"size"`
}

type SignedUrlOptions struct {
	ExpirationSeconds int64  `json:"expirationSeconds"`
	Size              string `json:"size"`
	Upscale           bool   `json:"upscale"`
}

type UploadFileOptions struct {
	FolderId   string        `json:"folderId"`
	FolderPath string        `json:"folderPath"`
	Options    UploadOptions `json:"options"`
}

type UploadOptions struct {
	Access string `json:"access"`
	Ttl    string `json:"ttl"`
}

type UpdateFileOptions struct {
	CharsetHunch string        `json:"charsetHunch,omitempty"`
	Options      UpdateOptions `json:"options"`
}

type UpdateOptions struct {
	Access    string `json:"access,omitempty"`
	ExpiresAt string `json:"expiresAt,omitempty"`
}

type FileStatResponse struct {
	File   FileStat   `json:"file"`
	Folder FolderStat `json:"folder,omitempty"`
}

type FileStat struct {
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

type FolderStat struct {
	ID             string `json:"id"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
	Archived      bool   `json:"archived"`
	ParentFolderID string `json:"parentFolderId"`
	Name          string `json:"name"`
	Path          string `json:"path"`
}
