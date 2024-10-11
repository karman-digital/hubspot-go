package hubspotmodels

type CRMCardResponse struct {
	Response string `json:"response"`
	Message  string `json:"message"`
}

type CrmCardAction struct {
	Type                       string   `json:"type"`
	HttpMethod                 string   `json:"httpMethod"`
	Width                      int      `json:"width"`
	Height                     int      `json:"height"`
	URI                        string   `json:"uri"`
	Label                      string   `json:"label"`
	AssociatedObjectProperties []string `json:"associatedObjectProperties,omitempty"`
}

type CrmCard struct {
	PrimaryAction *CrmCardAction  `json:"primaryAction,omitempty"`
	Results       []CrmCardResult `json:"results,omitempty"`
}

type CrmCardResult struct {
	ObjectId int    `json:"objectId"`
	Title    string `json:"title"`
	Link     string `json:"link"`
}
