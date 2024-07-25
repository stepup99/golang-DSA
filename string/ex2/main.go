package main

import "fmt"

func main() {
	road := "left-right-left-right"
	// for i := 0; i < len(road); i++ {
	// 	fmt.Printf("%c -  %T \n", road[i], road[i])
	// }
	for _, v := range road {
		fmt.Printf("char is - %v - type %T\n", v, v)
	}
}
