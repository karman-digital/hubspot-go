package hshelpers

import (
	"encoding/json"

	webhookmodels "github.com/karman-digital/hubspot/hubspot/api/models/webhooks"
)

func ConvertBodyBytesToAccountDataPayload(bodyBytes []byte) (webhookmodels.AccountDataPayload, error) {
	payload := webhookmodels.AccountDataPayload{}
	if err := json.Unmarshal(bodyBytes, &payload); err != nil {
		return payload, err
	}
	return payload, nil
}
