package main

import (
	"fmt"
	"time"
)

func main() {
	go count("sheep")
	count("dolphin")
}

func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(500 * time.Millisecond)
	}
}
