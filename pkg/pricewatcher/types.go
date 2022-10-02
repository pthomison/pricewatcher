package pricewatcher

import (
	"time"

	"gorm.io/gorm"
)

type TickData struct {
	gorm.Model

	Price     float32
	Coin      string
	Currency  string
	Timestamp time.Time
}
