package shared

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	communicationmodels "github.com/karman-digital/hubspot/hubspot/api/models/communicationpreferences"
	crmmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm"
	listsmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm/lists"
	filesmodels "github.com/karman-digital/hubspot/hubspot/api/models/files"
	campaignassetsmodels "github.com/karman-digital/hubspot/hubspot/api/models/marketing/campaigns/assets"
	campaignmodels "github.com/karman-digital/hubspot/hubspot/api/models/marketing/campaigns"
	emailanalyticscampaignmodels "github.com/karman-digital/hubspot/hubspot/api/models/marketing/emailanalytics/campaigns"
	emailanalyticseventsmodels "github.com/karman-digital/hubspot/hubspot/api/models/marketing/emailanalytics/events"
	emailmodels "github.com/karman-digital/hubspot/hubspot/api/models/marketing/emails"
	sharedmodels "github.com/karman-digital/hubspot/hubspot/api/models/shared"
	usermodels "github.com/karman-digital/hubspot/hubspot/api/models/users"
)

func HandleBatchResponse(resp *http.Response, method string) (batchResp crmmodels.BatchResponse, err error) {
	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return batchResp, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 201 && resp.StatusCode != 200 && resp.StatusCode != 207 {
		var errorResp sharedmodels.ErrorResponseBody
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

func HandleBatchCommunicationPreferencesResponse(resp *http.Response) (batchResp communicationmodels.BatchCommunicationPreferencesResponse, err error) {
	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return batchResp, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 201 && resp.StatusCode != 200 && resp.StatusCode != 207 {
		var errorResp sharedmodels.ErrorResponseBody
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

func HandleBatchResponseCodes(errorResp sharedmodels.ErrorResponseBody, statusCode int) error {
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

func HandleUserResponse(resp *http.Response) (userResp usermodels.UserBody, err error) {
	rawBody, err := handleBasicResponseCode(resp)
	if err != nil {
		return userResp, err
	}
	err = json.Unmarshal(rawBody, &userResp)
	if err != nil {
		return userResp, fmt.Errorf("error parsing body: %s", err)
	}
	return userResp, nil
}

func HandleResponse(resp *http.Response) (objResp crmmodels.Result, err error) {
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

func HandleError(resp *http.Response, returnedErr error) (objResp crmmodels.Result, err error) {
	if _, err = handleBasicResponseCode(resp); err != nil {
		return objResp, err
	}
	return objResp, nil
}

func HandleListResponse(resp *http.Response) (listResp crmmodels.ListResponse, err error) {
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

func HandleCreateResponse(resp *http.Response) (objResp crmmodels.Result, err error) {
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

func HandleSearchResponse(resp *http.Response) (searchResp crmmodels.SearchResponse, err error) {
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

func HandleBasicResponseCode(resp *http.Response) (rawBody []byte, err error) {
	return handleBasicResponseCode(resp)
}

func handleBasicResponseCode(resp *http.Response) (rawBody []byte, err error) {
	rawBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != http.StatusAccepted && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return rawBody, ErrResourceNotFound
		}
		var errorResp sharedmodels.ErrorResponseBody
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
		var errorResp sharedmodels.ErrorResponseBody
		err := json.Unmarshal(rawBody, &errorResp)
		if err != nil {
			return nil, fmt.Errorf("error parsing error body: %s", err)
		}
		return rawBody, fmt.Errorf("error returned by endpoint: %+v", errorResp)
	}
	return rawBody, nil
}

func HandleFileImportResponse(resp *http.Response) (fileImportResp filesmodels.FileImportResponse, err error) {
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

func HandleFileImportStatusResponse(resp *http.Response) (fileImportStatusResp filesmodels.FileImportStatusResponse, err error) {
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

func HandleListsSearchResponse(resp *http.Response) (searchResp listsmodels.SearchListsResponse, err error) {
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

func HandleListMembershipsResponse(resp *http.Response) (membershipsResp listsmodels.ListMembershipsResponse, err error) {
	rawBody, err := handleBasicResponseCode(resp)
	if err != nil {
		return membershipsResp, err
	}
	err = json.Unmarshal(rawBody, &membershipsResp)
	if err != nil {
		return membershipsResp, fmt.Errorf("error parsing body: %s", err)
	}
	return membershipsResp, nil
}

func HandleListsByIdResponse(resp *http.Response) (listsResp listsmodels.ListsByIdResponse, err error) {
	rawBody, err := handleBasicResponseCode(resp)
	if err != nil {
		return listsResp, err
	}
	err = json.Unmarshal(rawBody, &listsResp)
	if err != nil {
		return listsResp, fmt.Errorf("error parsing body: %s", err)
	}
	return listsResp, nil
}

func HandleCampaignsResponse(resp *http.Response) (campaignsResp campaignmodels.CampaignsResponse, err error) {
	rawBody, err := handleBasicResponseCode(resp)
	if err != nil {
		return campaignsResp, err
	}
	err = json.Unmarshal(rawBody, &campaignsResp)
	if err != nil {
		return campaignsResp, fmt.Errorf("error parsing body: %s", err)
	}
	return campaignsResp, nil
}

func HandleCampaignResponse(resp *http.Response) (campaignResp campaignmodels.Campaign, err error) {
	rawBody, err := handleBasicResponseCode(resp)
	if err != nil {
		return campaignResp, err
	}
	err = json.Unmarshal(rawBody, &campaignResp)
	if err != nil {
		return campaignResp, fmt.Errorf("error parsing body: %s", err)
	}
	return campaignResp, nil
}

func HandleCampaignAssetsResponse(resp *http.Response) (assetsResp campaignassetsmodels.CampaignAssetsResponse, err error) {
	rawBody, err := handleBasicResponseCode(resp)
	if err != nil {
		return assetsResp, err
	}
	err = json.Unmarshal(rawBody, &assetsResp)
	if err != nil {
		return assetsResp, fmt.Errorf("error parsing body: %s", err)
	}
	return assetsResp, nil
}

func HandleMarketingEmailResponse(resp *http.Response) (emailResp emailmodels.MarketingEmail, err error) {
	rawBody, err := handleBasicResponseCode(resp)
	if err != nil {
		return emailResp, err
	}
	err = json.Unmarshal(rawBody, &emailResp)
	if err != nil {
		return emailResp, fmt.Errorf("error parsing body: %s", err)
	}
	return emailResp, nil
}

func HandleEmailCampaignsResponse(resp *http.Response) (campaignsResp emailanalyticscampaignmodels.EmailCampaignsResponse, err error) {
	rawBody, err := handleBasicResponseCode(resp)
	if err != nil {
		return campaignsResp, err
	}
	err = json.Unmarshal(rawBody, &campaignsResp)
	if err != nil {
		return campaignsResp, fmt.Errorf("error parsing body: %s", err)
	}
	return campaignsResp, nil
}

func HandleEmailCampaignDetailResponse(resp *http.Response) (campaignDetailResp emailanalyticscampaignmodels.EmailCampaignDetail, err error) {
	rawBody, err := handleBasicResponseCode(resp)
	if err != nil {
		return campaignDetailResp, err
	}
	err = json.Unmarshal(rawBody, &campaignDetailResp)
	if err != nil {
		return campaignDetailResp, fmt.Errorf("error parsing body: %s", err)
	}
	return campaignDetailResp, nil
}

func HandleEmailEventsResponse(resp *http.Response) (eventsResp emailanalyticseventsmodels.EmailEventsResponse, err error) {
	rawBody, err := handleBasicResponseCode(resp)
	if err != nil {
		return eventsResp, err
	}
	err = json.Unmarshal(rawBody, &eventsResp)
	if err != nil {
		return eventsResp, fmt.Errorf("error parsing body: %s", err)
	}
	return eventsResp, nil
}
