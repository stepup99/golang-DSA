package main

import "fmt"

func main() {
	address := "1.1.1.1"
	var defang string

	for _, ch := range address {
		if ch == '.' {
			defang += string("[.]")
		} else {
			defang += string(ch)
		}
	}
	fmt.Println("defang >>>>> ", defang)
}
