package invoices

import "testing"

func TestInvoicePaths(t *testing.T) {
	tests := map[string]string{
		invoicePath(""):                                    "/crm/v3/objects/invoices",
		invoicePath("123"):                                 "/crm/v3/objects/invoices/123",
		invoiceSearchPath():                                "/crm/v3/objects/invoices/search",
		invoiceCompanyAssociationsPath("123"):              "/crm/v4/objects/invoices/123/associations/companies",
		invoiceCompanyAssociationPath("123", "456"):        "/crm/v4/objects/invoices/123/associations/default/companies/456",
		invoiceCompanyAssociationDeletePath("123", "456"):  "/crm/v4/objects/invoices/123/associations/companies/456",
		invoiceAssociationsPath("123", "2-32302481"):       "/crm/v4/objects/invoices/123/associations/2-32302481",
		invoiceAssociationPath("123", "2-32302481", "789"): "/crm/v4/objects/invoices/123/associations/2-32302481/789",
	}
	for got, want := range tests {
		if got != want {
			t.Fatalf("path = %q, want %q", got, want)
		}
	}
}
