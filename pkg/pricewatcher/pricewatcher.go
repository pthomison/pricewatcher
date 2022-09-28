package pricewatcher

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/pthomison/dbutils"
	"github.com/pthomison/dbutils/sqlite"
	"github.com/pthomison/errcheck"
	"github.com/spf13/cobra"
	coingecko "github.com/superoo7/go-gecko/v3"
	"gorm.io/gorm"
)

const (
	CommandName = "pricewatcher"

	SleepTime = time.Second * 30
)

type Args struct {
	Coin     string
	Currency string
}

func RegisterFlags(cmd *cobra.Command, cmdArgs *Args) {
	cmd.PersistentFlags().StringVarP(&cmdArgs.Coin, "coin", "", "ethereum", "coin to price")
	cmd.PersistentFlags().StringVarP(&cmdArgs.Currency, "currency", "", "usd", "currency to price in")
}

func Run(args *Args, output io.Writer) {
	cg := newCGClient()
	dbc := &sqlite.SQLiteClient{
		SQLiteFile: "pricewatcher.gorm",
	}

	dbc.Connect(&gorm.Config{})

	dbc.DB().AutoMigrate(&TickData{})

	for {
		price := getPrice(cg, args)

		td := &TickData{
			Price:     price,
			Coin:      args.Coin,
			Currency:  args.Currency,
			Timestamp: time.Now(),
		}

		arrTD := []*TickData{td}

		dbutils.Create(dbc, arrTD)

		fmt.Fprintf(output, "%+v\n", td)

		time.Sleep(SleepTime)
	}
}

func newCGClient() *coingecko.Client {
	httpClient := &http.Client{
		Timeout: SleepTime,
	}
	CG := coingecko.NewClient(httpClient)

	return CG
}

func getPrice(cg *coingecko.Client, args *Args) float32 {
	ids := []string{args.Coin}
	vc := []string{args.Currency}

	sp, err := cg.SimplePrice(ids, vc)
	errcheck.Check(err)

	price := (*sp)[args.Coin][args.Currency]

	return price
}
