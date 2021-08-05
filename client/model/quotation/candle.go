package quotation

import (
	"encoding/json"
)

type Candle struct {
	Market               string  `json:"market"`
	CandleDateTimeUtc    string  `json:"candle_date_time_utc"`
	CandleDateTimeKst    string  `json:"candle_date_time_kst"`
	OpeningPrice         float64 `json:"opening_price"`
	HighPrice            float64 `json:"high_price"`
	LowPrice             float64 `json:"low_price"`
	TradePrice           float64 `json:"trade_price"`
	TimeStamp            int64   `json:"timestamp"`
	CandleAccTradePrice  float64 `json:"candle_acc_trade_price"`
	CandleAccTradeVolume float64 `json:"candle_acc_trade_volume"`
	Unit                 int64   `json:"unit,omitempty"`
	PrevClosingPrice     float64 `json:"prev_closing_price,omitempty"`
	ChangePrice          float64 `json:"change_price,omitempty"`
	ChangeRate           float64 `json:"change_rate,omitempty"`
	ConvertedTradePrice  float64 `json:"converted_trade_price,omitempty"`
	FirstDayOfPeriod     string  `json:"first_day_of_period,omitempty"`
}

type Candles struct {
	Items []Candle
}

func GetCandlesFromJson(b []byte) (*Candles, error) {
	var candles []Candle

	e := json.Unmarshal(b, &candles)
	return &Candles{
		Items: candles,
	}, e
}
