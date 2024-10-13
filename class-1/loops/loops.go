package main

import "fmt"

func main() {
	// There are no while loops in go
	// There is only a for loop
	// It's defined like this:

	// for <initialization>; <condition>; <increment> {
	fmt.Println("For loop: ")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// inorder to make a while loop, you can just ommit the initialization and increment

	n := 0

	fmt.Println("While loop (actually a for): ")
	for n < 10 {
		fmt.Println(n)
		n++
	}

	// you can also ommit any combination of parts in the for loop.
	// For example, you can ommit the initialization part, but keep the condition and increment/decrement

	k := 0
	fmt.Println("For loop without initialization: ")
	for ; k < 10; k++ {
		fmt.Println(k)
	}

	// you can make an infinite loop like this:

	// for {
	//   fmt.Println("Infinite Loop")
	// }

}
