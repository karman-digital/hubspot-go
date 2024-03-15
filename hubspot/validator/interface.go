package hubspotvalidator

type Validator interface {
	ValidateHubspotToken() error
}
