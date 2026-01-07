package emailmodels

import "time"

type MarketingEmail struct {
	Id                 string                 `json:"id"`
	Name               string                 `json:"name"`
	ActiveDomain       string                 `json:"activeDomain,omitempty"`
	Archived           bool                   `json:"archived,omitempty"`
	Campaign           string                 `json:"campaign,omitempty"`
	Content            map[string]interface{} `json:"content,omitempty"`
	CreatedAt          time.Time              `json:"createdAt,omitempty"`
	CreatedById        int                    `json:"createdById,omitempty"`
	From               *EmailFrom             `json:"from,omitempty"`
	IsPublished        bool                   `json:"isPublished,omitempty"`
	IsTransactional    bool                   `json:"isTransactional,omitempty"`
	Language           string                 `json:"language,omitempty"`
	PublishDate        *time.Time             `json:"publishDate,omitempty"`
	SendOnPublish      bool                   `json:"sendOnPublish,omitempty"`
	State              string                 `json:"state,omitempty"`
	Subcategory        string                 `json:"subcategory,omitempty"`
	Subject            string                 `json:"subject,omitempty"`
	SubscriptionDetails *SubscriptionDetails  `json:"subscriptionDetails,omitempty"`
	To                 *EmailTo               `json:"to,omitempty"`
	Type               string                 `json:"type,omitempty"`
	UpdatedAt          time.Time              `json:"updatedAt,omitempty"`
	UpdatedById        int                    `json:"updatedById,omitempty"`
	Webversion         *Webversion            `json:"webversion,omitempty"`
}

type EmailFrom struct {
	FromName string `json:"fromName,omitempty"`
	ReplyTo  string `json:"replyTo,omitempty"`
}

type SubscriptionDetails struct {
	OfficeLocationId string `json:"officeLocationId,omitempty"`
}

type EmailTo struct {
	ContactIds       map[string]interface{} `json:"contactIds,omitempty"`
	ContactIlsLists  *ContactIlsLists        `json:"contactIlsLists,omitempty"`
	SuppressGraymail bool                    `json:"suppressGraymail,omitempty"`
}

type ContactIlsLists struct {
	Exclude []int `json:"exclude,omitempty"`
	Include []int `json:"include,omitempty"`
}

type Webversion struct {
	ExpiresAt       *time.Time `json:"expiresAt,omitempty"`
	MetaDescription string     `json:"metaDescription,omitempty"`
	RedirectToPageId int       `json:"redirectToPageId,omitempty"`
	RedirectToUrl   string     `json:"redirectToUrl,omitempty"`
}

