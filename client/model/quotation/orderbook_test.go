package quotation

import (
	"reflect"
	"testing"
)

func TestGetOrderbooksFromJson(t *testing.T) {
	cases := []struct {
		json     []byte
		expected *Orderbooks
	}{
		{
			[]byte(`[
				{"market":"KRW-BTC","timestamp":1628164914070,"total_ask_size":6.76294392,"total_bid_size":2.10571502,
				"orderbook_units":[
					{"ask_price":4.4271E7,"bid_price":4.427E7,"ask_size":0.03523783,"bid_size":0.88567329},
					{"ask_price":5.4271E7,"bid_price":5.427E7,"ask_size":1.03523783,"bid_size":1.88567329}]}]`),
			&Orderbooks{
				Items: []Orderbook{
					{
						Market:       "KRW-BTC",
						Timestamp:    1628164914070,
						TotalAskSize: 6.76294392,
						TotalBidSize: 2.10571502,
						OrderbookUnits: []OrderbookUnit{
							{
								AskPrice: 4.4271e7,
								BidPrice: 4.427e7,
								AskSize:  0.03523783,
								BidSize:  0.88567329,
							},
							{
								AskPrice: 5.4271e7,
								BidPrice: 5.427e7,
								AskSize:  1.03523783,
								BidSize:  1.88567329,
							},
						},
					},
				},
			},
		},
	}

	for _, c := range cases {
		actual, e := GetOrderbooksFromJson(c.json)
		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("GetOrderbooksFromJson(%s) == \n%#v\nexpected \n%#v\n",
				c.json, actual, c.expected)
		}
		if e != nil {
			t.Errorf("error : %s", e.Error())
		}
	}
}
