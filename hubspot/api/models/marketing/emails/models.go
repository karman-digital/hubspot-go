package emailmodels

import "time"

type MarketingEmail struct {
	Id                   string                 `json:"id"`
	Name                 string                 `json:"name"`
	ActiveDomain         string                 `json:"activeDomain,omitempty"`
	AllEmailCampaignIds  []string               `json:"allEmailCampaignIds,omitempty"`
	Archived             bool                   `json:"archived,omitempty"`
	BusinessUnitId       string                 `json:"businessUnitId,omitempty"`
	Campaign             string                 `json:"campaign,omitempty"`
	CampaignName         string                 `json:"campaignName,omitempty"`
	CampaignUtm          string                 `json:"campaignUtm,omitempty"`
	Content              map[string]interface{} `json:"content,omitempty"`
	CreatedAt            time.Time              `json:"createdAt,omitempty"`
	CreatedById          string                 `json:"createdById,omitempty"`
	EmailTemplateMode    string                 `json:"emailTemplateMode,omitempty"`
	From                 *EmailFrom             `json:"from,omitempty"`
	IsAb                 bool                   `json:"isAb,omitempty"`
	IsPublished          bool                   `json:"isPublished,omitempty"`
	IsTransactional      bool                   `json:"isTransactional,omitempty"`
	JitterSendTime       bool                   `json:"jitterSendTime,omitempty"`
	Language             string                 `json:"language,omitempty"`
	PreviewKey           string                 `json:"previewKey,omitempty"`
	PrimaryEmailCampaignId string               `json:"primaryEmailCampaignId,omitempty"`
	PublishDate          *time.Time             `json:"publishDate,omitempty"`
	PublishedAt          *time.Time             `json:"publishedAt,omitempty"`
	PublishedByEmail     string                 `json:"publishedByEmail,omitempty"`
	PublishedById        string                 `json:"publishedById,omitempty"`
	PublishedByName      string                 `json:"publishedByName,omitempty"`
	SendOnPublish        bool                   `json:"sendOnPublish,omitempty"`
	State                string                 `json:"state,omitempty"`
	Subcategory          string                 `json:"subcategory,omitempty"`
	Subject              string                 `json:"subject,omitempty"`
	SubscriptionDetails  *SubscriptionDetails   `json:"subscriptionDetails,omitempty"`
	To                   *EmailTo               `json:"to,omitempty"`
	Type                 string                 `json:"type,omitempty"`
	UpdatedAt            time.Time              `json:"updatedAt,omitempty"`
	UpdatedById          string                 `json:"updatedById,omitempty"`
	Webversion           *Webversion            `json:"webversion,omitempty"`
}

type EmailFrom struct {
	FromName string `json:"fromName,omitempty"`
	ReplyTo  string `json:"replyTo,omitempty"`
}

type SubscriptionDetails struct {
	OfficeLocationId string `json:"officeLocationId,omitempty"`
	SubscriptionId   string `json:"subscriptionId,omitempty"`
	SubscriptionName string `json:"subscriptionName,omitempty"`
}

type EmailTo struct {
	ContactIds       *ContactIdsList `json:"contactIds,omitempty"`
	ContactIlsLists  *ContactIlsLists `json:"contactIlsLists,omitempty"`
	ContactLists     *ContactLists    `json:"contactLists,omitempty"`
	SuppressGraymail bool             `json:"suppressGraymail,omitempty"`
}

type ContactIdsList struct {
	Exclude []string `json:"exclude,omitempty"`
	Include []string `json:"include,omitempty"`
}

type ContactIlsLists struct {
	Exclude []string `json:"exclude,omitempty"`
	Include []string `json:"include,omitempty"`
}

type ContactLists struct {
	Exclude []string `json:"exclude,omitempty"`
	Include []string `json:"include,omitempty"`
}

type Webversion struct {
	Domain          string `json:"domain,omitempty"`
	Enabled         bool   `json:"enabled,omitempty"`
	IsPageRedirected bool  `json:"isPageRedirected,omitempty"`
	Slug            string `json:"slug,omitempty"`
	Url             string `json:"url,omitempty"`
}

