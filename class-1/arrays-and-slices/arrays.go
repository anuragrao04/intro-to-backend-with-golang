package main

import (
	"fmt"
	"reflect"
)

func main() {
	var names [3]string = [3]string{"Anurag", "Shreya", "Amit"}
	// this is an array
	// it's fixed size. You can't just push elements into it
	fmt.Println(reflect.TypeOf(names))
	// you can reassign values inside the array
	names[0] = "Ashwini"
	fmt.Println(names)
}
