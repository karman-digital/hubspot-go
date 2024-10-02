package hubspotmodels

type BatchDeleteBody struct {
	Inputs []BatchInput `json:"inputs"`
}

type BatchInput struct {
	Id int `json:"id"`
}
