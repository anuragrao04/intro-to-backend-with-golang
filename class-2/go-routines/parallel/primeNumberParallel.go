package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	for i := 0; i <= 1000000; i++ {
		wg.Add(1) // Increment the counter for each goroutine
		go printIfPrime(i)
	}

	wg.Wait() // Wait for all goroutines to finish
}

func printIfPrime(i int) {
	threshold := i/2 + 1
	for k := 2; k < threshold; k++ {
		if i%k == 0 {
			wg.Done() // Decrement the counter when the goroutine is done
			return
		}
	}
	fmt.Println(i)
	wg.Done() // Decrement the counter when the goroutine is done
}
