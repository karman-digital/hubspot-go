package hubspotmodels

type AccountDataResponse struct {
	Response AccountsResponse `json:"response"`
}

type AccountsResponse struct {
	Accounts []Account `json:"accounts"`
}

type Account struct {
	AccountId   string `json:"accountId"`
	AccountName string `json:"accountName"`
}

type SettingsDropdownResponse struct {
	Response SettingsDropdownOptions `json:"response"`
}

type SettingsDropdownOptions struct {
	Options        []SettingsDropdown `json:"options"`
	SelectedOption string             `json:"selectedOption"`
}

type SettingsDropdown struct {
	Text  string `json:"text"`
	Value string `json:"value"`
}

type SettingsActionPayload struct {
	ActionType     string `json:"actionType"`
	PortalId       string `json:"portalId"`
	UserId         string `json:"userId"`
	UserEmail      string `json:"userEmail"`
	AppId          string `json:"appId"`
	AccountId      string `json:"accountId"`
	SelectedOption string `json:"selectedOption"`
}
