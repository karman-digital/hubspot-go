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
	ExpiresIn    int          `json:"expires_in"`
	RefreshToken RefreshToken `json:"refresh_token"`
	TokenType    string       `json:"token_type"`
	IDToken      string       `json:"id_token"`
}

type BearerTokenBody struct {
	Token                     string   `json:"token"`
	User                      string   `json:"user"`
	HubDomain                 string   `json:"hub_domain"`
	Scopes                    []string `json:"scopes"`
	ScopeToScopeGroupPks      []string `json:"scope_to_scope_group_pks"`
	TrialScopes               []string `json:"trial_scopes"`
	TrialScopeToScopeGroupPks []string `json:"trial_scope_to_scope_group_pks"`
	HubID                     int      `json:"hub_id"`
	AppID                     int      `json:"app_id"`
	ExpiresIn                 int      `json:"expires_in"`
	UserID                    int      `json:"user_id"`
	TokenType                 string   `json:"token_type"`
}
