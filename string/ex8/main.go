package main

import (
	"fmt"
	"strings"
)

func main() {
	garland := "rose.jasmine.tulip.hydra"
	flowers := strings.Split(garland, ".")
	fmt.Println(flowers)
	for _, v := range flowers {
		fmt.Println(v)
	}
}
