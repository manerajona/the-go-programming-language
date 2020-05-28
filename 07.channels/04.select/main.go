package main

import (
	"fmt"
)

func main() {

	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go supplier(eve, odd, quit)

	// receive
	consumer(eve, odd, quit)

}

func supplier(e, o, q chan<- int) {

	for i := 100; true; i-- {
		switch {
		case i == 0:
			q <- 0
			return
		case i%2 == 0:
			e <- i
		case i%2 == 1:
			o <- i
		}

	}
}

func consumer(e, o, q <-chan int) {
	for {
		select {
		case value := <-e:
			fmt.Println("Value eve:", value)
		case value := <-o:
			fmt.Println("Value odd:", value)
		case value := <-q:
			fmt.Println("Value quit:", value)
			return
		}
	}
}
