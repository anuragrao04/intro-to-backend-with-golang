package main

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var wg sync.WaitGroup

func main() {
	router := gin.Default()
	router.GET("/primes/serial", handleSerialPrimesRequest)
	router.GET("/primes/parallel", handleParallalPrimesRequest)
	router.Run()
}

func IsPrime(i int) bool {
	threshold := i/2 + 1
	for k := 2; k < threshold; k++ {
		if i%k == 0 {
			return false
		}
	}
	return true
}

func handleSerialPrimesRequest(c *gin.Context) {

	listOfPrimes := make([]int, 0)
	for i := 0; i <= 10000; i++ {
		if IsPrime(i) {
			listOfPrimes = append(listOfPrimes, i)
		}
	}
	c.JSON(200, gin.H{
		"primes": listOfPrimes,
		"random": uuid.New().String(),
	})

}

func handleParallalPrimesRequest(c *gin.Context) {
	listOfPrimes := make([]int, 0)
	listOfPrimesMutext := sync.Mutex{}

	for i := 0; i <= 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if IsPrime(i) {
				listOfPrimesMutext.Lock()
				listOfPrimes = append(listOfPrimes, i)
				listOfPrimesMutext.Unlock()
			}
		}()
	}

	wg.Wait()
	c.JSON(200, gin.H{"primes": listOfPrimes})
}
