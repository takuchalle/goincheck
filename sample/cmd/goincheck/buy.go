package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/takuyaohashi/goincheck"
)

var buyCmd = &cobra.Command{
	Use:   "buy",
	Short: "Buy coincheck in Japanese Yen.",
	Long:  "Buy.",
	RunE:  buy,
}

func usage_buy() string {
	return "Usage: ./goincheck buy yen"
}

func buy_bitcoin(yen int) error {
	client, err := goincheck.NewClient(accessKey, secretAccesskey)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	order, err := client.OrderToMarketBuy(ctx, yen)
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", order)

	return nil
}

func buy(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New(usage_buy())
	}
	yen, err := strconv.Atoi(args[0])
	if err != nil || yen <= 0 {
		return errors.New(usage_buy())

	}

	return buy_bitcoin(yen)
}

func init() {
	RootCmd.AddCommand(buyCmd)
}
