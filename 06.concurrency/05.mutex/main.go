package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex

func main() {
	wg.Add(2)
	go count("sheep")
	go count("dolphin")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func count(thing string) {
	for i := 1; i <= 100; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)

		mutex.Lock()
		{
			counter++
			fmt.Println(i, thing, "Counter:", counter)
		}
		mutex.Unlock()
	}
	wg.Done()
}
