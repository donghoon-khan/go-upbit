package quotation

import (
	"reflect"
	"testing"
)

func TestGetTradesTicksFromJson(t *testing.T) {
	cases := []struct {
		json     []byte
		expected *Ticks
	}{
		{
			[]byte(`[{"market":"KRW-BTC","trade_date_utc":"2021-08-05","trade_time_utc":"11:13:06",
			"timestamp":1628161986000,"trade_price":44534000.00000000,"trade_volume":0.08800287,
			"prev_closing_price":45718000.00000000,"change_price":-1184000.00000000,"ask_bid":"ASK",
			"sequential_id":1628161986000003}]`),
			&Ticks{
				Items: []Tick{
					{
						Market:           "KRW-BTC",
						TradeDateUtc:     "2021-08-05",
						TradeTimeUtc:     "11:13:06",
						Timestamp:        1628161986000,
						TradePrice:       44534000.00000000,
						TradeVolume:      0.08800287,
						PrevClosingPrice: 45718000.00000000,
						ChangePrice:      -1184000.00000000,
						AskBid:           "ASK",
						SequentialId:     1628161986000003,
					},
				},
			},
		},
	}

	for _, c := range cases {
		actual, e := GetTicksFromJson(c.json)
		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("GetTicksFromJson(%s) == \n%#v\nexpected \n%#v\n",
				c.json, actual, c.expected)
		}
		if e != nil {
			t.Errorf("error : %s", e.Error())
		}
	}
}
