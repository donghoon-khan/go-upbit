package quotation

import (
	"encoding/json"
)

type Tick struct {
	Market           string  `json:"market"`
	TradeDateUtc     string  `json:"trade_date_utc"`
	TradeTimeUtc     string  `json:"trade_time_utc"`
	Timestamp        int64   `json:"timestamp"`
	TradePrice       float64 `json:"trade_price"`
	TradeVolume      float64 `json:"trade_volume"`
	PrevClosingPrice float64 `json:"prev_closing_price"`
	ChangePrice      float64 `json:"change_price"`
	AskBid           string  `json:"ask_bid"`
	SequentialId     int64   `json:"sequential_id"`
}

type Ticks struct {
	Items []Tick
}

func GetTicksFromJson(b []byte) (*Ticks, error) {
	var ticks []Tick

	e := json.Unmarshal(b, &ticks)
	return &Ticks{
		Items: ticks,
	}, e
}
