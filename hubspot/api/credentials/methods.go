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
	authmodels "github.com/karman-digital/hubspot/hubspot/api/models/auth"
)

const (
	oauthTokenURL      = "https://api.hubapi.com/oauth/2026-09/token"
	oauthIntrospectURL = "https://api.hubapi.com/oauth/2026-09/token/introspect"
	accessTokenHint    = "access_token"
)

func GenerateTokenPair(code string, clientID string, clientSecret string, redirectURI string) (authmodels.TokenBody, error) {
	result := authmodels.TokenBody{}
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("redirect_uri", redirectURI)

	req, err := http.NewRequest(http.MethodPost, oauthTokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return result, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := (&http.Client{Timeout: 30 * time.Second}).Do(req)
	if err != nil {
		return result, err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return result, err
	}
	if res.StatusCode != http.StatusOK {
		return result, oauthResponseError(res.StatusCode, resBody)
	}
	if err := json.Unmarshal(resBody, &result); err != nil {
		return result, err
	}
	return result, nil
}

func (c *Credentials) RefreshTokenPair() error {
	data := url.Values{
		"grant_type":    []string{"refresh_token"},
		"redirect_uri":  []string{c.RedirectUri().String()},
		"client_id":     []string{c.ClientId().String()},
		"client_secret": []string{c.ClientSecret().String()},
		"refresh_token": []string{c.RefreshToken().String()},
	}
	req, err := retryablehttp.NewRequest(http.MethodPost, oauthTokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return fmt.Errorf("create refresh request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := c.Client().Do(req)
	if err != nil {
		return fmt.Errorf("refresh token: %w", err)
	}
	defer resp.Body.Close()

	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read refresh response: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return oauthResponseError(resp.StatusCode, rawBody)
	}
	var tokenBody authmodels.TokenBody
	if err := json.Unmarshal(rawBody, &tokenBody); err != nil {
		return fmt.Errorf("decode refresh response: %w", err)
	}
	return c.SetTokens(tokenBody.AccessToken, tokenBody.RefreshToken)
}

func (c *Credentials) SetTokens(accessToken authmodels.AccessToken, refreshToken authmodels.RefreshToken) error {
	if err := c.SetAccessToken(accessToken.String()); err != nil {
		return fmt.Errorf("error setting access token: %w", err)
	}
	if err := c.SetRefreshToken(refreshToken.String()); err != nil {
		return fmt.Errorf("error setting refresh token: %w", err)
	}
	return nil
}

func (c *Credentials) ValidateBearerToken() (bool, error) {
	result, err := c.GetBearerTokenData(c.AccessToken().String(), accessTokenHint)
	if err != nil {
		return false, err
	}
	return result.Active && result.ExpiresIn >= 150, nil
}

// GetBearerTokenData introspects a token using the client credentials held by
// this Credentials instance. tokenTypeHint must be access_token or refresh_token.
func (c *Credentials) GetBearerTokenData(token string, tokenTypeHint string) (authmodels.TokenInfoResponse, error) {
	return introspectToken(c.Client(), token, c.ClientId().String(), c.ClientSecret().String(), tokenTypeHint)
}

// GetBearerTokenData introspects a token when no Credentials instance is
// available. Client credentials are required by the 2026-09 endpoint.
func GetBearerTokenData(token string, clientID string, clientSecret string, tokenTypeHint string) (authmodels.TokenInfoResponse, error) {
	client := retryablehttp.NewClient()
	client.Logger = nil
	return introspectToken(client, token, clientID, clientSecret, tokenTypeHint)
}

func introspectToken(client *retryablehttp.Client, token string, clientID string, clientSecret string, tokenTypeHint string) (authmodels.TokenInfoResponse, error) {
	var result authmodels.TokenInfoResponse
	if client == nil {
		return result, errors.New("oauth client is required")
	}
	data := url.Values{}
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("token", token)
	data.Set("token_type_hint", tokenTypeHint)
	req, err := retryablehttp.NewRequest(http.MethodPost, oauthIntrospectURL, strings.NewReader(data.Encode()))
	if err != nil {
		return result, fmt.Errorf("create token introspection request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return result, fmt.Errorf("introspect token: %w", err)
	}
	defer resp.Body.Close()

	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, fmt.Errorf("read token introspection response: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return result, oauthResponseError(resp.StatusCode, rawBody)
	}
	if err := json.Unmarshal(rawBody, &result); err != nil {
		return result, fmt.Errorf("decode token introspection response: %w", err)
	}
	return result, nil
}

func oauthResponseError(statusCode int, rawBody []byte) error {
	var response authmodels.OAuthErrorResponse
	if err := json.Unmarshal(rawBody, &response); err == nil {
		if response.Error != "" || response.ErrorDescription != "" {
			return fmt.Errorf("oauth request returned %d: %s: %s", statusCode, response.Error, response.ErrorDescription)
		}
		if response.Message != "" {
			return fmt.Errorf("oauth request returned %d: %s", statusCode, response.Message)
		}
	}
	return fmt.Errorf("oauth request returned %d", statusCode)
}
