package quotation

import (
	"reflect"
	"testing"
)

func TestGetCandlesFromJson(t *testing.T) {
	cases := []struct {
		json     []byte
		expected *Candles
	}{
		{
			[]byte(`[{"market":"KRW-BTC","candle_date_time_utc":"2021-08-04T14:57:00","candle_date_time_kst":"2021-08-04T23:57:00",
			"opening_price":45301000.00000000,"high_price":45340000.00000000,"low_price":45281000.00000000,"trade_price":45338000.00000000,
			"timestamp":1628089069735,"candle_acc_trade_price":155953930.04602000,"candle_acc_trade_volume":3.44149774,"unit":1}]`),
			&Candles{
				Items: []Candle{
					{
						Market:               "KRW-BTC",
						CandleDateTimeUtc:    "2021-08-04T14:57:00",
						CandleDateTimeKst:    "2021-08-04T23:57:00",
						OpeningPrice:         45301000.00000000,
						HighPrice:            45340000.00000000,
						LowPrice:             45281000.00000000,
						TradePrice:           45338000.00000000,
						TimeStamp:            1628089069735,
						CandleAccTradePrice:  155953930.04602000,
						CandleAccTradeVolume: 3.44149774,
						Unit:                 1,
					},
				},
			},
		},
		{
			[]byte(`[{"market":"KRW-BTC","candle_date_time_utc":"2021-08-04T00:00:00","candle_date_time_kst":"2021-08-04T09:00:00",
			"opening_price":44950000.00000000,"high_price":45511000.00000000,"low_price":44000000.00000000,"trade_price":45410000.00000000,
			"timestamp":1628089373345,"candle_acc_trade_price":212050153710.24596000,"candle_acc_trade_volume":4744.81280109,
			"prev_closing_price":44969000.00000000,"change_price":441000.00000000,"change_rate":0.0098067558}]`),
			&Candles{
				Items: []Candle{
					{
						Market:               "KRW-BTC",
						CandleDateTimeUtc:    "2021-08-04T00:00:00",
						CandleDateTimeKst:    "2021-08-04T09:00:00",
						OpeningPrice:         44950000.00000000,
						HighPrice:            45511000.00000000,
						LowPrice:             44000000.00000000,
						TradePrice:           45410000.00000000,
						TimeStamp:            1628089373345,
						CandleAccTradePrice:  212050153710.24596000,
						CandleAccTradeVolume: 4744.81280109,
						Unit:                 0,
						PrevClosingPrice:     44969000.00000000,
						ChangePrice:          441000.00000000,
						ChangeRate:           0.0098067558,
					},
				},
			},
		},
		// 주 캔들
		{
			[]byte(`[{"market":"KRW-BTC","candle_date_time_utc":"2021-08-04T00:00:00","candle_date_time_kst":"2021-08-04T09:00:00",
			"opening_price":44950000.00000000,"high_price":45511000.00000000,"low_price":44000000.00000000,"trade_price":45410000.00000000,
			"timestamp":1628089373345,"candle_acc_trade_price":212050153710.24596000,"candle_acc_trade_volume":4744.81280109,
			"first_day_of_period":"2021-08-02"}]`),
			&Candles{
				Items: []Candle{
					{
						Market:               "KRW-BTC",
						CandleDateTimeUtc:    "2021-08-04T00:00:00",
						CandleDateTimeKst:    "2021-08-04T09:00:00",
						OpeningPrice:         44950000.00000000,
						HighPrice:            45511000.00000000,
						LowPrice:             44000000.00000000,
						TradePrice:           45410000.00000000,
						TimeStamp:            1628089373345,
						CandleAccTradePrice:  212050153710.24596000,
						CandleAccTradeVolume: 4744.81280109,
						Unit:                 0,
						FirstDayOfPeriod:     "2021-08-02",
					},
				},
			},
		},
		{
			[]byte(`[{"market":"KRW-BTC","candle_date_time_utc":"2021-08-04T00:00:00","candle_date_time_kst":"2021-08-04T09:00:00",
			"opening_price":44950000.00000000,"high_price":45511000.00000000,"low_price":44000000.00000000,"trade_price":45410000.00000000,
			"timestamp":1628089373345,"candle_acc_trade_price":212050153710.24596000,"candle_acc_trade_volume":4744.81280109,
			"first_day_of_period":"2021-08-02"}]`),
			&Candles{
				Items: []Candle{
					{
						Market:               "KRW-BTC",
						CandleDateTimeUtc:    "2021-08-04T00:00:00",
						CandleDateTimeKst:    "2021-08-04T09:00:00",
						OpeningPrice:         44950000.00000000,
						HighPrice:            45511000.00000000,
						LowPrice:             44000000.00000000,
						TradePrice:           45410000.00000000,
						TimeStamp:            1628089373345,
						CandleAccTradePrice:  212050153710.24596000,
						CandleAccTradeVolume: 4744.81280109,
						Unit:                 0,
						FirstDayOfPeriod:     "2021-08-02",
					},
				},
			},
		},
	}

	for _, c := range cases {
		actual, e := GetCandlesFromJson(c.json)
		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("GetCandlesFromJson(%s) == \n%#v\nexpected \n%#v\n",
				c.json, actual, c.expected)
		}
		if e != nil {
			t.Errorf("error : %s", e.Error())
		}
	}
}
