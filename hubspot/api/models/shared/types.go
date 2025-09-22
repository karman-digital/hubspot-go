package sharedmodels

import "time"

type ErrorResponseBody struct {
	Status        string `json:"status"`
	Message       string `json:"message"`
	Category      string `json:"category"`
	CorrelationId string `json:"correlationId"`
}

type NestedMessage struct {
	IsValid bool   `json:"isValid"`
	Message string `json:"message"`
	Error   string `json:"error"`
	Name    string `json:"name"`
}

type ParsedMessage struct {
	PropertyValues []NestedMessage `json:"Property values were not valid"`
}

func (e ErrorResponseBody) Error() string {
	return e.Message
}

func (p ParsedMessage) EmailInvalidError() bool {
	emailValidation := false
	for _, msg := range p.PropertyValues {
		if msg.Error == "INVALID_EMAIL" {
			emailValidation = true
		}
	}
	return emailValidation
}

type BatchResponseBase struct {
	CompletedAt time.Time         `json:"completedAt"`
	NumErrors   int               `json:"numErrors"`
	RequestedAt time.Time         `json:"requestedAt"`
	StartedAt   time.Time         `json:"startedAt"`
	Links       map[string]string `json:"links"`
	Errors      []ErrorDetail     `json:"errors"`
	Status      string            `json:"status"`
}

type Paging struct {
	Next NextPage `json:"next,omitempty"`
}

type NextPage struct {
	After string `json:"after,omitempty"`
	Link  string `json:"link,omitempty"`
}

type Navigation struct {
	Link   string `json:"link"`
	After  string `json:"after,omitempty"`
	Before string `json:"before,omitempty"`
}

type ErrorDetail struct {
	SubCategory any                 `json:"subCategory"`
	Context     map[string][]string `json:"context"`
	Links       map[string]string   `json:"links"`
	ID          string              `json:"id"`
	Category    string              `json:"category"`
	Message     string              `json:"message"`
	Errors      []NestedError       `json:"errors"`
	Status      string              `json:"status"`
}

type NestedError struct {
	SubCategory string              `json:"subCategory"`
	Code        string              `json:"code"`
	In          string              `json:"in"`
	Context     map[string][]string `json:"context"`
	Message     string              `json:"message"`
}

type GetOptions struct {
	Properties            []string `url:"properties,omitempty"`
	PropertiesWithHistory []string `url:"propertiesWithHistory,omitempty"`
	Associations          []string `url:"associations,omitempty"`
	Archived              bool     `url:"archived,omitempty"`
	After                 string   `url:"after,omitempty"`
	Limit                 int      `url:"limit,omitempty"`
	IdProperty            string   `url:"idProperty,omitempty"`
}
