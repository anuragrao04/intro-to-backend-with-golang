package main

import "fmt"

func printIfPrime(i int) {
	threshold := i/2 + 1
	for k := 2; k < threshold; k++ {
		if i%k == 0 {
			return
		}
	}
	fmt.Println(i)
}

func main() {
	for i := 0; i <= 1000000; i++ {
		go printIfPrime(i)
	}
}
