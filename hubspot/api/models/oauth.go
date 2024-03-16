package hubspotmodels

type TokenBody struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	IDToken      string `json:"id_token"`
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
