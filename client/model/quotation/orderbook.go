package quotation

import "encoding/json"

type OrderbookUnit struct {
	AskPrice float64 `json:"ask_price"`
	BidPrice float64 `json:"bid_price"`
	AskSize  float64 `json:"ask_size"`
	BidSize  float64 `json:"bid_size"`
}

type Orderbook struct {
	Market         string          `json:"market"`
	Timestamp      int64           `json:"timestamp"`
	TotalAskSize   float64         `json:"total_ask_size"`
	TotalBidSize   float64         `json:"total_bid_size"`
	OrderbookUnits []OrderbookUnit `json:"orderbook_units"`
}

type Orderbooks struct {
	Items []Orderbook
}

func GetOrderbooksFromJson(b []byte) (*Orderbooks, error) {
	var orderbooks []Orderbook

	e := json.Unmarshal(b, &orderbooks)
	return &Orderbooks{
		Items: orderbooks,
	}, e
}
