package main

import (
	"fmt"
	"sync"
)

var counter int64

const iterations = 1000

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	wg.Add(iterations)
	for i := 0; i < iterations; i++ {
		go func() {
			mutex.Lock()
			defer mutex.Unlock()
			counter++
			defer wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
