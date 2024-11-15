package credentials

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
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

func (c *Credentials) RefreshTokenPair() error {
	tokenBody := hubspotmodels.TokenBody{}
	data := url.Values{
		"grant_type":    []string{"refresh_token"},
		"redirect_uri":  []string{c.RedirectUri().String()},
		"client_id":     []string{c.ClientId().String()},
		"client_secret": []string{c.ClientSecret().String()},
		"refresh_token": []string{c.RefreshToken().String()},
	}
	req, err := retryablehttp.NewRequest("POST", "https://api.hubapi.com/oauth/v1/token", strings.NewReader(data.Encode()))
	if err != nil {
		return fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := c.Client().Do(req)
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
	return c.SetTokens(tokenBody.AccessToken, tokenBody.RefreshToken)
}

func (c *Credentials) SetTokens(accessToken hubspotmodels.AccessToken, refreshToken hubspotmodels.RefreshToken) error {
	if err := c.SetAccessToken(accessToken.String()); err != nil {
		return fmt.Errorf("error setting access token: %w", err)
	}
	if err := c.SetRefreshToken(refreshToken.String()); err != nil {
		return fmt.Errorf("error setting refresh token: %w", err)
	}
	return nil
}

func (c *Credentials) ValidateBearerToken() (bool, error) {
	resBodyStruct := hubspotmodels.BearerTokenBody{}
	res, err := http.Get(fmt.Sprintf("https://api.hubapi.com/oauth/v1/access-tokens/%s", c.AccessToken().String()))
	if err != nil {
		return false, err
	}
	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return false, err
	}
	err = json.Unmarshal(resBody, &resBodyStruct)
	if err != nil {
		return false, err
	}
	if res.StatusCode != 200 || resBodyStruct.ExpiresIn < 300 {
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
	err = json.Unmarshal(resBody, &resBodyStruct)
	if err != nil {
		return resBodyStruct, err
	}
	return resBodyStruct, nil
}
