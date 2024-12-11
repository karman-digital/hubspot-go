package files

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (f *FilesService) ImportFileViaUrl(body hubspotmodels.FileImportBody) (hubspotmodels.FileImportStatusResponse, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return hubspotmodels.FileImportStatusResponse{}, err
	}
	resp, err := f.SendRequest("POST", "/files/v3/files/import-from-url/async", reqBody)
	if err != nil {
		return hubspotmodels.FileImportStatusResponse{}, err
	}
	return f.handleFileUploadResponse(resp)
}

func (f *FilesService) handleFileUploadResponse(resp *http.Response) (hubspotmodels.FileImportStatusResponse, error) {
	uploadResp, err := shared.HandleFileImportResponse(resp)
	if err != nil {
		return hubspotmodels.FileImportStatusResponse{}, err
	}
	statusLink := uploadResp.Links.Status
	var importStatusResp hubspotmodels.FileImportStatusResponse
	var uploaded bool
	checkCount := 0
	maxChecks := 3
	for !uploaded && checkCount < maxChecks {
		resp, err = f.Client().Get(statusLink)
		if err != nil {
			return hubspotmodels.FileImportStatusResponse{}, err
		}
		importStatusResp, err = shared.HandleFileImportStatusResponse(resp)
		if err != nil {
			return hubspotmodels.FileImportStatusResponse{}, err
		}
		if importStatusResp.Status == "COMPLETE" {
			uploaded = true
		} else if importStatusResp.Status == "FAILED" {
			return hubspotmodels.FileImportStatusResponse{}, fmt.Errorf("file import failed")
		}
		if !uploaded {
			checkCount++
			if checkCount < maxChecks {
				time.Sleep(2 * time.Minute)
			}
		}
	}
	if !uploaded {
		return hubspotmodels.FileImportStatusResponse{}, fmt.Errorf("file import timed out after %d checks", maxChecks)
	}
	return importStatusResp, nil
}
