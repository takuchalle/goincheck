package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/takuyaohashi/goincheck"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get Ticker.",
	Long:  "Get Ticker. For further information, type ./goincheck get usage",
	RunE:  get,
}

func usage() string {
	return "Usage: ./goincheck get [ticker]"
}

func get_ticker() error {
	client, err := goincheck.NewClient(accessKey, secretAccesskey)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ticker, err := client.GetTicker(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("Ticker\n")
	fmt.Printf("last:\t %v\n", ticker.Last)
	fmt.Printf("bid:\t %v\n", ticker.Bid)
	fmt.Printf("ask:\t %v\n", ticker.Ask)
	fmt.Printf("high:\t %v\n", ticker.High)
	fmt.Printf("low:\t %v\n", ticker.Low)
	fmt.Printf("volume:\t %v\n", ticker.Volume)
	return nil
}

func get(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New(usage())
	}

	switch args[0] {
	case "ticker":
		return get_ticker()
	default:
		return errors.New(usage())
	}
	return nil
}

func init() {
	RootCmd.AddCommand(getCmd)
}
