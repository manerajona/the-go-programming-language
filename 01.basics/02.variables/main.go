package main

import "fmt"

func main() {

	// Declare -> var name type
	var x int

	fmt.Println("Default values", x)

	// Assignment -> =
	x = 5

	// Declaration and assignment -> :=
	y := 8

	// Declare in the same line
	var w, z int

	// Declare block
	var (
		a, b int
		c    int
	)

	// Assign in the same line
	a, b, c, w, z = 1, 1, 2, 13, 34

	// Declare and assing in the same value
	xx, yy := 55, 89

	fmt.Println("my variables:", a, b, c, x, y, w, z, xx, yy)
}
