package files

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
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

func (f *FilesService) GetSignedUrl(fileId string, signedUrlOptions ...hubspotmodels.SignedUrlOptions) (hubspotmodels.SignedUrlResponse, error) {
	endpoint := fmt.Sprintf("/files/v3/files/%s/signed-url", fileId)
	if len(signedUrlOptions) > 0 {
		queryParams := url.Values{}
		if signedUrlOptions[0].ExpirationSeconds > 0 {
			queryParams.Add("expirationSeconds", strconv.FormatInt(signedUrlOptions[0].ExpirationSeconds, 10))
		}
		if signedUrlOptions[0].Size != "" {
			queryParams.Add("size", signedUrlOptions[0].Size)
		}
		if signedUrlOptions[0].Upscale && signedUrlOptions[0].Size != "" {
			queryParams.Add("upscale", "true")
		}
		endpoint += "?" + queryParams.Encode()
	}
	resp, err := f.SendRequest("GET", endpoint, nil)
	if err != nil {
		return hubspotmodels.SignedUrlResponse{}, err
	}
	rawBody, err := shared.HandleBasicResponseCode(resp)
	if err != nil {
		return hubspotmodels.SignedUrlResponse{}, err
	}
	var signedUrlResponse hubspotmodels.SignedUrlResponse
	err = json.Unmarshal(rawBody, &signedUrlResponse)
	if err != nil {
		return hubspotmodels.SignedUrlResponse{}, err
	}
	return signedUrlResponse, nil
}
