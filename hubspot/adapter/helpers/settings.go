package hshelpers

import (
	"encoding/json"

	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
)

func ConvertBodyBytesToAccountDataPayload(bodyBytes []byte) (hubspotmodels.AccountDataPayload, error) {
	payload := hubspotmodels.AccountDataPayload{}
	if err := json.Unmarshal(bodyBytes, &payload); err != nil {
		return payload, err
	}
	return payload, nil
}
