package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	// send
	go supplier(ch1)

	// receive
	go consumer(ch1, ch2)

	for value := range ch2 {
		fmt.Println("Task:", value, "millis")
	}

}

func supplier(c chan<- int) {
	for i := 0; i < 50; i++ {
		c <- i
	}
	close(c)
}

func consumer(c1, c2 chan int) {
	var wg sync.WaitGroup

	for v := range c1 {
		wg.Add(1)
		go func(value int) {
			c2 <- task(value)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)
}

func task(num int) (duration int) {
	duration = num + rand.Intn(500)
	time.Sleep(time.Millisecond * time.Duration(duration))
	return
}
