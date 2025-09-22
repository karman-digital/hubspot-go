package graphqlmodels

type GraphQLResponse struct {
	Extensions map[string]interface{}   `json:"extensions"`
	Data       map[string]interface{}   `json:"data"`
	Errors     []map[string]interface{} `json:"errors"`
}
