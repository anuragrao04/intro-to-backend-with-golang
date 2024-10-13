package main

import (
	"supermarket/stock"
	// this package will be responsible for keeping stock in a supermarket
)

func main() {
	stock.Init(50)
	stock.PrintStock()
	stock.Sell(10)
	stock.PrintStock()
	stock.Refill(5)
	stock.PrintStock()
}
