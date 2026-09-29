package schemas

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	schemasmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm/schemas"
)

func (s *SchemaService) GetSchema(objectType string) (schemasmodels.Schema, error) {
	response, err := s.SendRequest(http.MethodGet, fmt.Sprintf("/crm-object-schemas/2026-09/schemas/%s", url.PathEscape(objectType)), nil)
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
