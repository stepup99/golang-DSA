package main

import (
	"fmt"
	"strings"
)

func main() {
	haystack := "sabutsad"
	needle := "sad"
	idx := strings.Index(haystack, needle)
	fmt.Print(idx)
}
