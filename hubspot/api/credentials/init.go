package credentials

import "github.com/hashicorp/go-retryablehttp"

func NewHubspotOauthCredentials(clientId string, clientSecret string, redirect_uri string, accessToken string, refreshToken string) *Credentials {
	var creds Credentials
	client := retryablehttp.NewClient()
	client.Logger = nil
	creds.SetClient(client)
	creds.SetAccessToken(accessToken)
	creds.SetRefreshToken(refreshToken)
	creds.SetClientId(clientId)
	creds.SetClientSecret(clientSecret)
	creds.SetRedirectUri(redirect_uri)
	return &creds
}
