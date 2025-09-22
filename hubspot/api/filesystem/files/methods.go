package files

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	filesmodels "github.com/karman-digital/hubspot/hubspot/api/models/files"
	sharedmodels "github.com/karman-digital/hubspot/hubspot/api/models/shared"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (f *FilesService) ImportFileViaUrl(body filesmodels.FileImportBody) (filesmodels.FileUploadResult, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return filesmodels.FileUploadResult{}, err
	}
	resp, err := f.SendRequest("POST", "/files/v3/files/import-from-url/async", reqBody)
	if err != nil {
		return filesmodels.FileUploadResult{}, err
	}
	return f.handleFileUploadResponse(resp)
}

func (f *FilesService) handleFileUploadResponse(resp *http.Response) (filesmodels.FileUploadResult, error) {
	uploadResp, err := shared.HandleFileImportResponse(resp)
	if err != nil {
		return filesmodels.FileUploadResult{}, err
	}
	taskId := uploadResp.ID
	var fileUploadResult filesmodels.FileUploadResult
	var uploaded bool
	checkCount := 0
	maxChecks := 6
	for !uploaded && checkCount < maxChecks {
		if checkCount == 0 {
			time.Sleep(10 * time.Second)
		}
		resp, err = f.SendRequest("GET", fmt.Sprintf("/files/v3/files/import-from-url/async/tasks/%s/status", taskId), nil)
		if err != nil {
			return filesmodels.FileUploadResult{}, err
		}
		importStatusResp, err := shared.HandleFileImportStatusResponse(resp)
		if err != nil {
			return filesmodels.FileUploadResult{}, err
		}
		if importStatusResp.Status == "COMPLETE" {
			uploaded = true
			fileUploadResult = importStatusResp.Result
		} else if importStatusResp.Status == "FAILED" {
			return filesmodels.FileUploadResult{}, fmt.Errorf("file import failed")
		}
		if !uploaded {
			checkCount++
			if checkCount < maxChecks {
				time.Sleep(30 * time.Second)
			}
		}
	}
	if !uploaded {
		return filesmodels.FileUploadResult{}, fmt.Errorf("file import timed out after %d checks", maxChecks)
	}
	return fileUploadResult, nil
}

func (f *FilesService) GetSignedUrl(fileId string, signedUrlOptions ...filesmodels.SignedUrlOptions) (filesmodels.SignedUrlResponse, error) {
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
		return filesmodels.SignedUrlResponse{}, err
	}
	rawBody, err := shared.HandleBasicResponseCode(resp)
	if err != nil {
		return filesmodels.SignedUrlResponse{}, err
	}
	var signedUrlResponse filesmodels.SignedUrlResponse
	err = json.Unmarshal(rawBody, &signedUrlResponse)
	if err != nil {
		return filesmodels.SignedUrlResponse{}, err
	}
	return signedUrlResponse, nil
}

func (f *FilesService) UploadFile(fileName string, fileContent []byte, opts ...filesmodels.UploadFileOptions) (filesmodels.FileUploadResult, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return filesmodels.FileUploadResult{}, err
	}
	_, err = part.Write(fileContent)
	if err != nil {
		return filesmodels.FileUploadResult{}, err
	}
	writer.WriteField("fileName", fileName)
	if len(opts) > 0 {
		if opts[0].FolderId != "" {
			err = writer.WriteField("folderId", opts[0].FolderId)
			if err != nil {
				return filesmodels.FileUploadResult{}, err
			}
		}
		if opts[0].FolderPath != "" {
			err = writer.WriteField("folderPath", opts[0].FolderPath)
			if err != nil {
				return filesmodels.FileUploadResult{}, err
			}
		}
		options := map[string]string{}
		if opts[0].Options.Access != "" {
			options["access"] = opts[0].Options.Access
		}
		if opts[0].Options.Ttl != "" {
			options["ttl"] = opts[0].Options.Ttl
		}
		optionsJson, err := json.Marshal(options)
		if err != nil {
			return filesmodels.FileUploadResult{}, err
		}
		err = writer.WriteField("options", string(optionsJson))
		if err != nil {
			return filesmodels.FileUploadResult{}, err
		}
	}

	err = writer.Close()
	if err != nil {
		return filesmodels.FileUploadResult{}, err
	}
	req, err := retryablehttp.NewRequest("POST", "https://api.hubapi.com/files/v3/files", body)
	if err != nil {
		return filesmodels.FileUploadResult{}, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", f.Credentials.AccessToken()))
	resp, err := f.Credentials.Client().Do(req)
	if err != nil {
		return filesmodels.FileUploadResult{}, fmt.Errorf("error making request: %s", err)
	}
	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return filesmodels.FileUploadResult{}, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 201 {
		if resp.StatusCode == 404 {
			return filesmodels.FileUploadResult{}, fmt.Errorf("resource not found")
		}
		var errorResp sharedmodels.ErrorResponseBody
		err := json.Unmarshal(rawBody, &errorResp)
		if err != nil {
			return filesmodels.FileUploadResult{}, fmt.Errorf("error parsing error body: %s", err)
		}
		return filesmodels.FileUploadResult{}, fmt.Errorf("error returned by endpoint: %+v", errorResp)
	}
	var result filesmodels.FileUploadResult
	err = json.Unmarshal(rawBody, &result)
	if err != nil {
		return filesmodels.FileUploadResult{}, fmt.Errorf("error parsing body: %s", err)
	}
	return result, nil
}
