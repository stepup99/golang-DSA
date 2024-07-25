package main

import "fmt"

func main() {
	necklace := "---------oxo------------"

	for idx, pearl := range necklace {
		if pearl == 'x' {
			fmt.Printf("position of the pendant is %d", idx)
		}
	}
}
