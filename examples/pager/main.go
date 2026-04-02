package main

import (
	"context"
	"fmt"
	"log"

	"github.com/subhanjanOps/tornSDK/client"
	"github.com/subhanjanOps/tornSDK/market"
)

func main() {
	// Example: iterate the market auctionhouse with the SDK's raw requester.
	// This example disables the rate limiter for local demo by setting
	// RequestsPerMinute to -1. In production, provide a valid API key and
	// enable throttling.
	sdk := client.New(client.Config{
		APIKey:            "", // set your API key when running for real
		RequestsPerMinute: -1,
	})

	pager := market.NewPager(sdk, "market/auctionhouse", nil, 50)
	ctx := context.Background()

	for page := 0; ; page++ {
		data, done, err := pager.Next(ctx)
		if err != nil {
			log.Fatalf("pager.Next error: %v", err)
		}

		fmt.Printf("page=%d len=%d\n", page+1, len(data))

		if done {
			break
		}
	}
}
