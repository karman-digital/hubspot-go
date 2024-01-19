package hubspot

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
	apptypes "github.com/karman-digital/integrations/types"
)

func GenerateTokenPair(code string, clientId string, clientSecret string, redirectURI string) (hubspotmodels.TokenBody, error) {
	resBodyStruct := hubspotmodels.TokenBody{}
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("client_id", clientId)
	data.Set("client_secret", clientSecret)
	data.Set("redirect_uri", redirectURI)
	req, err := http.NewRequest(http.MethodPost, "https://api.hubapi.com/oauth/v1/token", strings.NewReader(data.Encode()))
	if err != nil {
		return resBodyStruct, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		return resBodyStruct, err
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return resBodyStruct, err
	}
	if res.StatusCode != 200 {
		return resBodyStruct, errors.New(string(resBody))
	}
	err = json.Unmarshal(resBody, &resBodyStruct)
	if err != nil {
		return resBodyStruct, err
	}
	return resBodyStruct, nil
}

func (c *credentials) RefreshTokenPair(clientSecret string, clientId string, redirectUri string) error {
	var accessToken apptypes.AccessToken
	var refreshToken apptypes.RefreshToken
	tokenBody := hubspotmodels.TokenBody{}
	client := &http.Client{}

	data := url.Values{
		"grant_type":    []string{"refresh_token"},
		"redirect_uri":  []string{redirectUri},
		"client_id":     []string{clientId},
		"client_secret": []string{clientSecret},
		"refresh_token": []string{c.RefreshToken.String()},
	}

	req, err := http.NewRequest("POST", "https://api.hubapi.com/oauth/v1/token", strings.NewReader(data.Encode()))
	if err != nil {
		return fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making post request: %s", err)
	}
	defer resp.Body.Close()
	tokenRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error returned by endpoint: %s", tokenRawBody)
	}
	err = json.Unmarshal(tokenRawBody, &tokenBody)
	if err != nil {
		return fmt.Errorf("error parsing body: %s", err)
	}
	c.SetAccessToken(accessToken)
	c.SetRefreshToken(refreshToken)
	return nil
}

func (c *credentials) ValidateBearerToken() (bool, error) {
	res, err := http.Get(fmt.Sprintf("https://api.hubapi.com/oauth/v1/access-tokens/%s", c.AccessToken))
	if err != nil {
		return false, err
	}
	if res.StatusCode != 200 {
		return false, nil
	}
	return true, nil
}

func GetBearerTokenData(bearerToken string) (hubspotmodels.BearerTokenBody, error) {
	resBodyStruct := hubspotmodels.BearerTokenBody{}

	res, err := http.Get(fmt.Sprintf("https://api.hubapi.com/oauth/v1/access-tokens/%s", bearerToken))
	if err != nil {
		return resBodyStruct, err
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return resBodyStruct, err
	}
	if res.StatusCode != 200 {
		return resBodyStruct, errors.New(string(resBody))
	}
	json.Unmarshal(resBody, &resBodyStruct)

	return resBodyStruct, nil
}
