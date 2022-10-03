package pricewatcher

import (
	"gorm.io/gorm"
)

type TickData struct {
	gorm.Model

	Price         float32
	Coin          string
	Currency      string
	UnixTimestamp int64
}
