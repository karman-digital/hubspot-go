package hshelpers

import (
	"encoding/json"
	"fmt"

	sharedmodels "github.com/karman-digital/hubspot/hubspot/api/models/shared"
)

func ParseValidationErrors(errResponse sharedmodels.ErrorResponseBody) (sharedmodels.ParsedMessage, error) {
	var parsedMessage sharedmodels.ParsedMessage
	err := json.Unmarshal([]byte(fmt.Sprintf(`{"Property values were not valid":%s}`, errResponse.Message[len("Property values were not valid: "):])), &parsedMessage)
	if err != nil {
		return parsedMessage, errResponse
	}
	return parsedMessage, nil
}

func ParseErrorResponse(responseBody []byte) (sharedmodels.ErrorResponseBody, error) {
	var errResponse sharedmodels.ErrorResponseBody
	if err := json.Unmarshal(responseBody, &errResponse); err != nil {
		return errResponse, fmt.Errorf("error parsing error response: %s", err)
	}
	return errResponse, nil
}

func PasrseErrorAsErrorResponse(err error) (sharedmodels.ErrorResponseBody, error) {
	if parsedError, ok := err.(sharedmodels.ErrorResponseBody); ok {
		return parsedError, nil
	}
	return sharedmodels.ErrorResponseBody{}, err
}

func ParseErrorAsValidationErrors(err error) (sharedmodels.ParsedMessage, error) {
	errResp, err := PasrseErrorAsErrorResponse(err)
	if err != nil {
		return sharedmodels.ParsedMessage{}, err
	}
	return ParseValidationErrors(errResp)
}

func IsEmailInvalidError(err sharedmodels.ErrorResponseBody) bool {
	if err.Category == "VALIDATION_ERROR" {
		message, err := ParseValidationErrors(err)
		if err != nil {
			return false
		}
		return message.EmailInvalidError()
	}
	return false
}
