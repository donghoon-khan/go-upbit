package goupbit

import (
	"net/http"
	"time"
)

const (
	UPBIT_BASE_URL = "https://api.upbit.com/v1"
)

type client struct {
	baseURL    string
	accessKey  string
	secretKey  string
	httpClient *http.Client
}

func newClientForQuotationApi() *client {
	return &client{
		baseURL: UPBIT_BASE_URL,
		httpClient: &http.Client{
			Timeout: time.Second,
		},
	}
}
