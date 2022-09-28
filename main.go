package main

import (
	"os"

	"github.com/pthomison/errcheck"
	"github.com/pthomison/pricewatcher/pkg/pricewatcher"
	"github.com/spf13/cobra"
)

var (
	pricewatcherArguments = &pricewatcher.Args{}

	pricewatcherCmd = &cobra.Command{
		Use:   pricewatcher.CommandName,
		Short: pricewatcher.CommandName,
		Run:   pricewatcherRun,
	}
)

func pricewatcherRun(cmd *cobra.Command, args []string) {
	pricewatcher.Run(pricewatcherArguments, os.Stdout)
}

func Execute() error {
	return pricewatcherCmd.Execute()
}

func init() {
	pricewatcher.RegisterFlags(pricewatcherCmd, pricewatcherArguments)
}

func main() {
	errcheck.Check(Execute())
}
