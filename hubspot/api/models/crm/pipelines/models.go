package pipelinesmodels

type ListResponse struct {
	Results []Pipeline `json:"results"`
}

type Pipeline struct {
	ID       string  `json:"id"`
	Archived bool    `json:"archived"`
	Stages   []Stage `json:"stages"`
}

type Stage struct {
	ID       string   `json:"id"`
	Archived bool     `json:"archived"`
	Metadata Metadata `json:"metadata"`
}

type Metadata struct {
	IsClosed    string `json:"isClosed"`
	Probability string `json:"probability"`
}
