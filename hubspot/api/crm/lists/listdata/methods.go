package listdata

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	listsmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm/lists"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (l *ListDataService) SearchLists(body listsmodels.SearchListsBody) (listsmodels.SearchListsResponse, error) {
	reqUrl := "/crm/v3/lists/search"
	reqBody, err := json.Marshal(body)
	if err != nil {
		return listsmodels.SearchListsResponse{}, fmt.Errorf("error marshalling search body: %s", err)
	}
	resp, err := l.SendRequest(http.MethodPost, reqUrl, reqBody)
	if err != nil {
		return listsmodels.SearchListsResponse{}, fmt.Errorf("error creating request: %s", err)
	}
	return shared.HandleListsSearchResponse(resp)
}

func (l *ListDataService) GetLists(listIds []string, includeFilters bool) (listsmodels.ListsByIdResponse, error) {
	reqUrl := "/crm/v3/lists/"
	queryParams := url.Values{}
	for _, listId := range listIds {
		queryParams.Add("listIds", listId)
	}
	if includeFilters {
		queryParams.Add("includeFilters", "true")
	}
	if encoded := queryParams.Encode(); encoded != "" {
		reqUrl = fmt.Sprintf("%s?%s", reqUrl, encoded)
	}
	resp, err := l.SendRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return listsmodels.ListsByIdResponse{}, fmt.Errorf("error creating request: %s", err)
	}
	return shared.HandleListsByIdResponse(resp)
}
