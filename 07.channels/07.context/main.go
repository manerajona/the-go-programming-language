package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("Error check:", ctx.Err())
	fmt.Println("Num of goroutines:", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(100 * time.Millisecond)
				fmt.Println(n, "Working...")
			}
		}
	}()

	time.Sleep(time.Second)
	fmt.Println("Error check:", ctx.Err())
	fmt.Println("Num of goroutines:", runtime.NumGoroutine())

	cancel()
}
