package blogtagmodels

type BatchBlogTagResponse struct {
	Status      string    `json:"status"`
	Results     []BlogTag `json:"results"`
	StartedAt   string    `json:"startedAt"`
	CompletedAt string    `json:"completedAt"`
}

type BlogTag struct {
	Created   string `json:"created"`
	DeletedAt string `json:"deletedAt"`
	ID        string `json:"id"`
	Language  string `json:"language"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	Updated   string `json:"updated"`
}

type BlogTagsBatchInput struct {
	Inputs []string `json:"inputs"`
}
