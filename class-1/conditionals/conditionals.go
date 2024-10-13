package main

import "fmt"

func main() {
	var name string
	fmt.Print("Enter Your Name: ")
	fmt.Scanln(&name)
	if name == "Anurag" {
		fmt.Println("Hello Anurag")
	} else if name == "Shreya" {
		fmt.Println("Hello Shreya")
	} else {
		fmt.Println("Hello Stranger")
	}
}
