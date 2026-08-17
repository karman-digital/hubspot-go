package schemas

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	schemasmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm/schemas"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (s *SchemaService) CreateAssociationDefinition(objectType string, body schemasmodels.AssociationDefinitionBody) error {
	requestBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("error marshalling association definition: %w", err)
	}
	response, err := s.SendRequest(http.MethodPost, fmt.Sprintf("/crm/v3/schemas/%s/associations", objectType), requestBody)
	if err != nil {
		return fmt.Errorf("error creating association definition: %w", err)
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusConflict {
		return shared.ErrResourceAlreadyExists
	}
	if response.StatusCode != http.StatusCreated {
		responseBody, readErr := io.ReadAll(response.Body)
		if readErr != nil {
			return fmt.Errorf("error reading association definition response: %w", readErr)
		}
		return fmt.Errorf("error returned by association definition endpoint. status code: %s, error: %s", response.Status, responseBody)
	}
	return nil
}
