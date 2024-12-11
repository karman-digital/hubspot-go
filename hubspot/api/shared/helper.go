package shared

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
)

func HandleBatchResponse(resp *http.Response, method string) (batchResp hubspotmodels.BatchResponse, err error) {
	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return batchResp, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 201 && resp.StatusCode != 200 && resp.StatusCode != 207 {
		var errorResp hubspotmodels.ErrorResponseBody
		err := json.Unmarshal(rawBody, &errorResp)
		if err != nil {
			return batchResp, fmt.Errorf("error parsing error body: %s", err)
		}
		return batchResp, HandleBatchResponseCodes(errorResp, resp.StatusCode)
	}
	err = json.Unmarshal(rawBody, &batchResp)
	if err != nil {
		return batchResp, fmt.Errorf("error parsing body: %s", err)
	}
	if resp.StatusCode == 207 {
		if method == http.MethodGet {
			return batchResp, ErrBatchGet
		}
		return batchResp, ErrBatchCreate
	}
	return batchResp, nil
}

func HandleBatchCommunicationPreferencesResponse(resp *http.Response) (batchResp hubspotmodels.BatchCommunicationPreferencesResponse, err error) {
	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return batchResp, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 201 && resp.StatusCode != 200 && resp.StatusCode != 207 {
		var errorResp hubspotmodels.ErrorResponseBody
		err := json.Unmarshal(rawBody, &errorResp)
		if err != nil {
			return batchResp, fmt.Errorf("error parsing error body: %s", err)
		}
		return batchResp, HandleBatchResponseCodes(errorResp, resp.StatusCode)
	}
	err = json.Unmarshal(rawBody, &batchResp)
	if err != nil {
		return batchResp, fmt.Errorf("error parsing body: %s", err)
	}
	if resp.StatusCode == 207 {
		return batchResp, ErrBatchCreate
	}
	return batchResp, nil
}

func HandleBatchResponseCodes(errorResp hubspotmodels.ErrorResponseBody, statusCode int) error {
	switch statusCode {
	case 200:
		return nil
	case 400:
		if errorResp.Category == "VALIDATION" {
			fmt.Printf("validation error returned by endpoint: %s", errorResp.Message)
			return ErrPropertyValidation
		}
		return fmt.Errorf("error returned by endpoint: %s", errorResp.Message)
	case 409:
		fmt.Printf("object already exists: %s", errorResp.Message)
		return ErrObjectAlreadyExists
	default:
		return ErrApiCall
	}
}

func HandleResponse(resp *http.Response) (objResp hubspotmodels.Result, err error) {
	rawBody, err := handleBasicResponseCode(resp)
	if err != nil {
		return objResp, err
	}
	err = json.Unmarshal(rawBody, &objResp)
	if err != nil {
		return objResp, fmt.Errorf("error parsing body: %s", err)
	}
	return objResp, nil
}

func HandleError(resp *http.Response, returnedErr error) (objResp hubspotmodels.Result, err error) {
	if _, err = handleBasicResponseCode(resp); err != nil {
		return objResp, err
	}
	return objResp, nil
}

func HandleListResponse(resp *http.Response) (listResp hubspotmodels.ListResponse, err error) {
	rawBody, err := handleBasicResponseCode(resp)
	if err != nil {
		return listResp, err
	}
	err = json.Unmarshal(rawBody, &listResp)
	if err != nil {
		return listResp, fmt.Errorf("error parsing body: %s", err)
	}
	return listResp, nil
}

func HandleCreateResponse(resp *http.Response) (objResp hubspotmodels.Result, err error) {
	rawBody, err := handleCustomResponseCode(resp, http.StatusCreated)
	if err != nil {
		return objResp, err
	}
	err = json.Unmarshal(rawBody, &objResp)
	if err != nil {
		return objResp, fmt.Errorf("error parsing body: %s", err)
	}
	return objResp, nil
}

func HandleSearchResponse(resp *http.Response) (searchResp hubspotmodels.SearchResponse, err error) {
	rawBody, err := handleBasicResponseCode(resp)
	if err != nil {
		return searchResp, err
	}
	err = json.Unmarshal(rawBody, &searchResp)
	if err != nil {
		return searchResp, fmt.Errorf("error parsing body: %s", err)
	}
	return searchResp, nil
}

func handleBasicResponseCode(resp *http.Response) (rawBody []byte, err error) {
	rawBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 404 {
			return rawBody, ErrResourceNotFound
		}
		var errorResp hubspotmodels.ErrorResponseBody
		err := json.Unmarshal(rawBody, &errorResp)
		if err != nil {
			return nil, fmt.Errorf("error parsing error body: %s", err)
		}
		return rawBody, fmt.Errorf("error returned by endpoint: %+v", errorResp)
	}
	return rawBody, nil
}

func handleCustomResponseCode(resp *http.Response, code int) (rawBody []byte, err error) {
	rawBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != code {
		if resp.StatusCode == 404 {
			return rawBody, ErrResourceNotFound
		}
		var errorResp hubspotmodels.ErrorResponseBody
		err := json.Unmarshal(rawBody, &errorResp)
		if err != nil {
			return nil, fmt.Errorf("error parsing error body: %s", err)
		}
		return rawBody, fmt.Errorf("error returned by endpoint: %+v", errorResp)
	}
	return rawBody, nil
}

func HandleFileImportResponse(resp *http.Response) (fileImportResp hubspotmodels.FileImportResponse, err error) {
	rawBody, err := handleCustomResponseCode(resp, http.StatusAccepted)
	if err != nil {
		return fileImportResp, err
	}
	err = json.Unmarshal(rawBody, &fileImportResp)
	if err != nil {
		return fileImportResp, fmt.Errorf("error parsing body: %s", err)
	}
	return fileImportResp, nil
}

func HandleDeleteResponse(resp *http.Response) (err error) {
	_, err = handleCustomResponseCode(resp, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

func HandleFileImportStatusResponse(resp *http.Response) (fileImportStatusResp hubspotmodels.FileImportStatusResponse, err error) {
	rawBody, err := handleCustomResponseCode(resp, http.StatusOK)
	if err != nil {
		return fileImportStatusResp, err
	}
	err = json.Unmarshal(rawBody, &fileImportStatusResp)
	if err != nil {
		return fileImportStatusResp, fmt.Errorf("error parsing body: %s", err)
	}
	return fileImportStatusResp, nil
}
