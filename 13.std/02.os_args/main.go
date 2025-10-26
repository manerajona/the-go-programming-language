package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 || args[0] == "help" {
		fmt.Println("Usage: go run main.go [total-amount] [tip-amount]")
		fmt.Println("Example: go run main.go 20 10")
	} else {
		total, _ := strconv.ParseFloat(args[0], 32)
		tip, _ := strconv.ParseFloat(args[1], 32)
		fmt.Printf("You total will be %.2f\n", calculateTotal(float32(total), float32(tip)))
	}
}

func calculateTotal(total float32, tip float32) float32 {
	return total + (total * (tip / 100))
}
