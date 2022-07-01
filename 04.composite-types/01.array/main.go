package main

import "fmt"

var x [10]int

func main() {

	fmt.Println("Default value of array:", x)

	// Assign value using x[index]
	for i := 0; i < len(x); i++ {
		var index = i
		var value = i * 5
		x[index] = value
	}
	fmt.Println("New value of x:", x)

	// Assign value using range
	for index := range x {
		x[index] += 5
	}
	fmt.Println("New value of x:", x)

}
