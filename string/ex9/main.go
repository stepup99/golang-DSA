package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello World"
	vowelsCounts := 0
	for _, ch := range str {
		switch ch {
		case 'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U':
			vowelsCounts++
		}
	}

	fmt.Println(vowelsCounts)
	anotherWay(str)
}

func anotherWay(str string) {
	i := 0
	for _, v := range str {
		if strings.ContainsRune("aeiouAEIO", v) {
			i++
		}
	}
	fmt.Println("another way -> ", i)
}
