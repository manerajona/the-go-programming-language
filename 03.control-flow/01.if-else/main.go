package main

import "fmt"

func main() {

	// if true {}
	x := 1
	if x == 1 {
		fmt.Println("x is 1")
	}

	if x == 0 {
		fmt.Println("x is 0") // never get there
	}

	// if condition {} else {}
	if x%2 == 0 {
		fmt.Println("x is even")
	} else {
		fmt.Println("x is odd")
	}

	// if initialization statement; condition {}
	if y := 20; y > 0 {
		fmt.Println("y is", y)
	}

	// if initialization statement; condition {} else if condition {} else {}
	if num := 128; num < 1 {
		fmt.Println("num < 1")
	} else if num < 10 {
		fmt.Println("num < 10")
	} else if num < 100 {
		fmt.Println("num < 100")
	} else {
		fmt.Println("num > 100")
	}

}
