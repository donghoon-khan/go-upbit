package goupbit

import "testing"

func TestNewClientForQuotationApi(t *testing.T) {
	quotationApi := newClientForQuotationApi()
	if quotationApi == nil {
		t.Fatalf("NewClientForQuotationApi(): Expected client not to be nil")
	}
}
