package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("./big.txt")
	text := string(data)
	if err != nil {
		panic("failed to read from big.txt")
	}

	vowels := "aeiouAEIOU"
	counter := 0
	// reading serially
	for _, c := range text {
		// if character is a vowel
		if strings.ContainsRune(vowels, c) {
			counter = counter + 1
		}
	}

	fmt.Println(counter)

}
