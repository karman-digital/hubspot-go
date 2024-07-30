package hubspotmodels

type PostBody struct {
	Properties   Properties                  `json:"properties"`
	Associations []ObjectCreationAssociation `json:"associations,omitempty"`
}
