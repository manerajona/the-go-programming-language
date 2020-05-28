package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a channel
	channel := make(chan int)

	// Create a new thread to supply
	go func() {
		for i := 1; i <= 10; i++ {
			// Supply the channel
			channel <- i
		}
	}()

	// Create a new thread to consume
	go func() {
		for {
			// Consume the channel
			fmt.Println(<-channel)
		}
	}()

	// Thead main finished after a second
	time.Sleep(time.Second)
}
