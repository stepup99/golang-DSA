package main

import (
	"fmt"
)

func sender(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("waiting")
	ch <- 4
	fmt.Println("sending 4")

}

func receiver(ch chan int) {
	<-ch
}

func main() {
	ch := make(chan int, 3)
	go sender(ch)
	// time.Sleep(2 * time.Second)
	// go receiver(ch)
	// time.Sleep(4 * time.Second)
}
