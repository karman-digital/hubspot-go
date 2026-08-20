package invoices

import (
	"encoding/json"
	"fmt"
	"net/http"

	crmmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm"
	associationsmodels "github.com/karman-digital/hubspot/hubspot/api/models/crm/associations"
	sharedmodels "github.com/karman-digital/hubspot/hubspot/api/models/shared"
	"github.com/karman-digital/hubspot/hubspot/api/shared"
)

func (s *Service) CreateInvoice(body crmmodels.PostBody) (crmmodels.Result, error) {
	requestBody, err := json.Marshal(body)
	if err != nil {
		return crmmodels.Result{}, fmt.Errorf("marshal invoice: %w", err)
	}
	response, err := s.SendRequest(http.MethodPost, invoicePath(""), requestBody)
	if err != nil {
		return crmmodels.Result{}, fmt.Errorf("create invoice: %w", err)
	}
	return shared.HandleCreateResponse(response)
}

func (s *Service) UpdateInvoice(id string, body crmmodels.PatchBody) (crmmodels.Result, error) {
	requestBody, err := json.Marshal(body)
	if err != nil {
		return crmmodels.Result{}, fmt.Errorf("marshal invoice: %w", err)
	}
	response, err := s.SendRequest(http.MethodPatch, invoicePath(id), requestBody)
	if err != nil {
		return crmmodels.Result{}, fmt.Errorf("update invoice: %w", err)
	}
	return shared.HandleResponse(response)
}

func (s *Service) GetInvoice(id string, options ...sharedmodels.GetOptions) (crmmodels.Result, error) {
	response, err := s.SendRequest(http.MethodGet, invoicePath(id), nil, options...)
	if err != nil {
		return crmmodels.Result{}, fmt.Errorf("get invoice: %w", err)
	}
	return shared.HandleResponse(response)
}

func (s *Service) SearchInvoices(body crmmodels.SearchBody) (crmmodels.SearchResponse, error) {
	requestBody, err := json.Marshal(body)
	if err != nil {
		return crmmodels.SearchResponse{}, fmt.Errorf("marshal invoice search: %w", err)
	}
	response, err := s.SendRequest(http.MethodPost, invoiceSearchPath(), requestBody)
	if err != nil {
		return crmmodels.SearchResponse{}, fmt.Errorf("search invoices: %w", err)
	}
	return shared.HandleSearchResponse(response)
}

func (s *Service) GetCompanyAssociations(id string) (associationsmodels.AssociationGetResponse, error) {
	response, err := s.SendRequest(http.MethodGet, invoiceCompanyAssociationsPath(id), nil)
	if err != nil {
		return associationsmodels.AssociationGetResponse{}, fmt.Errorf("get invoice company associations: %w", err)
	}
	defer response.Body.Close()
	body, err := shared.HandleBasicResponseCode(response)
	if err != nil {
		return associationsmodels.AssociationGetResponse{}, err
	}
	var result associationsmodels.AssociationGetResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return result, fmt.Errorf("decode invoice company associations: %w", err)
	}
	return result, nil
}

func (s *Service) AssociateCompany(invoiceID, companyID string) error {
	response, err := s.SendRequest(http.MethodPut, invoiceCompanyAssociationPath(invoiceID, companyID), nil)
	if err != nil {
		return fmt.Errorf("associate invoice company: %w", err)
	}
	defer response.Body.Close()
	_, err = shared.HandleBasicResponseCode(response)
	return err
}

func (s *Service) RemoveCompanyAssociation(invoiceID, companyID string) error {
	response, err := s.SendRequest(http.MethodDelete, invoiceCompanyAssociationDeletePath(invoiceID, companyID), nil)
	if err != nil {
		return fmt.Errorf("remove invoice company association: %w", err)
	}
	defer response.Body.Close()
	return shared.HandleDeleteResponse(response)
}

func (s *Service) GetInvoiceAssociations(invoiceID, objectType string) (associationsmodels.AssociationGetResponse, error) {
	response, err := s.SendRequest(http.MethodGet, invoiceAssociationsPath(invoiceID, objectType), nil)
	if err != nil {
		return associationsmodels.AssociationGetResponse{}, fmt.Errorf("get invoice associations: %w", err)
	}
	defer response.Body.Close()
	body, err := shared.HandleBasicResponseCode(response)
	if err != nil {
		return associationsmodels.AssociationGetResponse{}, err
	}
	var result associationsmodels.AssociationGetResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return result, fmt.Errorf("decode invoice associations: %w", err)
	}
	return result, nil
}

func (s *Service) AssociateInvoice(invoiceID, objectType, objectID string, associationTypeID int) error {
	body, err := json.Marshal([]associationsmodels.AssociationType{{
		AssociationCategory: "USER_DEFINED", AssociationTypeId: associationTypeID,
	}})
	if err != nil {
		return fmt.Errorf("marshal invoice association: %w", err)
	}
	response, err := s.SendRequest(http.MethodPut, invoiceAssociationPath(invoiceID, objectType, objectID), body)
	if err != nil {
		return fmt.Errorf("associate invoice: %w", err)
	}
	defer response.Body.Close()
	_, err = shared.HandleBasicResponseCode(response)
	return err
}

func (s *Service) RemoveInvoiceAssociation(invoiceID, objectType, objectID string) error {
	response, err := s.SendRequest(http.MethodDelete, invoiceAssociationPath(invoiceID, objectType, objectID), nil)
	if err != nil {
		return fmt.Errorf("remove invoice association: %w", err)
	}
	defer response.Body.Close()
	return shared.HandleDeleteResponse(response)
}

func invoicePath(id string) string {
	if id == "" {
		return "/crm/v3/objects/invoices"
	}
	return "/crm/v3/objects/invoices/" + id
}

func invoiceSearchPath() string {
	return invoicePath("") + "/search"
}

func invoiceCompanyAssociationsPath(invoiceID string) string {
	return "/crm/v4/objects/invoices/" + invoiceID + "/associations/companies"
}

func invoiceCompanyAssociationPath(invoiceID, companyID string) string {
	return "/crm/v4/objects/invoices/" + invoiceID + "/associations/default/companies/" + companyID
}

func invoiceCompanyAssociationDeletePath(invoiceID, companyID string) string {
	return invoiceCompanyAssociationsPath(invoiceID) + "/" + companyID
}

func invoiceAssociationsPath(invoiceID, objectType string) string {
	return "/crm/v4/objects/invoices/" + invoiceID + "/associations/" + objectType
}

func invoiceAssociationPath(invoiceID, objectType, objectID string) string {
	return invoiceAssociationsPath(invoiceID, objectType) + "/" + objectID
}
