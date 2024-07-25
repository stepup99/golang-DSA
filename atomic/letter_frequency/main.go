package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var mu sync.Mutex

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string, frequency *[26]int32, wg *sync.WaitGroup) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	for _, b := range body {
		c := strings.ToLower(string(b))
		// mu.Lock()
		index := strings.Index(allLetters, c)
		if index >= 0 {
			atomic.AddInt32(&frequency[index], 1)
		}
		// mu.Unlock()
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	fmt.Println("main fn ...........")
	start := time.Now()
	var frequency [26]int32
	for i := 1000; i <= 1000; i++ {
		wg.Add(1)
		go countLetters(fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i), &frequency, &wg)
	}

	fmt.Println("Done ...... ")
	// fmt.Println(frequency)
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("completed in ", elapsed)
	for i, f := range frequency {
		fmt.Printf("%v -> %v \n", string(allLetters[i]), f)
	}

}
