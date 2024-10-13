package main

import "fmt"

func main() {
	var x int // variable declaration
	x = 69
	var y int = 5 // declaration and initiialized
	// go can infer the type when a value is initialized to a variable
	// so we can skip the datatye
	var z = 10 // valid
	// you can also skip the var keyword and use the := shorthand
	p := 10 // valid

	fmt.Println(x, y, z, p)
}
