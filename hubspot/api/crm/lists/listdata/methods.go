package listdata

import (
	"encoding/json"
	"fmt"
	"net/http"

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

