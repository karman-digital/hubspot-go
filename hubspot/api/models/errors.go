package hubspotmodels

type ErrorResponseBody struct {
	Status        string `json:"status"`
	Message       string `json:"message"`
	Category      string `json:"category"`
	CorrelationId string `json:"correlationId"`
}

type NestedMessage struct {
	IsValid bool   `json:"isValid"`
	Message string `json:"message"`
	Error   string `json:"error"`
	Name    string `json:"name"`
}

type ParsedMessage struct {
	PropertyValues []NestedMessage `json:"Property values were not valid"`
}

func (e ErrorResponseBody) Error() string {
	return e.Message
}
