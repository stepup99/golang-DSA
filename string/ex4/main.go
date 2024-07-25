package main

import "fmt"

func main() {
	necklace := "-o-o-o-o-o-o-o-o-o-o-o-o-o-o-"
	count := 0
	for _, v := range necklace {
		if v == 'o' {
			count++
		}
	}
	fmt.Printf("number of pearls is rosy %d\n", count)
}
