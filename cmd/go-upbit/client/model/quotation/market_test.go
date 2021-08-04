package quotation

import (
	"reflect"
	"testing"
)

func TestGetMarketsFromJson(t *testing.T) {
	cases := []struct {
		json     []byte
		expected *Markets
	}{
		{
			[]byte(`[{"market":"KRW-BTC", "korean_name":"비트코인", "english_name":"Bitcoin", "market_warning":"NONE"}]`),
			&Markets{
				Items: []Market{
					{
						Market:      "KRW-BTC",
						KoreanName:  "비트코인",
						EnglishName: "Bitcoin",
						Warning:     "NONE",
					},
				},
			},
		},
		{
			[]byte(`[{"market":"USDT-LTC", "korean_name":"라이트코인", "english_name":"Litecoin", "market_warning":""},
			{"market":"BTC-MANA", "korean_name":"디센트럴랜드", "english_name":"Decentraland", "market_warning":"CAUTION"}]`),
			&Markets{
				Items: []Market{
					{
						Market:      "USDT-LTC",
						KoreanName:  "라이트코인",
						EnglishName: "Litecoin",
						Warning:     "",
					},
					{
						Market:      "BTC-MANA",
						KoreanName:  "디센트럴랜드",
						EnglishName: "Decentraland",
						Warning:     "CAUTION",
					},
				},
			},
		},
	}

	for _, c := range cases {
		actual, e := GetMarketsFromJson(c.json)
		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("GetMarketsFromJson(%s) == \n%#v\nexpected \n%#v\n",
				c.json, actual, c.expected)
		}
		if e != nil {
			t.Errorf("error : %s", e.Error())
		}
	}
}
