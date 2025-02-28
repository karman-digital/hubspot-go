package hubspotmodels

type HubDBRowValues map[string]interface{}

type HubDBRowResponse struct {
	Path         string         `json:"path"`
	CreatedAt    string         `json:"createdAt"`
	ChildTableID string         `json:"childTableId"`
	PublishedAt  string         `json:"publishedAt"`
	Values       HubDBRowValues `json:"values"`
	Name         string         `json:"name"`
	ID           string         `json:"id"`
	UpdatedAt    string         `json:"updatedAt"`
}
