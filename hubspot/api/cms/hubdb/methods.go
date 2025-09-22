package hubdb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	hubdbmodels "github.com/karman-digital/hubspot/hubspot/api/models/cms/hubdb"
)

func (h *HubDBService) GetTableRow(tableId, rowId string) (hubdbmodels.HubDBRowResponse, error) {
	resp, err := h.SendRequest(http.MethodGet, fmt.Sprintf("/cms/v3/hubdb/tables/%s/rows/%s", tableId, rowId), nil)
	if err != nil {
		return hubdbmodels.HubDBRowResponse{}, fmt.Errorf("error making request: %s", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return hubdbmodels.HubDBRowResponse{}, fmt.Errorf("error reading response body: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		return hubdbmodels.HubDBRowResponse{}, fmt.Errorf("error getting table row: %s", string(body))
	}
	var hubDBRowResponse hubdbmodels.HubDBRowResponse
	err = json.Unmarshal(body, &hubDBRowResponse)
	if err != nil {
		return hubdbmodels.HubDBRowResponse{}, fmt.Errorf("error unmarshalling response body: %s", err)
	}
	return hubDBRowResponse, nil
}
