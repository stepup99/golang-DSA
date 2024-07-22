package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	money int32 = 0
	mu    sync.Mutex
)

func stingy(done chan struct{}) {
	for i := 0; i < 10; i++ {
		// mu.Lock()
		atomic.AddInt32(&money, 10)
		fmt.Println("stingy sees the balance of ", money)
		// mu.Unlock()
	}
	done <- struct{}{}
}

func spendy(done chan struct{}) {
	for i := 0; i < 10; i++ {
		// mu.Lock()
		atomic.AddInt32(&money, -10)
		fmt.Println("spendy sees balance of ", money)
		// mu.Unlock()
	}
	done <- struct{}{}
}

func main() {
	done := make(chan struct{})
	go stingy(done)
	go spendy(done)
	<-done
	<-done
	fmt.Println("Final balance:", money)
}
