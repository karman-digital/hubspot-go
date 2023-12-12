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
