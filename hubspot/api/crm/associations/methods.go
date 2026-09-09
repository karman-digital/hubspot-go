package associations

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	crmmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm"
	associationsmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm/associations"
	sharedmodels "github.com/karman-digital/hubspot/hubspot/api/models/shared"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

const batchGetAssociationsLimit = 1000

func (c *AssociationService) CreateDefaultAssociation(fromObject, toObject string, fromId, toId int) (crmmodels.BatchResponse, error) {
	var associationResp crmmodels.BatchResponse
	resp, err := c.SendRequest(
		http.MethodPut,
		fmt.Sprintf("/crm/v4/objects/%s/%d/associations/default/%s/%d", fromObject, fromId, toObject, toId),
		nil,
	)
	if err != nil {
		return associationResp, err
	}
	defer resp.Body.Close()
	associationRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return associationResp, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 200 {
		return associationResp, fmt.Errorf("error returned by endpoint: %s", associationRawBody)
	}
	err = json.Unmarshal(associationRawBody, &associationResp)
	if err != nil {
		return associationResp, fmt.Errorf("error parsing body: %s", err)
	}

	return associationResp, nil
}

func (c *AssociationService) GetAssociations(fromObject, toObject string, id int, opts ...sharedmodels.GetOptions) (associationsmodels.AssociationGetResponse, error) {
	var association associationsmodels.AssociationGetResponse
	resp, err := c.SendRequest(
		http.MethodGet,
		fmt.Sprintf("/crm/v4/objects/%s/%d/associations/%s", fromObject, id, toObject),
		nil,
		opts...,
	)
	if err != nil {
		return association, err
	}
	defer resp.Body.Close()
	associationRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return association, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 200 {
		return association, fmt.Errorf("error returned by endpoint: %s", associationRawBody)
	}
	err = json.Unmarshal(associationRawBody, &association)
	if err != nil {
		return association, fmt.Errorf("error parsing body: %s", err)
	}
	return association, nil
}

func (c *AssociationService) BatchCreateDefaultAssociations(fromObject, toObject string, associations associationsmodels.BatchCreateDefaultAssociationsBody) (crmmodels.BatchResponse, error) {
	var associationResp crmmodels.BatchResponse
	reqBody, err := json.Marshal(associations)
	if err != nil {
		return associationResp, fmt.Errorf("error marshalling post body: %s", err)
	}
	resp, err := c.SendRequest(
		http.MethodPost,
		fmt.Sprintf("/crm/v4/associations/%s/%s/batch/associate/default", fromObject, toObject),
		reqBody,
	)
	if err != nil {
		return associationResp, err
	}
	defer resp.Body.Close()
	associationRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return associationResp, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 200 {
		return associationResp, fmt.Errorf("error returned by endpoint: %s", associationRawBody)
	}
	err = json.Unmarshal(associationRawBody, &associationResp)
	if err != nil {
		return associationResp, fmt.Errorf("error parsing body: %s", err)
	}
	return associationResp, nil
}

func (c *AssociationService) BatchGetAssociations(fromObject, toObject string, body associationsmodels.BatchGetAssociationsBody) (associationsmodels.BatchAssociationGetResponse, error) {
	var batchResp associationsmodels.BatchAssociationGetResponse
	reqBody, err := json.Marshal(body)
	if err != nil {
		return batchResp, fmt.Errorf("error marshalling post body: %s", err)
	}
	resp, err := c.SendRequest(
		http.MethodPost,
		fmt.Sprintf("/crm/v4/associations/%s/%s/batch/read", fromObject, toObject),
		reqBody,
	)
	if err != nil {
		return batchResp, err
	}
	defer resp.Body.Close()
	contactRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return batchResp, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 200 && resp.StatusCode != 207 {
		var errorResp sharedmodels.ErrorResponseBody
		err := json.Unmarshal(contactRawBody, &errorResp)
		if err != nil {
			return batchResp, fmt.Errorf("error parsing error body: %s", err)
		}
		return batchResp, shared.HandleBatchResponseCodes(errorResp, resp.StatusCode)
	}
	err = json.Unmarshal(contactRawBody, &batchResp)
	if err != nil {
		return batchResp, fmt.Errorf("error parsing body: %s", err)
	}
	if resp.StatusCode == 207 {
		return batchResp, shared.ErrBatchGet
	}
	return batchResp, nil
}

func (c *AssociationService) BatchGetAllAssociations(fromObject, toObject string, sourceIDs []string) ([]associationsmodels.BatchAssociationResult, error) {
	if len(sourceIDs) == 0 {
		return []associationsmodels.BatchAssociationResult{}, nil
	}
	if len(sourceIDs) > batchGetAssociationsLimit {
		return nil, fmt.Errorf("association batch read accepts at most %d source IDs", batchGetAssociationsLimit)
	}

	resultIndexes := make(map[string]int, len(sourceIDs))
	seenCursors := make(map[string]map[string]struct{}, len(sourceIDs))
	seenTargets := make(map[string]map[int]struct{}, len(sourceIDs))
	results := make([]associationsmodels.BatchAssociationResult, len(sourceIDs))
	pending := make([]associationsmodels.BatchGetAssociationsInput, len(sourceIDs))
	for index, sourceID := range sourceIDs {
		if strings.TrimSpace(sourceID) == "" {
			return nil, fmt.Errorf("association batch read source ID is required")
		}
		if _, duplicate := resultIndexes[sourceID]; duplicate {
			return nil, fmt.Errorf("association batch read has duplicate source ID %q", sourceID)
		}
		resultIndexes[sourceID] = index
		seenCursors[sourceID] = map[string]struct{}{"": {}}
		seenTargets[sourceID] = make(map[int]struct{})
		results[index] = associationsmodels.BatchAssociationResult{
			From: associationsmodels.From{ID: sourceID},
			To:   make([]associationsmodels.ToItem, 0),
		}
		pending[index] = associationsmodels.BatchGetAssociationsInput{Id: sourceID}
	}

	for len(pending) > 0 {
		response, err := c.BatchGetAssociations(fromObject, toObject, associationsmodels.BatchGetAssociationsBody{Inputs: pending})
		if err != nil {
			return nil, fmt.Errorf("read complete associations: %w", err)
		}
		if response.NumErrors != 0 {
			return nil, fmt.Errorf("association batch read reported %d errors", response.NumErrors)
		}

		expected := make(map[string]struct{}, len(pending))
		for _, input := range pending {
			expected[input.Id] = struct{}{}
		}
		returned := make(map[string]struct{}, len(response.Results))
		next := make([]associationsmodels.BatchGetAssociationsInput, 0)
		for _, source := range response.Results {
			sourceID := source.From.ID
			if _, requested := expected[sourceID]; !requested {
				return nil, fmt.Errorf("association batch read returned unknown source ID %q", sourceID)
			}
			if _, duplicate := returned[sourceID]; duplicate {
				return nil, fmt.Errorf("association batch read returned duplicate source ID %q", sourceID)
			}
			returned[sourceID] = struct{}{}

			index := resultIndexes[sourceID]
			for _, target := range source.To {
				if _, duplicate := seenTargets[sourceID][target.ToObjectId]; duplicate {
					return nil, fmt.Errorf("association batch read returned duplicate target %d for source ID %q", target.ToObjectId, sourceID)
				}
				seenTargets[sourceID][target.ToObjectId] = struct{}{}
				results[index].To = append(results[index].To, target)
			}

			cursor := source.Paging.Next.After
			if cursor != "" {
				if _, repeated := seenCursors[sourceID][cursor]; repeated {
					return nil, fmt.Errorf("association batch read returned repeated cursor %q for source ID %q", cursor, sourceID)
				}
				seenCursors[sourceID][cursor] = struct{}{}
				next = append(next, associationsmodels.BatchGetAssociationsInput{Id: sourceID, After: cursor})
			}
		}
		for _, input := range pending {
			if _, found := returned[input.Id]; !found {
				return nil, fmt.Errorf("association batch read omitted source ID %q", input.Id)
			}
		}
		pending = next
	}

	return results, nil
}

func (c *AssociationService) BatchCreateAssociations(fromObject, toObject string, body associationsmodels.BatchCreateAssociationsBody) (crmmodels.BatchResponse, error) {
	var batchResp crmmodels.BatchResponse
	reqBody, err := json.Marshal(body)
	if err != nil {
		return batchResp, fmt.Errorf("error marshalling post body: %s", err)
	}
	resp, err := c.SendRequest(
		http.MethodPost,
		fmt.Sprintf("/crm/v4/associations/%s/%s/batch/create", fromObject, toObject),
		reqBody,
	)
	if err != nil {
		return batchResp, err
	}
	defer resp.Body.Close()
	associationRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return batchResp, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 201 {
		return batchResp, fmt.Errorf("error returned by endpoint: %s", associationRawBody)
	}
	err = json.Unmarshal(associationRawBody, &batchResp)
	if err != nil {
		return batchResp, fmt.Errorf("error parsing body: %s", err)
	}
	return batchResp, nil
}

func (c *AssociationService) CreateAssociation(fromObject, toObject, fromObjectType, toObjectType string, body []associationsmodels.AssociationType) (crmmodels.Result, error) {
	var associationResp crmmodels.Result
	reqBody, err := json.Marshal(body)
	if err != nil {
		return associationResp, fmt.Errorf("error marshalling post body: %s", err)
	}
	resp, err := c.SendRequest(
		http.MethodPut,
		fmt.Sprintf("/crm/v4/objects/%s/%s/associations/%s/%s", fromObjectType, fromObject, toObjectType, toObject),
		reqBody,
	)
	if err != nil {
		return associationResp, err
	}
	defer resp.Body.Close()
	associationRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return associationResp, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 201 {
		return associationResp, fmt.Errorf("error returned by endpoint: %s", associationRawBody)
	}
	err = json.Unmarshal(associationRawBody, &associationResp)
	if err != nil {
		return associationResp, fmt.Errorf("error parsing body: %s", err)
	}
	return associationResp, nil
}

func (c *AssociationService) BatchArchiveAssociationLabels(fromObject, toObject string, body associationsmodels.BatchCreateAssociationsBody) error {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("error marshalling post body: %s", err)
	}
	resp, err := c.SendRequest(
		http.MethodPost,
		fmt.Sprintf("/crm/v4/associations/%s/%s/batch/labels/archive", fromObject, toObject),
		reqBody,
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	associationRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("error returned by endpoint: %s", associationRawBody)
	}
	return nil
}
