package main

import "fmt"

func ToLower(s string) string {
	var ch string
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			ch += string(v + 32)
		} else {
			ch += string(v)
		}
	}
	return ch
}

func main() {
	shrink := "HONEY, i SHRUNK THE KIDS"
	result := ToLower(shrink)

	fmt.Println(result)
}
