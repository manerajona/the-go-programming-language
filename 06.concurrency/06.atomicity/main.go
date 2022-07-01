package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int64

func main() {
	wg.Add(2)
	go count("sheep")
	go count("dolphin")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func count(thing string) {
	for i := 1; i <= 100; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)

		atomic.AddInt64(&counter, 1)

		fmt.Println(i, thing, "Counter:", atomic.LoadInt64(&counter)) // access without race
	}
	wg.Done()
}
