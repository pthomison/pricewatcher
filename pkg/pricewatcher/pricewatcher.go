package pricewatcher

import (
	"fmt"
	"io"
	"time"

	"github.com/pthomison/dbutils"
	"github.com/pthomison/dbutils/sqlite"
	"github.com/pthomison/pricewatcher/pkg/coinbase"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

const (
	CommandName = "pricewatcher"

	SleepTime = time.Second * 30
)

type Args struct {
	DBFile string
}

func RegisterFlags(cmd *cobra.Command, cmdArgs *Args) {
	cmd.PersistentFlags().StringVarP(&cmdArgs.DBFile, "dbfile", "", "pricewatcher.sqlite.db", "location for database file")
}

func Run(args *Args, output io.Writer) {
	dbc := &sqlite.SQLiteClient{
		SQLiteFile: args.DBFile,
	}

	dbc.Connect(&gorm.Config{})

	dbc.DB().AutoMigrate(&CoinbaseBuyPrice{})
	dbc.DB().AutoMigrate(&CoinbaseSellPrice{})
	dbc.DB().AutoMigrate(&CoinbaseSpotPrice{})

	for {
		buyPrice := &CoinbaseBuyPrice{}
		sellPrice := &CoinbaseSellPrice{}
		spotPrice := &CoinbaseSpotPrice{}

		buyPrice.Consume(coinbase.GetEthBuyPrice())
		sellPrice.Consume(coinbase.GetEthSellPrice())
		spotPrice.Consume(coinbase.GetEthSpotPrice())

		dbutils.Create(dbc, []*CoinbaseBuyPrice{buyPrice})
		dbutils.Create(dbc, []*CoinbaseSellPrice{sellPrice})
		dbutils.Create(dbc, []*CoinbaseSpotPrice{spotPrice})

		fmt.Fprintf(output, "%+v\n", buyPrice)
		fmt.Fprintf(output, "%+v\n", sellPrice)
		fmt.Fprintf(output, "%+v\n", spotPrice)

		time.Sleep(SleepTime)
	}
}
