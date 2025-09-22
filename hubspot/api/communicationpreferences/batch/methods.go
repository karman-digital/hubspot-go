package batchcommunicationpreferences

import (
	"encoding/json"
	"fmt"
	"net/http"

	communicationmodels "github.com/karman-digital/hubspot/hubspot/api/models/communicationpreferences"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (b *BatchCommunicationPreferencesService) BatchUpdateCommunicationPreferences(body communicationmodels.BatchCommunicationPreferencesPostBody) (communicationmodels.BatchCommunicationPreferencesResponse, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return communicationmodels.BatchCommunicationPreferencesResponse{}, fmt.Errorf("error marshalling post body: %s", err)
	}
	resp, err := b.SendRequest(http.MethodPost, "/communication-preferences/v4/statuses/batch/write", reqBody)
	if err != nil {
		return communicationmodels.BatchCommunicationPreferencesResponse{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleBatchCommunicationPreferencesResponse(resp)
}
