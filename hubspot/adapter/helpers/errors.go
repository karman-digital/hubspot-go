package hshelpers

import (
	"encoding/json"
	"fmt"
	"log"

	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
)

func ParseValidationErrors(errResponse hubspotmodels.ErrorResponseBody) (hubspotmodels.ParsedMessage, error) {
	var parsedMessage hubspotmodels.ParsedMessage
	err := json.Unmarshal([]byte(fmt.Sprintf(`{"Property values were not valid":%s}`, errResponse.Message[len("Property values were not valid: "):])), &parsedMessage)
	if err != nil {
		log.Fatalf("Error unmarshaling nested JSON: %v", err)
	}
	return parsedMessage, nil
}

func ParseErrorResponse(responseBody []byte) (hubspotmodels.ErrorResponseBody, error) {
	var errResponse hubspotmodels.ErrorResponseBody
	if err := json.Unmarshal(responseBody, &errResponse); err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}
	return errResponse, nil
}
