package quotation

import "encoding/json"

type MarketCode struct {
	Market      string `json:"market"`
	KoreanName  string `json:"korean_name"`
	EnglishName string `json:"english_name"`
	Warning     string `json:"market_warning"`
}

type MarketCodes struct {
	items []MarketCode
}

func MarketCodesFromJson(b []byte) (*MarketCodes, error) {
	var marketCodes []MarketCode

	e := json.Unmarshal(b, &marketCodes)
	return &MarketCodes{
		items: marketCodes,
	}, e
}
