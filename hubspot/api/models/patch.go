package husbpotmodels

type PatchBody struct {
	Properties Properties `json:"properties"`
}

type Properties map[string]any
