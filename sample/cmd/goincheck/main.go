package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/takuyaohashi/goincheck"
)

const (
	your_key       = ""
	your_secretkey = ""
)

func main() {
	client, _ := goincheck.NewClient(your_key, your_secretkey)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	tikcer, err := client.GetTicker(ctx)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	fmt.Printf("Tikcer = %+v\n", tikcer)
	os.Exit(0)
}
