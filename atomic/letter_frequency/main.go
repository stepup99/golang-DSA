package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string, frequency *[26]int32) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	for _, b := range body {
		c := strings.ToLower(string(b))
		index := strings.Index(allLetters, c)
		if index >= 0 {
			frequency[index] += 1
		}
	}
}

func main() {
	fmt.Println("main fn ...........")
	start := time.Now()
	var frequency [26]int32
	for i := 1000; i <= 1002; i++ {
		countLetters(fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i), &frequency)
	}

	fmt.Println("Done ...... ")
	// fmt.Println(frequency)
	elapsed := time.Since(start)
	fmt.Println("completed in ", elapsed)
	for i, f := range frequency {
		fmt.Printf("%v -> %v \n", string(allLetters[i]), f)
	}

}
