package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64

const iterations = 1000

func main() {
	var wg sync.WaitGroup

	wg.Add(iterations)
	for i := 0; i < iterations; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			defer wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
