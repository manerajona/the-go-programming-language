package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

	// send
	go supplier(c)

	// receive
	consumer(c)

}

// directional chan<- (send only)
func supplier(channel chan<- int) {
	for i := 1; i <= 10; i++ {
		// Supply the channel
		channel <- i
	}
}

// directional <-chan (receive only)
func consumer(channel <-chan int) {
	for i := 1; i <= 10; i++ {
		// Consume the channel
		fmt.Println(<-channel)
	}
}
