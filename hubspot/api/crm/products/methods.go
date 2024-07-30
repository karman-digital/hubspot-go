package products

import (
	"fmt"

	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (p *ProductService) GetProductByUniqueId(uniqueId string, opts ...hubspotmodels.GetOptions) (hubspotmodels.Result, error) {
	if opts[0].IdProperty == "" {
		return hubspotmodels.Result{}, fmt.Errorf("idProperty must be set for unique property search")
	}
	resp, err := p.SendRequest("GET", fmt.Sprintf("/crm/v3/objects/products/%s", uniqueId), nil, opts...)
	if err != nil {
		return hubspotmodels.Result{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleResponse(resp)
}
