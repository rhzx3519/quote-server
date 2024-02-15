package quote

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var LISTED_SYMBOLS = []string{"AAPL", "TSLA", "NVDA", "ARKK"}

type Pool struct {
	quit   chan string
	stocks map[string]Quote
}

func NewPool() (*Pool, error) {
	var stocks = make(map[string]Quote)
	for _, symbol := range LISTED_SYMBOLS {
		current := 300 + rand.Intn(100)
		stocks[symbol] = Quote{
			Symbol:    symbol,
			Code:      symbol,
			Exchange:  "NASDAQ",
			Name:      symbol,
			Status:    "1",
			Current:   fmt.Sprintf("%.2f", float64(current)),
			Currency:  "USD",
			Timestamp: time.Now().Unix(),
			Open:      "open",
			LastClose: fmt.Sprintf("%.2f", float64(current-rand.Intn(10))),
			High:      fmt.Sprintf("%.2f", float64(current+rand.Intn(10))),
			Low:       fmt.Sprintf("%.2f", float64(current-rand.Intn(10))),
			AvgPrice:  fmt.Sprintf("%.2f", float64(current)),
			Volume:    "100",
			Amount:    "0",
			Amplitude: "88",
		}
	}
	return &Pool{
		quit:   make(chan string),
		stocks: stocks,
	}, nil
}

func (p *Pool) Start() error {
	var err error

	return err
}

func (p *Pool) GetQuote(symbol string, exchange string) (*Item, error) {
	var item *Item
	for _, sb := range LISTED_SYMBOLS {
		if sb == symbol {
			q := p.stocks[sb]
			current, _ := strconv.ParseFloat(q.Current, 64)
			q.Current = fmt.Sprintf("%.2f", current+rand.Float64())
			item = &Item{
				Market: Market{
					Region:   "US",
					Status:   "1",
					TimeZone: time.FixedZone("", 0).String(),
				},
				Quote: q,
			}
		}
	}

	return item, nil
}
