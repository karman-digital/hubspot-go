package blogtags

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
)

func (s *BlogTagsService) GetBatchBlogTags(opts hubspotmodels.BlogTagsBatchInput) (hubspotmodels.BatchBlogTagResponse, error) {
	reqBody, err := json.Marshal(opts)
	if err != nil {
		return hubspotmodels.BatchBlogTagResponse{}, fmt.Errorf("error marshalling post body: %s", err)
	}
	resp, err := s.SendRequest(http.MethodPost, "/cms/v3/blogs/tags/batch/read", reqBody)
	if err != nil {
		return hubspotmodels.BatchBlogTagResponse{}, fmt.Errorf("error making request: %s", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return hubspotmodels.BatchBlogTagResponse{}, fmt.Errorf("error reading response body: %s", err)
	}
	var batchBlogTagResponse hubspotmodels.BatchBlogTagResponse
	err = json.Unmarshal(body, &batchBlogTagResponse)
	if err != nil {
		return hubspotmodels.BatchBlogTagResponse{}, fmt.Errorf("error unmarshalling response body: %s", err)
	}
	return batchBlogTagResponse, nil
}
