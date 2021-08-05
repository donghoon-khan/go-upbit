package client

import (
	"net/http"
	"reflect"
	"testing"
)

var qc = &QuotationClient{http.DefaultClient}

func TestQuotationClient_GetMarketCodeList(t *testing.T) {
	cases := []struct {
		isDetails bool
		expected  *MarketCode
	}{
		{
			true, &MarketCode{
				Market:      "",
				KoreanName:  "",
				EnglishName: "",
				Warning:     "",
			}},
	}
	for _, c := range cases {
		actual := qc.GetMarketCodes(false)
		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("GetMarketCodeList(%#v) == \n%#v\nexpected \n%#v\n", c.isDetails, actual, c.expected)
		}
	}
}
