package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("./big.txt")
	if err != nil {
		panic("failed to read from big.txt")
	}
	text := string(data)
	textLength := len(text)
	counterChannel := make(chan int, textLength/2048+1)
	// let's take one chunk size to be 2048 characters
	chunkCount := 0
	for i := 0; i < textLength; i += 2048 {
		var chunk string
		if i+2048 >= len(text) {
			chunk = text[i:]
			chunkCount++
			go countVowels(chunk, counterChannel)
			break
		} else {
			chunk = text[i : i+2048]
		}
		chunkCount++
		go countVowels(chunk, counterChannel)
	}

	// collect all of the sums
	sum := 0
	for i := 0; i < chunkCount; i++ {
		sum += <-counterChannel
	}

	fmt.Println(sum)
}

func countVowels(chunk string, counterChannel chan int) {
	vowels := "aeiouAEIOU"
	counter := 0
	for _, c := range chunk {
		if strings.ContainsRune(vowels, c) {
			counter += 1
		}
	}
	counterChannel <- counter
}
