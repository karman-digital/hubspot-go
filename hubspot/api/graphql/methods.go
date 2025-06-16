package graphql

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
)

func (g *GraphQLService) MakeRequest(query string, variables map[string]interface{}) (map[string]interface{}, error) {
	response, err := g.MakeRequestWithFullResponse(query, variables)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (g *GraphQLService) MakeRequestWithFullResponse(query string, variables map[string]interface{}) (hubspotmodels.GraphQLResponse, error) {
	reqBody := map[string]interface{}{
		"query": query,
	}
	if variables != nil {
		reqBody["variables"] = variables
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return hubspotmodels.GraphQLResponse{}, fmt.Errorf("error marshalling request body: %s", err)
	}

	resp, err := g.SendRequest(http.MethodPost, "/collector/graphql", body)
	if err != nil {
		return hubspotmodels.GraphQLResponse{}, fmt.Errorf("error making GraphQL request: %s", err)
	}

	return handleGraphQLResponse(resp)
}

func handleGraphQLResponse(resp *http.Response) (hubspotmodels.GraphQLResponse, error) {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return hubspotmodels.GraphQLResponse{}, fmt.Errorf("error reading response body: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		return hubspotmodels.GraphQLResponse{}, fmt.Errorf("GraphQL request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var result hubspotmodels.GraphQLResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return hubspotmodels.GraphQLResponse{}, fmt.Errorf("error unmarshalling response: %s", err)
	}

	if len(result.Errors) > 0 && result.Data == nil {
		return result, fmt.Errorf("GraphQL errors: %v", result.Errors)
	}

	return result, nil
}
