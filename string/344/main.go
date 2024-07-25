package main

import "fmt"

func main() {
	s := []string{"h", "e", "l", "l", "o", "p"}
	start := 0
	end := len(s) - 1

	for start <= end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}

	fmt.Println(s)
}
