package main

import (
	"fmt"

	"github.com/takuyaohashi/goincheck"
)

func main() {
	client, _ := goincheck.NewClient("hoge", "huga")
	tikcer, _ := client.GetTicker()
	fmt.Printf("Tikcer = %+v\n", tikcer)

	orderbook, _ := client.GetOrderBook()
	fmt.Printf("OrderBook = %+v\n", orderbook)
}
