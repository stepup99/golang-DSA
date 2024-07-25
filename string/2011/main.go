package main

import "fmt"

func main() {
	operations := []string{"X++", "++X"}
	X := 0

	for _, ops := range operations {
		if ops == "X++" || ops == "++X" {
			X++
		} else {
			X--
		}
	}

	fmt.Println(X)
}
