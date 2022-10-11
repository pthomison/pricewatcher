package coinbase

import (
	"fmt"
	"testing"
)

func TestCoinbase(t *testing.T) {

	fmt.Println(GetEthBuyPrice())
	fmt.Println(GetEthSellPrice())
	fmt.Println(GetEthSpotPrice())

}
