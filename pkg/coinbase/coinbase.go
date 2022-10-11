package coinbase

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/pthomison/errcheck"
)

const (
	ETH_PRICE_BUY_ENDPOINT  = "https://api.coinbase.com/v2/prices/ETH-USD/buy"
	ETH_PRICE_SELL_ENDPOINT = "https://api.coinbase.com/v2/prices/ETH-USD/sell"
	ETH_PRICE_SPOT_ENDPOINT = "https://api.coinbase.com/v2/prices/ETH-USD/spot"
)

type CoinbasePriceResponse struct {
	Data struct {
		AmountStr string `json:"amount"`
		Amount    float64
		Currency  string
		Coin      string `json:"base"`
		Time      int64
	}
}

func (r *CoinbasePriceResponse) Consolidate() {
	if &r.Data == nil || r.Data.AmountStr == "" {
		errcheck.Check(errors.New("Cannot consolidate a blank CoinbasePriceResponse"))
	}

	var err error

	r.Data.Amount, err = strconv.ParseFloat(r.Data.AmountStr, 64)
	errcheck.Check(err)

	r.Data.Time = time.Now().Unix()
}

func getCoinbasePrice(endpoint string) *CoinbasePriceResponse {
	r, err := http.Get(endpoint)
	errcheck.Check(err)

	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	errcheck.Check(err)

	response := &CoinbasePriceResponse{}

	err = json.Unmarshal(b, response)
	errcheck.Check(err)

	response.Consolidate()

	return response
}

func GetEthBuyPrice() *CoinbasePriceResponse {
	return getCoinbasePrice(ETH_PRICE_BUY_ENDPOINT)
}

func GetEthSellPrice() *CoinbasePriceResponse {
	return getCoinbasePrice(ETH_PRICE_SELL_ENDPOINT)
}

func GetEthSpotPrice() *CoinbasePriceResponse {
	return getCoinbasePrice(ETH_PRICE_SPOT_ENDPOINT)
}
