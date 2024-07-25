package main

import (
	"fmt"
	"strings"
)

func main() {
	word := " fly me to the moon   joyboyy    "
	word = strings.Trim(word, " ")

	d := strings.Split(word, " ")

	dlen := len(d[len(d)-1])

	fmt.Println(dlen)
}
