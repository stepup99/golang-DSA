package main

import (
	"fmt"
	"strings"
)

func reverse(words []string) []string {
	start := 0
	end := len(words) - 1
	for start <= end {
		words[start], words[end] = words[end], words[start]
		start++
		end--
	}
	return words
}

func main() {
	words := strings.Fields(strings.Trim(" a good   example ", " "))
	fmt.Println(words)

	fmt.Println(strings.Join(reverse(words), " "))
}
