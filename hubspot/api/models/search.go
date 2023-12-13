package hubspotmodels

type SearchBody struct {
	Query        string        `json:"query,omitempty"`
	Limit        int           `json:"limit,omitempty"`
	After        string        `json:"after,omitempty"`
	Sorts        []string      `json:"sorts,omitempty"`
	Properties   []string      `json:"properties,omitempty"`
	FilterGroups []FilterGroup `json:"filterGroups,omitempty"`
}

type FilterGroup struct {
	Filters []Filter `json:"filters,omitempty"`
}

type Filter struct {
	HighValue    string   `json:"highValue,omitempty"`
	PropertyName string   `json:"propertyName,omitempty"`
	Values       []string `json:"values,omitempty"`
	Value        string   `json:"value,omitempty"`
	Operator     string   `json:"operator,omitempty"`
}

type SearchResponse struct {
	Total   int      `json:"total,omitempty"`
	Paging  Paging   `json:"paging,omitempty"`
	Results []Result `json:"results,omitempty"`
}

type Paging struct {
	Next NextPage `json:"next,omitempty"`
}

type NextPage struct {
	After string `json:"after,omitempty"`
	Link  string `json:"link,omitempty"`
}
