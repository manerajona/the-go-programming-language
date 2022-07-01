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

func supplier(channel chan<- int) {
	for i := 1; i <= 10; i++ {
		channel <- i
	}
	// IMPORTANT! to close the channel
	close(channel)
}

func consumer(channel <-chan int) {
	for value := range channel {
		fmt.Println(value)
	}
}
