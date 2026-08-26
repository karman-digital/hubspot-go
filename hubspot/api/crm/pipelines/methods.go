package pipelines

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	pipelinemodels "github.com/karman-digital/hubspot/hubspot/api/models/crm/pipelines"
)

func (service *PipelineService) GetPipelines(objectType string) (pipelinemodels.ListResponse, error) {
	var responseBody pipelinemodels.ListResponse
	response, err := service.SendRequest(http.MethodGet, fmt.Sprintf("/crm/v3/pipelines/%s", url.PathEscape(objectType)), nil)
	if err != nil {
		return responseBody, err
	}
	defer response.Body.Close()
	rawBody, err := io.ReadAll(response.Body)
	if err != nil {
		return responseBody, fmt.Errorf("error reading body: %s", err)
	}
	if response.StatusCode != http.StatusOK {
		return responseBody, fmt.Errorf("error returned by endpoint: %s", rawBody)
	}
	if err := json.Unmarshal(rawBody, &responseBody); err != nil {
		return responseBody, fmt.Errorf("error parsing body: %s", err)
	}
	return responseBody, nil
}
