package main

import (
	"fmt"
	"sync"
)

func main() {
	var x = 10
	var xLock sync.Mutex

	xLock.Lock()
	x = 20
	xLock.Unlock()

	fmt.Println(x)
}
