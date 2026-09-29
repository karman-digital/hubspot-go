package products

import (
	"fmt"
	"net/url"

	crmmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm"
	sharedmodels "github.com/karman-digital/hubspot/hubspot/api/models/shared"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (p *ProductService) GetProductByUniqueId(uniqueId string, opts ...sharedmodels.GetOptions) (crmmodels.Result, error) {
	if len(opts) == 0 || opts[0].IdProperty == "" {
		return crmmodels.Result{}, fmt.Errorf("idProperty must be set for unique property search")
	}
	resp, err := p.SendRequest("GET", fmt.Sprintf("/crm/objects/2026-09/products/%s", url.PathEscape(uniqueId)), nil, opts...)
	if err != nil {
		return crmmodels.Result{}, fmt.Errorf("error making request: %s", err)
	}
	return shared.HandleResponse(resp)
}
