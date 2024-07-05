package shared

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
)

func HandleBatchResponse(resp *http.Response) (batchResp hubspotmodels.BatchResponse, err error) {
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
		return objResp, fmt.Errorf("error reading body: %s", err)
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
		return searchResp, fmt.Errorf("error reading body: %s", err)
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
		var errorResp hubspotmodels.ErrorResponseBody
		err := json.Unmarshal(rawBody, &errorResp)
		if err != nil {
			return nil, fmt.Errorf("error parsing error body: %s", err)
		}
		return rawBody, fmt.Errorf("error returned by endpoint: %+v", errorResp)
	}
	return rawBody, nil
}
