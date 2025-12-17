package emailanalyticscampaignmodels

type EmailCampaign struct {
	Id             int64  `json:"id"`
	AppId          int64  `json:"appId"`
	AppName        string `json:"appName"`
	LastUpdatedTime int64 `json:"lastUpdatedTime"`
}

type EmailCampaignsResponse struct {
	HasMore   bool           `json:"hasMore"`
	Offset    string         `json:"offset"`
	Campaigns []EmailCampaign `json:"campaigns"`
}

type EmailCampaignDetail struct {
	Id          int64                `json:"id"`
	AppId       int64                `json:"appId"`
	AppName     string               `json:"appName"`
	ContentId   int64                `json:"contentId"`
	Counters    EmailCampaignCounters `json:"counters"`
	Name        string               `json:"name"`
	NumIncluded int                  `json:"numIncluded"`
	NumQueued   int                  `json:"numQueued"`
	SubType     string               `json:"subType"`
	Subject     string               `json:"subject"`
	Type        string               `json:"type"`
}

type EmailCampaignCounters struct {
	Delivered int `json:"delivered"`
	Open      int `json:"open"`
	Processed int `json:"processed"`
	Sent      int `json:"sent"`
}

