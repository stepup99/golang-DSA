package main

import (
	"fmt"
	"time"
)

func data(ch chan string) {
	fmt.Println("i am waiting >>>>>>>  ")
	time.Sleep(4 * time.Second)
	ch <- "Hello World"
}

func main() {
	ch := make(chan string)
	go data(ch)
	// go func() {
	// 	fmt.Println("i am waiting >>>>>>>  ")
	// 	time.Sleep(4 * time.Second)
	// 	ch <- "Hello World"
	// }()

	msg := <-ch
	fmt.Println(msg)
}
