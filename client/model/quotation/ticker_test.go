package quotation

import (
	"reflect"
	"testing"
)

func TestGetTickersFromJson(t *testing.T) {
	cases := []struct {
		json     []byte
		expected *Tickers
	}{
		{
			[]byte(`[{"market":"KRW-BTC","trade_date":"20210805","trade_time":"115035",
			"trade_date_kst":"20210805","trade_time_kst":"205035","trade_timestamp":1628164235000,
			"opening_price":45718000.00000000,"high_price":45877000.00000000,"low_price":44250000.0,
			"trade_price":44362000.0,"prev_closing_price":45718000.00000000,"change":"FALL",
			"change_price":1356000.00000000,"change_rate":0.0296600901,"signed_change_price":-1356000.00000000,
			"signed_change_rate":-0.0296600901,"trade_volume":0.00571864,"acc_trade_price":178105486470.681000000,
			"acc_trade_price_24h":321511335721.21606000,"acc_trade_volume":3954.96824729,
			"acc_trade_volume_24h":7112.07039886,"highest_52_week_price":81994000.00000000,
			"highest_52_week_date":"2021-04-14","lowest_52_week_price":11860000.00000000,
			"lowest_52_week_date":"2020-09-07","timestamp":1628164235919}]`),
			&Tickers{
				Items: []Ticker{
					{
						Market:             "KRW-BTC",
						TradeDate:          "20210805",
						TradeTime:          "115035",
						TradeDateKst:       "20210805",
						TradeTimeKst:       "205035",
						TradeTimestamp:     1628164235000,
						OpeningPrice:       45718000.00000000,
						HighPrice:          45877000.00000000,
						LowPrice:           44250000.0,
						TradePrice:         44362000.0,
						PrevClosingPrice:   45718000.00000000,
						Change:             "FALL",
						ChangePrice:        1356000.00000000,
						ChangeRate:         0.0296600901,
						SignedChangePrice:  -1356000.00000000,
						SignedChangeRate:   -0.0296600901,
						TradeVolume:        0.00571864,
						AccTradePrice:      178105486470.681000000,
						AccTradePrice24h:   321511335721.21606000,
						AccTradeVolume:     3954.96824729,
						AccTradeVolume24h:  7112.07039886,
						Highest52WeekPrice: 81994000.00000000,
						Highest52WeekDate:  "2021-04-14",
						Lowest52WeekPrice:  11860000.00000000,
						Lowest52WeekDate:   "2020-09-07",
						Timestamp:          1628164235919,
					},
				},
			},
		},
	}

	for _, c := range cases {
		actual, e := GetTickersFromJson(c.json)
		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("GetTickersFromJson(%s) == \n%#v\nexpected \n%#v\n",
				c.json, actual, c.expected)
		}
		if e != nil {
			t.Errorf("error : %s", e.Error())
		}
	}
}
