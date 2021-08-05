package client

import (
	"net/http"
	"testing"

	"github.com/donghoon-khan/go-upbit/client/model/quotation"
	"github.com/stretchr/testify/assert"
)

var qc = &QuotationClient{http.DefaultClient}

func TestQuotationClient_GetMarketCodeList(t *testing.T) {
	cases := []struct {
		isDetails bool
		expected  *quotation.Market
	}{
		{
			true,
			&quotation.Market{
				Market:      "KRW-BTC",
				KoreanName:  "비트코인",
				EnglishName: "Bitcoin",
				Warning:     "",
			},
		},
		{
			false,
			&quotation.Market{
				Market:      "KRW-BTC",
				KoreanName:  "비트코인",
				EnglishName: "Bitcoin",
				Warning:     "NONE",
			},
		},
	}
	for _, c := range cases {
		_, e := qc.GetMarketCodes(c.isDetails)
		assert.NoError(t, e)
	}
}
