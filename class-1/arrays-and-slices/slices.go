package main

import (
	"fmt"
	"reflect"
)

func main() {
	// A slice is a dynamically-sized, flexible view into the elements of an array.
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:4] // prints element from index 1 to 4 (excluding 4)
	fmt.Println(slice)  // [2 3 4]
	fmt.Println("type of array: ", reflect.TypeOf(array))
	fmt.Println("type of slice: ", reflect.TypeOf(slice))

	// you can also create a slice without an array
	dynamic_array := make([]int, 5) // 5 is the initial size of the array
	// now we can keep appending elements
	dynamic_array[0] = 1
	dynamic_array = append(dynamic_array, 69) // you can't do this with a normal array
	// uninitalized elements will be 0
	fmt.Println(dynamic_array)
}
