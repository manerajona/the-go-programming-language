package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// Add a counter of how many goroutines will run
	wg.Add(2)

	go count("sheep")
	go count("dolphin")

	// Main thead wait till the wg counter is zero
	wg.Wait()
}

func count(thing string) {
	for i := 1; i <= 100; i++ {
		fmt.Println(i, thing)
		time.Sleep(10 * time.Millisecond)
	}
	// decrement the wg counter
	wg.Done()
}
