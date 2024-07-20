package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *SafeCounter) Increment(wg *sync.WaitGroup) {
	defer wg.Done()
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *SafeCounter) Decrement(wg *sync.WaitGroup) {
	defer wg.Done()
	c.mu.Lock()
	c.count--
	c.mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	s := SafeCounter{count: 0}

	for i := 1; i < 10; i++ {
		wg.Add(1)
		go s.Increment(&wg)
	}

	for i := 1; i < 10; i++ {
		wg.Add(1)
		go s.Decrement(&wg)
	}

	wg.Wait()

	fmt.Println(s.count)
}
