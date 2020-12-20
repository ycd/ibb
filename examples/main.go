package main

import (
	"context"
	"fmt"

	"github.com/ycd/ibb"
)

func main() {

	// Initialize the client
	ibb := ibb.NewClient()

	ctx := context.Background()

	// Get information about ticket prices of 2020
	ticketP2020, _ := ibb.IETTTicketPrices(ctx, 2020)
	for _, v := range ticketP2020.Records {
		fmt.Println(v)
	}

	// Get information about ISPARK's parking lots
	ispark, _ := ibb.IsparkParkingLots(ctx)
	for _, v := range ispark.Records {
		fmt.Println(v)
	}

	occupancyRates, _ := ibb.DamOccupancyRates(ctx)
	for _, v := range occupancyRates.Records {
		fmt.Println(v)
	}

}
