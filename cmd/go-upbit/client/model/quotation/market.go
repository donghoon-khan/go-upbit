package quotation

import "encoding/json"

type Market struct {
	Market      string `json:"market"`
	KoreanName  string `json:"korean_name"`
	EnglishName string `json:"english_name"`
	Warning     string `json:"market_warning,omitempty"`
}

type Markets struct {
	Items []Market
}

func GetMarketsFromJson(b []byte) (*Markets, error) {
	var marketCodes []Market

	e := json.Unmarshal(b, &marketCodes)
	return &Markets{
		Items: marketCodes,
	}, e
}
