package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	fmt.Print("Hello ", name) // name has the \n delimiter at the end
	fmt.Println("Your current go version:", runtime.Version())
	fmt.Printf("Running in %v_%v\n", runtime.GOOS, runtime.GOARCH)
}
