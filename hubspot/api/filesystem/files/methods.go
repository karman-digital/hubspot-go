package files

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (f *FilesService) ImportFileViaUrl(body hubspotmodels.FileImportBody) (hubspotmodels.FileUploadResult, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return hubspotmodels.FileUploadResult{}, err
	}
	resp, err := f.SendRequest("POST", "/files/v3/files/import-from-url/async", reqBody)
	if err != nil {
		return hubspotmodels.FileUploadResult{}, err
	}
	return f.handleFileUploadResponse(resp)
}

func (f *FilesService) handleFileUploadResponse(resp *http.Response) (hubspotmodels.FileUploadResult, error) {
	uploadResp, err := shared.HandleFileImportResponse(resp)
	if err != nil {
		return hubspotmodels.FileUploadResult{}, err
	}
	taskId := uploadResp.ID
	var fileUploadResult hubspotmodels.FileUploadResult
	var uploaded bool
	checkCount := 0
	maxChecks := 6
	for !uploaded && checkCount < maxChecks {
		time.Sleep(1 * time.Minute)
		resp, err = f.SendRequest("GET", fmt.Sprintf("/files/v3/files/import-from-url/async/tasks/%s/status", taskId), nil)
		if err != nil {
			return hubspotmodels.FileUploadResult{}, err
		}
		importStatusResp, err := shared.HandleFileImportStatusResponse(resp)
		if err != nil {
			return hubspotmodels.FileUploadResult{}, err
		}
		if importStatusResp.Status == "COMPLETE" {
			uploaded = true
			fileUploadResult = importStatusResp.Result
		} else if importStatusResp.Status == "FAILED" {
			return hubspotmodels.FileUploadResult{}, fmt.Errorf("file import failed")
		}
		if !uploaded {
			checkCount++
			if checkCount < maxChecks {
				time.Sleep(30 * time.Second)
			}
		}
	}
	if !uploaded {
		return hubspotmodels.FileUploadResult{}, fmt.Errorf("file import timed out after %d checks", maxChecks)
	}
	return fileUploadResult, nil
}
