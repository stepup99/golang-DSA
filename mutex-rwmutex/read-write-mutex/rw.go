package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type SafeMap struct {
	mu   sync.RWMutex
	data map[string]int
}

func (sm *SafeMap) Read(key string, id int, wg *sync.WaitGroup) (int, bool) {
	defer wg.Done()
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	fmt.Printf("Reader %v starting with id->%d \n\n", key, id)
	value, exists := sm.data[key]
	if !exists {
		fmt.Printf("Reader key  %v does not exists with id->%d \n\n", key, id)
	}
	fmt.Printf("Reader key ->%v , value ->%v  id->%d\n\n", key, sm.data[key], id)

	fmt.Printf("Reader %v done id->%d \n\n", key, id)

	return value, exists
}

func (sm *SafeMap) Write(key string, id int, value int, wg *sync.WaitGroup) {
	defer wg.Done()
	sm.mu.Lock()
	defer sm.mu.Unlock()
	fmt.Printf("Writer %v starting , id -> %v \n\n", key, id)
	sm.data[key] = value
	fmt.Printf("Writer, key ->%v, value -> %v  , id -> %v \n\n", key, value, id)
	fmt.Printf("Writer %v done , id -> %v \n\n", key, id)
}

func main() {
	sm := SafeMap{data: make(map[string]int)}
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go sm.Write("key1", i, i*10, &wg)
	}
	for i := 1; i < 4; i++ {
		wg.Add(1)
		go sm.Read("key1", i, &wg)
	}

	wg.Wait()
}
