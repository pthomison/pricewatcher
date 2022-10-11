package pricewatcher

import (
	"github.com/pthomison/pricewatcher/pkg/coinbase"
	"gorm.io/gorm"
)

type CoinbaseBuyPrice struct {
	gorm.Model

	Price    float64
	Coin     string
	Currency string
	Time     int64
}

func (c *CoinbaseBuyPrice) Consume(resp *coinbase.CoinbasePriceResponse) {
	c.Price = resp.Data.Amount
	c.Coin = resp.Data.Coin
	c.Currency = resp.Data.Currency
	c.Time = resp.Data.Time
}

type CoinbaseSellPrice struct {
	gorm.Model

	Price    float64
	Coin     string
	Currency string
	Time     int64
}

func (c *CoinbaseSellPrice) Consume(resp *coinbase.CoinbasePriceResponse) {
	c.Price = resp.Data.Amount
	c.Coin = resp.Data.Coin
	c.Currency = resp.Data.Currency
	c.Time = resp.Data.Time
}

type CoinbaseSpotPrice struct {
	gorm.Model

	Price    float64
	Coin     string
	Currency string
	Time     int64
}

func (c *CoinbaseSpotPrice) Consume(resp *coinbase.CoinbasePriceResponse) {
	c.Price = resp.Data.Amount
	c.Coin = resp.Data.Coin
	c.Currency = resp.Data.Currency
	c.Time = resp.Data.Time
}
