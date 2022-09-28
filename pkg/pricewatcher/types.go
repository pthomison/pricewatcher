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

// func (t *TickData) Persist(dbc dbutils.DBClient) {
// 	td := &TickData{
// 		Price:     price,
// 		Coin:      coin,
// 		Currency:  currency,
// 		Timestamp: time.Now(),
// 	}

// 	return td
// }
