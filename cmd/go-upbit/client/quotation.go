package client

import (
	"encoding/json"
	"net/http"

	"github.com/google/go-querystring/query"
)

type QuotationClient struct {
	*http.Client
}

type MarketCode struct {
	Market      string `json:"market"`
	KoreanName  string `json:"korean_name"`
	EnglishName string `json:"english_name"`
	Warning     string `json:"market_warning"`
}

func (qc *QuotationClient) GetMarketCodes(isDetails bool) []MarketCode {
	var result []MarketCode
	_ = qc.callByGetMethod(
		"/market/all",
		struct {
			IsDetails bool `url:"isDetails"`
		}{isDetails},
		&result,
	)
	return result
}

func (qc *QuotationClient) callByGetMethod(url string, queryValue interface{}, response interface{}) error {
	values, err := query.Values(queryValue)
	if err != nil {
		panic(err)
	}
	encodedQuery := values.Encode()

	req, err := http.NewRequest("GET", apiUrl+"/"+apiVersion+url+"?"+encodedQuery, nil)
	if err != nil {
		return err
	}

	result, _ := getResponse(qc.Client, req)
	return json.Unmarshal(result, &response)
}
