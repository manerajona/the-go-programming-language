package main

import (
	"fmt"
)

func main() {

	const buffer = 10

	// Create a channel with buffer=10
	channel := make(chan int, buffer)

	for i := 1; i <= buffer; i++ {
		// Supply the channel
		channel <- i
	}

	for i := buffer; i > 0; i-- {
		// Consume the channel
		fmt.Println(<-channel)
	}

}
