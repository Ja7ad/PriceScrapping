package main

import (
	"PriceScrapping/src"
	"fmt"
	"time"
)

func main() {
	for  {
		// Create Instance form Prices struct
		prc := src.Prices{}

		// Scrapping data and point to PriceData varaible
		PriceData := src.Scrapper(&prc)

		// Show Prices to output console
		fmt.Printf("Dollar : %s IRR\n Euro : %s IRR\n Gold : %s IRR\n Coin : %s IRR\n",
			PriceData["dollar"],
			PriceData["euro"],
			PriceData["gold"],
			PriceData["coin"],
			)
		fmt.Println("=====================================")

		// Sleep time for Scrapping 5 min
		time.Sleep(300 * time.Second)
	}
}
