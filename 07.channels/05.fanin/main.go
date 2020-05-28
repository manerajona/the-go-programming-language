package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {

	eve := make(chan int)
	odd := make(chan int)
	fanin := make(chan string)

	// send
	go supplier(eve, odd)

	// receive
	go consumer(eve, odd, fanin)

	for v := range fanin {
		fmt.Println("Value from fanin:", v)
	}

}

func supplier(e, o chan<- int) {

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
}

func consumer(e, o <-chan int, fanin chan<- string) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range e {
			fanin <- strconv.Itoa(v) + "(eve)"
		}
		wg.Done()
	}()

	go func() {
		for v := range o {
			fanin <- strconv.Itoa(v) + "(odd)"
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}
