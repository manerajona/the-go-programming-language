package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/trace"
	"time"
)

func main() {
	// open file using go tool trace tracing.out
	f, _ := os.Create("tracing.out")
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("Error closing file: %v", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("Error starting tracing: %v", err)
	}
	defer trace.Stop()

	RandomNumber()
}

func RandomNumber() {
	foo := rand.Intn(1000)
	time.Sleep(time.Duration(rand.Intn(foo)) * time.Millisecond)
	fmt.Printf("foo = %d\n", foo)
}
