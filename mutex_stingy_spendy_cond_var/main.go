package main

import (
	"fmt"
	"sync"
)

var (
	money          = 0
	mu             sync.Mutex
	moneyDeposited = sync.NewCond(&mu)
)

func stingy(done chan struct{}) {
	for i := 0; i < 100; i++ {
		mu.Lock()
		money += 10
		moneyDeposited.Broadcast() // Notify all waiting goroutines
		fmt.Println("stingy sees the balance of ", money)
		mu.Unlock()
	}
	done <- struct{}{}
}

func spendy(done chan struct{}) {
	for i := 0; i < 100; i++ {
		mu.Lock()
		for money < 20 { // Wait until there's enough money
			fmt.Println("for money < 20 { >>>>>>>>>.----- ")
			moneyDeposited.Wait()
		}
		money -= 20
		fmt.Println("spendy sees balance of ", money)
		mu.Unlock()
	}
	done <- struct{}{}
}

func main() {
	done := make(chan struct{})
	go stingy(done)
	go spendy(done)

	<-done
	fmt.Println("Final balance:", money)
}

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// var (
// 	money          = 100
// 	lock           = sync.Mutex{}
// 	moneyDeposited = sync.NewCond(&lock)
// )

// func stingy(done chan struct{}) {

// 	for i := 1; i <= 1000; i++ {
// 		lock.Lock()
// 		money += 10
// 		fmt.Println("Stingy sees balance of", money)
// 		moneyDeposited.Signal() // Signal one waiting goroutine
// 		lock.Unlock()
// 		time.Sleep(1 * time.Millisecond)
// 	}
// 	done <- struct{}{}
// 	println("Stingy Done")
// }

// func spendy(done chan struct{}) {

// 	for i := 1; i <= 1000; i++ {
// 		lock.Lock()
// 		for money-20 < 0 {
// 			fmt.Println("Spendy waiting, balance is", money)
// 			moneyDeposited.Wait() // Wait until condition is met
// 		}
// 		money -= 20
// 		fmt.Println("Spendy sees balance of", money)
// 		lock.Unlock()
// 		time.Sleep(1 * time.Millisecond)
// 	}
// 	done <- struct{}{}
// 	println("Spendy Done")
// }

// func main() {
// 	done := make(chan struct{})
// 	go stingy(done)
// 	go spendy(done)
// 	<-done
// 	fmt.Println("Final balance:", money) // Print the final balance
// }
