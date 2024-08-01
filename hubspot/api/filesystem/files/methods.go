package files

import (
	"encoding/json"

	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (f *FilesService) ImportFileViaUrl(body hubspotmodels.FileImportBody) (hubspotmodels.FileImportResponse, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return hubspotmodels.FileImportResponse{}, err
	}
	resp, err := f.SendRequest("POST", "/files/v3/files/import-from-url/async", reqBody)
	if err != nil {
		return hubspotmodels.FileImportResponse{}, err
	}
	return shared.HandleFileImportResponse(resp)
}
