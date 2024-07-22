package main

import (
	"fmt"
	"sync"
)

var (
	money = 0
	mu    sync.Mutex
)

func stingy(done chan struct{}, mu *sync.Mutex) {
	for i := 0; i < 10; i++ {
		mu.Lock()
		money += 10
		fmt.Println("stingy sees the balance of ", money)
		mu.Unlock()
	}
	done <- struct{}{}
}

func spendy(done chan struct{}, mu *sync.Mutex) {
	for i := 0; i < 10; i++ {
		mu.Lock()
		money -= 10
		fmt.Println("spendy sees balance of ", money)
		mu.Unlock()
	}
	done <- struct{}{}
}

func main() {
	done := make(chan struct{})
	go stingy(done, &mu)
	go spendy(done, &mu)
	<-done
	<-done
	fmt.Println("Final balance:", money)
}
