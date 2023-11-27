package hubspot

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/karman-digital/hubspot/hubspot/models"
)

func (c *credentials) UpdateContact(id int, patchBody models.PatchBody) (models.ContactResponse, error) {
	var contactResp models.ContactResponse
	reqUrl := fmt.Sprintf("https://api.hubapi.com/crm/v3/objects/contacts/%d", id)
	reqBody, err := json.Marshal(patchBody)
	if err != nil {
		return contactResp, fmt.Errorf("error marshalling patch body: %s", err)
	}
	req, err := retryablehttp.NewRequest("PATCH", reqUrl, reqBody)
	if err != nil {
		return contactResp, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
	resp, err := c.Client.Do(req)
	if err != nil {
		return contactResp, fmt.Errorf("error making request: %s", err)
	}
	defer resp.Body.Close()
	contactRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return contactResp, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != 200 {
		return contactResp, fmt.Errorf("error returned by endpoint: %s", contactRawBody)
	}
	err = json.Unmarshal(contactRawBody, &contactResp)
	if err != nil {
		return contactResp, fmt.Errorf("error parsing body: %s", err)
	}
	return contactResp, nil
}
