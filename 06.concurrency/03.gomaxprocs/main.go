package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	/* GOMAXPROCS sets the maximum number of CPUs that can be
	executing simultaneously and returns the previous setting.*/
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(2)
	go count("sheep")
	go count("dolphin")
	wg.Wait()
}

func count(thing string) {
	for i := 1; i < 100; i++ {
		fmt.Println(i, thing)
		time.Sleep(10 * time.Millisecond)
	}
	wg.Done()
}
