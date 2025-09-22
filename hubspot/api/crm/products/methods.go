package products

import (
	"fmt"

	crmmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm"
	sharedmodels "github.com/karman-digital/hubspot/hubspot/api/models/shared"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (p *ProductService) GetProductByUniqueId(uniqueId string, opts ...sharedmodels.GetOptions) (crmmodels.Result, error) {
	if opts[0].IdProperty == "" {
		return crmmodels.Result{}, fmt.Errorf("idProperty must be set for unique property search")
	}
	resp, err := p.SendRequest("GET", fmt.Sprintf("/crm/v3/objects/products/%s", uniqueId), nil, opts...)
	if err != nil {
		return crmmodels.Result{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleResponse(resp)
}
