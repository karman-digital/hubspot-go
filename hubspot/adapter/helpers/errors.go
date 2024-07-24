package hshelpers

import (
	"encoding/json"
	"fmt"

	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
)

func ParseValidationErrors(errResponse hubspotmodels.ErrorResponseBody) (hubspotmodels.ParsedMessage, error) {
	var parsedMessage hubspotmodels.ParsedMessage
	err := json.Unmarshal([]byte(fmt.Sprintf(`{"Property values were not valid":%s}`, errResponse.Message[len("Property values were not valid: "):])), &parsedMessage)
	if err != nil {
		return parsedMessage, errResponse
	}
	return parsedMessage, nil
}

func ParseErrorResponse(responseBody []byte) (hubspotmodels.ErrorResponseBody, error) {
	var errResponse hubspotmodels.ErrorResponseBody
	if err := json.Unmarshal(responseBody, &errResponse); err != nil {
		return errResponse, fmt.Errorf("error parsing error response: %s", err)
	}
	return errResponse, nil
}

func PasrseErrorAsErrorResponse(err error) (hubspotmodels.ErrorResponseBody, error) {
	if parsedError, ok := err.(hubspotmodels.ErrorResponseBody); ok {
		return parsedError, nil
	}
	return hubspotmodels.ErrorResponseBody{}, err
}

func ParseErrorAsValidationErrors(err error) (hubspotmodels.ParsedMessage, error) {
	errResp, err := PasrseErrorAsErrorResponse(err)
	if err != nil {
		return hubspotmodels.ParsedMessage{}, err
	}
	return ParseValidationErrors(errResp)
}

func IsEmailInvalidError(err hubspotmodels.ErrorResponseBody) bool {
	if err.Category == "VALIDATION_ERROR" {
		message, err := ParseValidationErrors(err)
		if err != nil {
			return false
		}
		return message.EmailInvalidError()
	}
}
