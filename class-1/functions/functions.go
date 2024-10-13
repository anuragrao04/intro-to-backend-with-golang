package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

// functions can return more than one value

func swap(x, y int) (int, int) {
	return y, x
}

// named returns
func divide(x, y int) (remainder, quotient int) {
	remainder = x % y
	quotient = int(x / y)
	return // you don't have to mention the variables to be returned, since they are already named
}

func main() {
	fmt.Println(add(10, 5))
	fmt.Println(swap(10, 5))
	fmt.Println(divide(10, 3))
}
