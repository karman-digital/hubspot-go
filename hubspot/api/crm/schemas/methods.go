package schemas

import (
	"encoding/json"
	"fmt"
	"net/http"

	schemasmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm/schemas"
)

func (s *SchemaService) GetSchema(objectType string) (schemasmodels.Schema, error) {
	response, err := s.SendRequest(http.MethodGet, fmt.Sprintf("/crm/v3/schemas/%s", objectType), nil)
	if err != nil {
		return schemasmodels.Schema{}, fmt.Errorf("error getting schema: %w", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return schemasmodels.Schema{}, fmt.Errorf("error returned by schema endpoint: %s", response.Status)
	}
	var schema schemasmodels.Schema
	if err := json.NewDecoder(response.Body).Decode(&schema); err != nil {
		return schemasmodels.Schema{}, fmt.Errorf("error decoding schema response: %w", err)
	}
	return schema, nil
}
