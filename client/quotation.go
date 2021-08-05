package client

import (
	"net/http"

	"github.com/donghoon-khan/go-upbit/client/model/quotation"
	"github.com/google/go-querystring/query"
)

type QuotationClient struct {
	*http.Client
}

func (qc *QuotationClient) GetMarketCodes(isDetails bool) (*quotation.Markets, error) {

	markets, err := qc.callByGetMethod(
		"/market/all",
		struct {
			IsDetails bool `url:"isDetails"`
		}{false})
	if err != nil {
		return nil, err
	}

	return quotation.GetMarketsFromJson(markets)
}

func (qc *QuotationClient) callByGetMethod(url string, queryValue interface{}) ([]byte, error) {
	values, err := query.Values(queryValue)
	if err != nil {
		panic(err)
	}
	encodedQuery := values.Encode()

	req, err := http.NewRequest("GET", apiUrl+"/"+apiVersion+url+"?"+encodedQuery, nil)
	if err != nil {
		return nil, err
	}

	return getResponse(qc.Client, req)
}
