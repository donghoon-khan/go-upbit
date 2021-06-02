package goupbit

import "testing"

func TestNewClient(t *testing.T) {
	client := newClient()
	if client == nil {
		t.Fatalf("newClientForQuotationApi(): Expected client not to be nil")
	}
}
