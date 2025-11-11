package main

import (
	_ "embed"
	"fmt"
)

var (
	//go:embed numbers.txt
	data []byte
)

func main() {
	fmt.Println(string(data))
}
