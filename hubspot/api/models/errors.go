package hubspotmodels

type ErrorResponseBody struct {
	Status        string `json:"status"`
	Message       string `json:"message"`
	Category      string `json:"category"`
	CorrelationId string `json:"correlationId"`
}
