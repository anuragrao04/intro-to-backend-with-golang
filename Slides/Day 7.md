###### Introduction To Backend With Golang

### Welcome To Class 7

---

### TODO

1. **Load Testing?**
	1. Why?
	2. What?
	3. How?

3. **Locust**

5. **Create**

---

#### Load Testing - Why?

- How do you know how many users can your backend system handle?
- At what point does it fail?

---

#### Load Testing - What?

- Swarm your backend system with tonnes of requests. See how it performs
- AKA Swarm your backend with users

---

#### Load Testing - How?

- Can you get a thousand users at a snap? Whenever you want? 
- I don't have that many friends, do you?
- All these users need to do is send requests. You can automate this.

---

### Locust

[https://locust.io/](https://locust.io/)

- Is a python library that helps you do load testing
- Why's it called locust? It swarms your backend just like how locusts swarm fields of crops.

---

#### Let's create

- Remember the program we wrote earlier to get prime numbers in parallel and in series? Let's convert that into an API.
- The code for which is below. You can copy and paste it. 

```go
// main.go
package main

import (
	"sync"

	"github.com/gin-gonic/gin"
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

```

---

#### Locust 

- Let's first set up python. You can use the same folder you used for the golang code. 
- Set up a python virtual environment:

```bash
python3 -m venv .venv

# for windows:
.venv\Scripts\activate.bat

# for linux and mac:
source .venv/bin/activate
```

---

#### Locust - Hands on.

- I want to show you how to read and use documentation, instead of just using LLMs. 
- LLMs often miss lead you, on to a deeper hole than you're already in
- So this demo will be you coming along with me, reading documentation and figuring locust out. 

---

#### Assignment

- You have created a backend for sign in and sign up for the assignment. Test these routes using locust. How many users can sign in at the same time?
- Remember, this is a post request. So you also have to include the post request body in the locustfile (i.e., the username and password)
- For the sign up route, you'll need to randomize the username and password since signing up with the same user again is pointless. That user is already signed up. 

---






