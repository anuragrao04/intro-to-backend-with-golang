
###### Introduction To Backend With Golang

### Welcome To Class 2

---

### TODO

- From the previous class: Packages & Modules 
- Concurrency & Parallelism
	- Go Routines
	- Race Conditions
	- Mutex Locks

---

### Packages & Modules

--- 

#### Packages

- Packages are a collection of files. A collection of files _belong_ to a package.

```go
package main // <- This defined what pacakge a particular file belongs to


import "fmt"
...
```

---

#### Module

- A list of packages is belong to a module
- Every go project is a module

---

##### How Do You Make A Module

```bash
go mod init <module name>
# run the above in a terminal
# REPLACE <module name> with a name you like
```
- Intializes a module in the current directory

---

- If you want to publish your module, this must be unique globally - around the world. Usually given the same name as the Github repo it's hosted at. For example: `github.com/anuragrao04/superlit-backend`
- For the class's purposes, you may use any name you like.

---

##### The `go.mod` file

- Once you have done a `go mod init`, a `go.mod` file is created. This file contains all the external modules you've installed in your project and also some information about your own project. 
- Analogous to `package.json` in javascript

---

##### My `go.mod`

```go
module pesuio

go 1.23.1
```

---

#### What's The Use Of All This Drama? 

- All Files That Belong To The Same Package Can Share Variables, Functions and Other Identifiers.
- You can import a package and use functions and variables from that package.
- This is really useful in real world scenarios to separate work

---

#### For Example - Superlit

---

#### An Example For You To Code Up

1. Create a folder and a module inside it. Call it `supermarket`:

```bash
# inside the folder called supermarket
go mod init supermarket
```

---

2. **The Execution Of A Go Program Begins At The Main Function Inside The Main Package**. Create a `main.go` inside this folder. Populate it with the following code:

---

```go
// supermarket/main.go
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
```

---

3. Create a folder `stock` inside the `supermarket` folder. **Every File That Belongs To A Package Must Be Inside One Folder**. Then create the following files inside it:

---

```go
// supermarket/stock/init.go

package stock // <- Notice The Package Name

var Stock = 0 
// shorthand variable declaration is only allowed inside function

func Init(initQuantity int) {
	Stock = initQuantity
}
```

---

```go
// supermarket/stock/operations.go

package stock // <- Notice The Package Name

import "fmt"

func Refill(refillQuantity int) {
	Stock += refillQuantity
	fmt.Println("Refilled By", refillQuantity)
}

func Sell(sellQuantity int) {
	Stock -= selQuantity
	fmt.Println("Sold", sellQuantity)
}
```

---

```go
// supermarket/stock/printers.go

package stock // <- Notice The Package Name

import "fmt"

func PrintStock() {
	fmt.Println("The Current Stock Is: ", Stock)
}
```

---

##### Run With `go run main.go` in the `supermarket` directory

- If all goes well, you should see this:

![[Pasted image 20241013215229.png]]

---


##### Things to Note:

- The `Stock` variable is shared between files
- All variables and functions that are called from other files, regardless of which package they belong to start with a Capital Letter. 

---

### Moving On...

---

#### Concurrency & Parallelism

- Concurrency is one person juggling between tasks
- Parallelism is multiple people doing their own tasks at the same time

---

##### Go Routines - Parallelism

```go
go funcName()
```

- Just prepend `go` to a function call. It becomes a go routine, executing in parallel

---

##### An Example

```go
// primeNumbersSerial.go
// this program lists all prime numbers from 0 to 1000.
// This does not use go routines

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
		printIfPrime(i)
	}
}
```

---

```go
// primeNumbersParallel.go
// this program lists all prime numbers from 0 to 1000.
// This uses go routines

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

```

---

##### Race Conditions

- An uncertain situation that arises because the final outcome depends on which of the two or more parallel tasks finished first.

---

##### An Example

- Imagine a bank's backend system. Your account balance is represented as a number:

`₹69000`

---

- Now you attempt to deposit 100 rupees. This is done in two steps:
	1. Fetch your current balance: 69000
	2. Add 100 to it, and write it back: 69100

---

- Imagine, you have set up an auto payment to spotify to deduct 100 rupees every month. This too happens in two steps:
	1. Fetch your current balance: 69000
	2. Deduct 100 from it, and write it back: 68900

---

#### What Happens If These Two Happen At The Same Time?

---

#### If Time Allows - How To Avoid A Race Condition?

- Mutex Locks!
- Further Reading: Semaphores, Waitgroups, How do you implement mutex locks in golang?

---



