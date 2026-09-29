package authmodels

import "errors"

type AccessToken string
type RefreshToken string

type ClientSecret string

type ClientId string

type RedirectUri string

func (r *RefreshToken) Set(s string) error {
	if s == "" {
		return errors.New("refresh token cannot be empty")
	}
	*r = RefreshToken(s)
	return nil
}

func (r RefreshToken) String() string {
	return string(r)
}

func (a *AccessToken) Set(s string) error {
	if s == "" {
		return errors.New("access token cannot be empty")
	}
	*a = AccessToken(s)
	return nil
}

func (a AccessToken) String() string {
	return string(a)
}

func (c ClientSecret) String() string {
	return string(c)
}

func (c *ClientSecret) Set(s string) {
	*c = ClientSecret(s)
}

func (c ClientId) String() string {
	return string(c)
}

func (c *ClientId) Set(s string) {
	*c = ClientId(s)
}

func (r RedirectUri) String() string {
	return string(r)
}

func (r *RedirectUri) Set(s string) {
	*r = RedirectUri(s)
}

type TokenBody struct {
	AccessToken  AccessToken  `json:"access_token"`
	ExpiresIn    int64        `json:"expires_in"`
	RefreshToken RefreshToken `json:"refresh_token"`
	TokenType    string       `json:"token_type"`
	TokenUse     string       `json:"token_use"`
	IDToken      string       `json:"id_token"`
	HubID        int          `json:"hub_id"`
	UserID       int          `json:"user_id"`
	Scopes       []string     `json:"scopes"`
}

type TokenInfoResponse struct {
	Active                bool               `json:"active"`
	AppID                 int                `json:"app_id"`
	ClientID              string             `json:"client_id"`
	ExpiresIn             int64              `json:"expires_in"`
	HubDomain             string             `json:"hub_domain"`
	HubID                 int                `json:"hub_id"`
	IsPrivateDistribution bool               `json:"is_private_distribution"`
	Scopes                []string           `json:"scopes"`
	SignedAccessToken     *SignedAccessToken `json:"signed_access_token,omitempty"`
	Token                 string             `json:"token"`
	TokenType             string             `json:"token_type"`
	TokenUse              string             `json:"token_use"`
	User                  string             `json:"user"`
	UserID                int                `json:"user_id"`
}

type SignedAccessToken struct {
	AppID                     int    `json:"appId"`
	AppInstallID              string `json:"appInstallId"`
	Audience                  string `json:"audience"`
	ExpiresAt                 int64  `json:"expiresAt"`
	HubID                     int    `json:"hubId"`
	Hublet                    string `json:"hublet"`
	InstallingUserID          int    `json:"installingUserId"`
	IsPrivateDistribution     bool   `json:"isPrivateDistribution"`
	IsServiceAccount          bool   `json:"isServiceAccount"`
	IsUserLevel               bool   `json:"isUserLevel"`
	NewSignature              string `json:"newSignature"`
	ScopeToScopeGroupPKs      string `json:"scopeToScopeGroupPks"`
	Scopes                    string `json:"scopes"`
	Signature                 string `json:"signature"`
	TrialScopeToScopeGroupPKs string `json:"trialScopeToScopeGroupPks"`
	TrialScopes               string `json:"trialScopes"`
	UserID                    int    `json:"userId"`
}

type OAuthErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	Status           string `json:"status"`
	Message          string `json:"message"`
}

// BearerTokenBody is retained as an alias while callers migrate to the dated
// token introspection model.
type BearerTokenBody = TokenInfoResponse
