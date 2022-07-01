package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/* The flag -race help us find race conditions
$ go run -race main.go
*/

var wg sync.WaitGroup
var counter int // shared resource

func main() {
	wg.Add(2)
	go count("sheep")
	go count("dolphin")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func count(thing string) {
	for i := 1; i <= 100; i++ {
		x := counter + 1
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		counter = x

		fmt.Println(i, thing, "Counter:", counter)
	}
	wg.Done()
}
