package stock // <- Notice The Package Name

import "fmt"

func Refill(refillQuantity int) {
	Stock += refillQuantity
	fmt.Println("Refilled By", refillQuantity)
}

func Sell(sellQuantity int) {
	Stock -= sellQuantity
	fmt.Println("Sold", sellQuantity)
}
