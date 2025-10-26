package main

import (
	"flag"
	"fmt"
)

func main() {
	arch := flag.String("arch", "amd64", "cpu architecture")
	flag.Parse()

	switch *arch {
	case "amd64":
		fmt.Println("Running in 64 bit mode")
	case "x86":
		fmt.Println("Running in 32 bit mode")
	default:
		fmt.Println("Running in unknown architecture")
	}
}
